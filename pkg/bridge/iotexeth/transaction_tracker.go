package ethereum

import (
	"context"
	"math"
	"math/big"
	"sort"
	"strings"
	"time"

	"github.com/IoTube-analytics/go-iotube-analytics/pkg/bridge"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/contracts/erc20"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/logging"
	typ "github.com/IoTube-analytics/go-iotube-analytics/pkg/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/pkg/timestamp"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/tsdb"

	log "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type TransactionTracker struct {
	logger log.Logger
	cfg    Config
	ctx    context.Context
	cncl   context.CancelFunc
	client *ethclient.Client
	db     *tsdb.DB
	engine *promql.Engine
	// Map: token address ->  token symbol.
	tokens map[string]ERC20
}

func NewTransactionTracker(ctx context.Context, client *ethclient.Client, logger log.Logger, cfg Config, db *tsdb.DB) (*TransactionTracker, error) {
	filterLog, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	logger = log.With(filterLog, "component", ComponentName)
	ctx, cncl := context.WithCancel(ctx)

	// promqlEngine
	opts := promql.EngineOpts{
		Logger:               logger,
		Reg:                  nil,
		MaxSamples:           30000,
		Timeout:              10 * time.Second,
		LookbackDelta:        5 * time.Minute,
		EnableAtModifier:     true,
		EnableNegativeOffset: true,
	}
	engine := promql.NewEngine(opts)

	// Getting tokens.
	tokens, err := getTokenList(client)
	if err != nil {
		return nil, errors.Wrap(err, "getting token list")
	}

	return &TransactionTracker{
		logger: logger,
		cfg:    cfg,
		ctx:    ctx,
		cncl:   cncl,
		db:     db,
		engine: engine,
		client: client,
		tokens: tokens,
	}, nil
}

func (tt *TransactionTracker) Stop() {
	tt.cncl()
	level.Debug(tt.logger).Log("msg", "ethereum watcher stopped")
}

func (tt *TransactionTracker) Start() error {
	level.Debug(tt.logger).Log("msg", "ethereum watcher started")
	// Ethereum blocktime ticker.
	ticker := time.NewTicker(20 * time.Second)
	for {
		select {
		case <-tt.ctx.Done():
			return nil
		default:
		}
		var (
			fromBlockNo, toBlockNo *big.Int
		)

		// Calculating head block number minus 18.
		header, err := tt.client.HeaderByNumber(tt.ctx, nil)
		if err != nil {
			level.Error(tt.logger).Log("msg", "getting latest block header", "err", err)
			<-ticker.C
			continue
		}
		toBlockNo = new(big.Int).Sub(header.Number, big.NewInt(18))

		// Get last checked block number from the db.
		lastCheckedBlockNo, err := bridge.LastCheckedBlockNo(tt.ctx, tt.engine, tt.db, "ethereum")
		if err != nil {
			fromBlockNo = big.NewInt(TokenCashierStartBlockNo)
			level.Info(tt.logger).Log("msg", "watching ethereum blockchain for the first time")

		} else {
			// Look ahead one block to make sure we didn't miss any new invoices.
			fromBlockNo = lastCheckedBlockNo
			if toBlockNo.Uint64() < lastCheckedBlockNo.Uint64() {
				<-ticker.C
				continue
			}
		}
		level.Info(tt.logger).Log("msg", "checking for new transactions",
			"fromBlockNo", fromBlockNo,
			"toBlockNo", toBlockNo,
		)
		txs, err := tt.traverse(fromBlockNo, toBlockNo)
		if err != nil {
			level.Error(tt.logger).Log("msg", "traversing the eth blockchain",
				"err", err,
				"fromBlockNo", fromBlockNo,
				"toBlockNo", toBlockNo,
			)
			<-ticker.C
			continue
		}
		if len(txs) > 0 {
			level.Info(tt.logger).Log("msg",
				"found new transactions",
				"count", len(txs),
			)
		}
		// Lets commit txs to the database.
		// FIXME: better logic needed here.
		err = tt.recordTxs(txs)
		if err != nil {
			err = tt.updateLastCheckedBlockNo(toBlockNo)
		}
		if err != nil {
			level.Error(tt.logger).Log("msg", "updating eth blockchain state",
				"err", err,
				"fromBlockNo", fromBlockNo,
				"toBlockNo", toBlockNo,
			)
		}
		<-ticker.C
	}
}

