package ethiotex

import (
	"context"
	"math"
	"math/big"
	"time"

	"github.com/IoTube-analytics/go-iotube-analytics/pkg/bridge"
	eth_bridge "github.com/IoTube-analytics/go-iotube-analytics/pkg/bridge/eth"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/contracts/tokenCashier"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/logging"
	typ "github.com/IoTube-analytics/go-iotube-analytics/pkg/types"
	"github.com/davecgh/go-spew/spew"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	log "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type TransactionTracker struct {
	logger log.Logger
	cfg    Config
	ctx    context.Context
	cncl   context.CancelFunc
	client *ethclient.Client
	store  *bridge.Store
	// Map: token address ->  token symbol.
	tokens map[string]eth_bridge.ERC20
}

func NewTransactionTracker(ctx context.Context, client *ethclient.Client, logger log.Logger, cfg Config, store *bridge.Store) (*TransactionTracker, error) {
	filterLog, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	logger = log.With(filterLog, "component", ComponentName)
	// Getting tokens.
	ctx1, _ := context.WithTimeout(ctx, 10*time.Second)
	tokens, err := eth_bridge.GetTokenList(ctx1, client, logger, StandardTokenListAddress, ProxyTokenListAddress)
	if err != nil {
		return nil, errors.Wrap(err, "getting token list")
	}
	level.Debug(logger).Log("msg", "supported tokens", "list", spew.Sdump(tokens))
	ctx, cncl := context.WithCancel(ctx)
	return &TransactionTracker{
		logger: logger,
		cfg:    cfg,
		ctx:    ctx,
		cncl:   cncl,
		store:  store,
		client: client,
		tokens: tokens,
	}, nil
}

func (self *TransactionTracker) Stop() {
	self.cncl()
	level.Debug(self.logger).Log("msg", "eth tx tracker stopped")
}

func (self *TransactionTracker) Start() error {
	level.Debug(self.logger).Log("msg", "eth tx tracker started")
	// Ethereum blocktime ticker.
	ticker := time.NewTicker(20 * time.Second)
	for {
		select {
		case <-self.ctx.Done():
			return nil
		default:
		}
		var (
			fromBlockNo, toBlockNo *big.Int
		)

		// Get last checked block number from the db.
		lastCheckedBlockNo, _ := self.store.LastCheckedBlockNo(typ.NetEthereum)
		if lastCheckedBlockNo == nil {
			fromBlockNo = big.NewInt(TokenCashierStartBlockNo)
			level.Info(self.logger).Log("msg", "watching ethereum blockchain for the first time")
		} else {
			// Look ahead one block to make sure we didn't miss any new invoices.
			fromBlockNo = lastCheckedBlockNo
		}
		// Calculating head block number minus 18.
		header, err := self.client.HeaderByNumber(self.ctx, nil)
		if err != nil {
			level.Error(self.logger).Log("msg", "getting latest block header", "err", err)
			<-ticker.C
			continue
		}

		// Min block to loop over.
		min := math.Min(float64(fromBlockNo.Uint64()+blockLimitBeforeCommit),
			float64(header.Number.Uint64()))
		toBlockNo = big.NewInt(int64(min))

		if toBlockNo.Cmp(fromBlockNo) == -1 {
			level.Debug(self.logger).Log("msg", "no new block to check, waiting...")
			<-ticker.C
			continue
		}
		level.Info(self.logger).Log("msg", "checking for new transactions",
			"fromBlockNo", fromBlockNo,
			"toBlockNo", toBlockNo,
		)
		txs, err := self.traverse(fromBlockNo, toBlockNo)
		if err != nil {
			level.Error(self.logger).Log("msg", "traversing the eth blockchain",
				"err", err,
				"fromBlockNo", fromBlockNo,
				"toBlockNo", toBlockNo,
			)
			<-ticker.C
			continue
		}
		level.Info(self.logger).Log("msg",
			"new transactions count",
			"count", len(txs),
		)
		// Lets commit txs to the database.
		// FIXME: better logic needed here.
		err = self.store.RecordTxs(txs)
		if err != nil {
			level.Error(self.logger).Log("msg", "recording txs",
				"err", err,
				"fromBlockNo", fromBlockNo,
				"toBlockNo", toBlockNo,
			)
		} else {
			err = self.store.UpdateLastCheckedBlockNo(toBlockNo, typ.NetEthereum)
			if err != nil {
				level.Error(self.logger).Log("msg", "updating eth blockchain state",
					"err", err,
					"lastCheckedBlockNo", toBlockNo,
				)
			}
		}

		<-ticker.C
	}
}

func (self *TransactionTracker) traverse(fromBlockNo, toBlockNo *big.Int) ([]typ.Transaction, error) {
	txs := make([]typ.Transaction, 0)

	tokenCashierFilterer, err := tokenCashier.NewTokenCashierFilterer(TokenCashierAddress, self.client)
	if err != nil {
		return nil, errors.Wrap(err, "getting tokenCashierFilterer")
	}
	ctx, _ := context.WithTimeout(self.ctx, 10*time.Second)
	level.Info(self.logger).Log("msg",
		"filtering Receipt events",
	)
	end := toBlockNo.Uint64()
	iter, err := tokenCashierFilterer.FilterReceipt(&bind.FilterOpts{Context: ctx,
		Start: fromBlockNo.Uint64(),
		End:   &end,
	}, []common.Address{}, []*big.Int{})
	if err != nil {
		return nil, errors.Wrap(err, "filtering the Receipt event")
	}

	level.Info(self.logger).Log("msg",
		"iterating over events",
	)
	for iter.Next() {
		select {
		case <-self.ctx.Done():
			return nil, errors.New("context cancelled")
		default:
		}

		level.Info(self.logger).Log("msg",
			"handle event", "event", iter.Event.Raw.TxHash.String(),
		)
		ctx, _ := context.WithTimeout(self.ctx, 2*time.Second)
		decimals, err := eth_bridge.GetTokenDecimals(ctx, self.client, iter.Event.Token)
		if err != nil {
			return nil, errors.Wrap(err, "getting token symbol")
		}
		ctx, _ = context.WithTimeout(self.ctx, 2*time.Second)
		symbol, err := eth_bridge.GetTokenSymbol(ctx, self.client, iter.Event.Token)
		if err != nil {
			return nil, errors.Wrap(err, "getting token symbol")
		}

		// Amount value.
		transferValue := big.NewFloat(0).SetInt(iter.Event.Amount)
		// Apply decimals.
		amount, _ := big.NewFloat(0).Quo(transferValue, big.NewFloat(math.Pow10(int(decimals)))).Float64()

		// Getting block metadata.
		// ctx, _ = context.WithTimeout(self.ctx, 10*time.Second)
		block, err := self.client.BlockByNumber(self.ctx, big.NewInt(int64(iter.Event.Raw.BlockNumber)))
		if err != nil {
			return nil, errors.Wrap(err, "getting block by number")
		}

		tx := typ.Transaction{
			Amount:     amount,
			BlockNo:    block.Header().Number.Uint64(),
			Hash:       iter.Event.Raw.TxHash.String(),
			To:         iter.Event.Recipient.String(),
			Symbol:     symbol,
			Bridge:     typ.EthereumIoteX,
			BridgeSide: typ.FromRight,
			From:       iter.Event.Sender.String(),
			Timestamp:  block.Header().Time,
		}
		txs = append(txs, tx)
	}

	return txs, nil
}
