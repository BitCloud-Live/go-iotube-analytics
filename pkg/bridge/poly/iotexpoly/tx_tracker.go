package iotexpoly

import (
	"context"
	"math"
	"math/big"
	"time"

	"github.com/IoTube-analytics/go-iotube-analytics/pkg/bridge"
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
	tokens map[string]bridge.ERC20
}

func NewTransactionTracker(ctx context.Context, client *ethclient.Client, logger log.Logger, cfg Config, store *bridge.Store) (*TransactionTracker, error) {
	filterLog, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	logger = log.With(filterLog, "component", ComponentName)

	// Getting tokens.
	ctxGetToken, cnclGetToken := context.WithTimeout(ctx, 10*time.Second)
	defer cnclGetToken()

	tokens, err := bridge.GetTokenListMethod2(ctxGetToken, client, logger, StandardTokenListAddress, ProxyTokenListAddress, StandardTokenListAddressStartBlockNo, ProxyTokenListAddressStartBlockNo)
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
	level.Debug(self.logger).Log("msg", "iotex tx tracker stopped")
}

func (self *TransactionTracker) Start() error {
	level.Debug(self.logger).Log("msg", "iotex tx tracker started")
	// IoTeX blocktime ticker.
	ticker := time.NewTicker(20 * time.Second)
	for {
		select {
		case <-self.ctx.Done():
			return nil
		case <-ticker.C:
		}
		var (
			fromBlockNo, toBlockNo *big.Int
		)

		// Get last checked block number from the db.
		lastCheckedBlockNo, err := self.store.LastCheckedBlockNo(typ.NetIoTeX, typ.NetPolygon)
		if lastCheckedBlockNo == nil {
			level.Error(self.logger).Log("msg", "getting lastCheckedBlockNo", "err", err)
			fromBlockNo = big.NewInt(TokenCashierStartBlockNo)
			level.Info(self.logger).Log("msg", "watching iotex blockchain for the first time")
		} else {
			// Look ahead one block to make sure we didn't miss any new invoices.
			fromBlockNo = lastCheckedBlockNo
		}
		// Calculating head block number minus 18.
		header, err := self.client.HeaderByNumber(self.ctx, nil)
		if err != nil {
			level.Error(self.logger).Log("msg", "getting latest block header", "err", err)
			continue
		}

		// Min block to loop over.
		min := math.Min(float64(fromBlockNo.Uint64()+blockLimitBeforeCommit),
			float64(header.Number.Uint64()))
		toBlockNo = big.NewInt(int64(min))

		if toBlockNo.Cmp(fromBlockNo) == -1 {
			level.Debug(self.logger).Log("msg", "no new block to check, waiting...")
			continue
		}
		level.Info(self.logger).Log("msg", "checking for new transactions",
			"fromBlockNo", fromBlockNo,
			"toBlockNo", toBlockNo,
		)
		txs, err := self.traverse(fromBlockNo, toBlockNo)
		if err != nil {
			level.Error(self.logger).Log("msg", "traversing the iotex blockchain",
				"err", err,
				"fromBlockNo", fromBlockNo,
				"toBlockNo", toBlockNo,
			)
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
			err = self.store.UpdateLastCheckedBlockNo(toBlockNo, typ.NetIoTeX, typ.NetPolygon)
			if err != nil {
				level.Error(self.logger).Log("msg", "updating iotex blockchain state",
					"err", err,
					"lastCheckedBlockNo", toBlockNo,
				)
			}
		}
	}
}

func (self *TransactionTracker) traverse(fromBlockNo, toBlockNo *big.Int) ([]typ.Transaction, error) {
	txs := make([]typ.Transaction, 0)

	tokenCashierFilterer, err := tokenCashier.NewTokenCashierFilterer(TokenCashierAddress, self.client)
	if err != nil {
		return nil, errors.Wrap(err, "getting tokenCashierFilterer")
	}
	ctx, cncl := context.WithTimeout(self.ctx, 10*time.Second)
	defer cncl()
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
			return nil, errors.New("context canceled")
		default:
		}

		level.Info(self.logger).Log("msg",
			"handle event", "event", iter.Event.Raw.TxHash.String(),
		)
		ctx, cncl := context.WithTimeout(self.ctx, 2*time.Second)
		defer cncl()
		decimals, err := bridge.GetTokenDecimals(ctx, self.client, iter.Event.Token)
		if err != nil {
			return nil, errors.Wrap(err, "getting token symbol")
		}
		ctx, cncl = context.WithTimeout(self.ctx, 2*time.Second)
		defer cncl()
		symbol, err := bridge.GetTokenSymbol(ctx, self.client, iter.Event.Token)
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
			Bridge:     typ.PolygonIoteX,
			BridgeSide: typ.FromRight,
			From:       iter.Event.Sender.String(),
			Timestamp:  block.Header().Time,
		}
		txs = append(txs, tx)
	}
	return txs, nil
}