func (tt *TransactionTracker) traverse(fromBlockNo, toBlockNo *big.Int) ([]typ.Transaction, error) {
	txs := make([]typ.Transaction, 0)

	// Iterate over blocks to 18 block before head.
	for i := fromBlockNo.Int64(); i <= toBlockNo.Int64(); i++ {
		block, err := tt.client.BlockByNumber(tt.ctx, big.NewInt(int64(i)))
		if err != nil {
			level.Error(tt.logger).Log("msg",
				"getting block by no",
				"err", err,
				"no", i)
			return nil, err
		}
		for _, tx := range block.Transactions() {
			symbol := "ETH"
			var amount float64
			if tx.To() == nil {
				// Skip on a contract creation tx.
				continue
			}
			erc20, ok := tt.tokens[tx.To().Hex()]
			if ok {
				level.Debug(tt.logger).Log("msg",
					"found transaction",
					"txHash", tx.Hash().String(),
				)
				symbol = erc20.Symbol
				amount, err = tt.getTransferAmount(erc20, tx.To(), tx.Hash())
				if err != nil {
					return nil, errors.Wrap(err, "getting transfer event")
				}
			} else {

				level.Debug(tt.logger).Log("msg",
					"eth deposit transaction",
					"txHash", tx.Hash().String(),
				)
				_amount, ok := big.NewFloat(0).SetString(tx.Value().String())
				if !ok {
					level.Error(tt.logger).Log("msg",
						"unexpected conversion error",
						"err", err,
					)
					return nil, err
				}
				amount, _ = new(big.Float).Quo(_amount, big.NewFloat(math.Pow10(18))).Float64()
			}

			msg, err := tx.AsMessage(types.NewEIP155Signer(tx.ChainId()), nil)
			if err != nil {
				level.Error(tt.logger).Log("msg",
					"getting tx from",
					"err", err,
				)
				return nil, err
			}
			tx := typ.Transaction{
				Amount:    amount,
				BlockNo:   block.NumberU64(),
				Hash:      tx.Hash().String(),
				To:        tx.To().Hex(),
				Symbol:    symbol,
				From:      msg.From().Hex(),
				Timestamp: block.Time(),
			}
			txs = append(txs, tx)
		}
	}
	return txs, nil
}

func (tt *TransactionTracker) recordTxs(txs []typ.Transaction) error {
	var err error

	appender := tt.db.Appender(tt.ctx)
	defer func() { // An appender always needs to be committed or rolled back.
		if err != nil {
			if err := appender.Rollback(); err != nil {
				level.Error(tt.logger).Log("msg", "db rollback failed", "err", err)
			}
			return
		}
		if errC := appender.Commit(); errC != nil {
			err = errors.Wrap(err, "db append commit failed")
		}
	}()

	for _, tx := range txs {
		ts := timestamp.FromFloatSeconds(float64(tx.Timestamp))
		lbls := labels.Labels{
			labels.Label{Name: "__name__", Value: "tx"},
			labels.Label{Name: "network", Value: string(tx.Network)},
			labels.Label{Name: "type", Value: "in"},
			labels.Label{Name: "symbol", Value: tx.Symbol},
		}

		sort.Sort(lbls) // This is important! The labels need to be sorted to avoid creating the same series with duplicate reference.

		_, err = appender.Append(0, lbls, ts, float64(tx.Amount))
		if err != nil {
			return errors.Wrap(err, "append values to the DB")
		}
	}
	return nil
}

func (tt *TransactionTracker) updateLastCheckedBlockNo(blockNo *big.Int) error {
	var err error
	ts := timestamp.FromTime(time.Now().Round(5 * time.Second))
	appender := tt.db.Appender(tt.ctx)
	defer func() { // An appender always needs to be committed or rolled back.
		if err != nil {
			if err := appender.Rollback(); err != nil {
				level.Error(tt.logger).Log("msg", "db rollback failed", "err", err)
			}
			return
		}
		if errC := appender.Commit(); errC != nil {
			err = errors.Wrap(err, "db append commit failed")
		}
	}()

	lbls := labels.Labels{
		labels.Label{Name: "__name__", Value: "blockchain"},
		labels.Label{Name: "network", Value: string(typ.NetEthereum)},
	}

	sort.Sort(lbls) // This is important! The labels need to be sorted to avoid creating the same series with duplicate reference.

	_, err = appender.Append(0, lbls, ts, float64(blockNo.Uint64()))
	if err != nil {
		return errors.Wrap(err, "append values to the DB")
	}
	return nil
}

func (tt *TransactionTracker) getTransferAmount(erc20Token ERC20, tokenAddress *common.Address, txHash common.Hash) (float64, error) {
	abi, err := abi.JSON(strings.NewReader(string(erc20.Erc20ABI)))
	if err != nil {
		return 0, err
	}
	tx, err := tt.client.TransactionReceipt(context.TODO(), txHash)
	if err != nil {
		return 0, err
	}
	var transferEvent struct {
		From  common.Address
		To    common.Address
		Value *big.Int
	}

	for _, vLog := range tx.Logs {
		err := abi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
		if err == nil {
			transferEvent.To = common.BytesToAddress(vLog.Topics[2].Bytes())
			if transferEvent.To == TokenSafeAddress {
				transferValue := big.NewFloat(0).SetInt(transferEvent.Value.SetBytes(vLog.Topics[3].Bytes()))
				// Apply decimals.
				amount, _ := big.NewFloat(0).Quo(transferValue, big.NewFloat(math.Pow10(int(erc20Token.Decimals)))).Float64()
				return amount, nil
			}
		}
	}
	return 0, errors.New("couldn't find a transfer")
}
