package ethereum

import (
	"context"
	"math"
	"math/big"
	"sort"
	"time"

	"github.com/IoTube-analytics/go-iotube-analytics/pkg/bridge"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/logging"
	typ "github.com/IoTube-analytics/go-iotube-analytics/pkg/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/tsdb"
	"github.com/tellor-io/telliot/pkg/format"

	log "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/shopspring/decimal"
)

type EthereumWatcher struct {
	logger log.Logger
	cfg    Config
	ctx    context.Context
	cncl   context.CancelFunc
	client *ethclient.Client
	db     *tsdb.DB
	engine *promql.Engine
	tokens []common.Address
}

func NewEthereumWatcher(ctx context.Context, client *ethclient.Client, logger log.Logger, cfg Config, db *tsdb.DB) (*EthereumWatcher, error) {
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

	return &EthereumWatcher{
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

func (ew *EthereumWatcher) Stop() {
	ew.cncl()
	level.Debug(ew.logger).Log("msg", "ethereum watcher stopped")
}

func (ew *EthereumWatcher) Start() error {
	level.Debug(ew.logger).Log("msg", "ethereum watcher started")
	// Ethereum blocktime ticker.
	ticker := time.NewTicker(20 * time.Second)
	for {
		select {
		case <-ew.ctx.Done():
			return nil
		default:
		}
		var (
			fromBlockNo, toBlockNo *big.Int
		)

		// Calculating head block number minus 18.
		header, err := ew.client.HeaderByNumber(ew.ctx, nil)
		if err != nil {
			level.Error(ew.logger).Log("msg", "getting latest block header", "err", err)
			<-ticker.C
			continue
		}
		toBlockNo = new(big.Int).Sub(header.Number, big.NewInt(18))

		// Get last checked block number from the db.
		lastCheckedBlockNo, err := bridge.LastCheckedBlockNo(ew.ctx, ew.engine, ew.db, "ethereum")
		if err != nil {
			fromBlockNo = new(big.Int).Sub(toBlockNo, big.NewInt(18))
			level.Info(ew.logger).Log("msg", "watching ethereum blockchain for the first time")

		} else {
			// Look ahead one block to make sure we didn't miss any new invoices.
			fromBlockNo = lastCheckedBlockNo
			if toBlockNo.Uint64() < lastCheckedBlockNo.Uint64() {
				<-ticker.C
				continue
			}
		}
		level.Info(ew.logger).Log("msg", "checking for new transactions",
			"fromBlockNo", fromBlockNo,
			"toBlockNo", toBlockNo,
		)
		txs, err := ew.traverse(fromBlockNo, toBlockNo)
		if err != nil {
			level.Error(ew.logger).Log("msg", "traversing the eth blockchain",
				"err", err,
				"fromBlockNo", fromBlockNo,
				"toBlockNo", toBlockNo,
			)
			<-ticker.C
			continue
		}
		if len(txs) > 0 {
			level.Info(ew.logger).Log("msg",
				"found new transactions",
				"count", len(txs),
			)
		}
		// Lets commit txs to the database.
		self.updateState(txs, toBlockNo)
		if err != nil {
			level.Error(ew.logger).Log("msg", "updating eth blockchain state",
				"err", err,
				"fromBlockNo", fromBlockNo,
				"toBlockNo", toBlockNo,
			)
		}
		<-ticker.C
	}
}

func (ew *EthereumWatcher) traverse(fromBlockNo, toBlockNo *big.Int) ([]typ.Transaction, error) {
	txs := make([]typ.Transaction, 0)

	// Iterate over blocks to 18 block before head.
	for i := fromBlockNo.Int64(); i <= toBlockNo.Int64(); i++ {
		block, err := ew.client.BlockByNumber(ew.ctx, big.NewInt(int64(i)))
		if err != nil {
			level.Error(ew.logger).Log("msg",
				"getting block by no",
				"err", err,
				"no", i)
			return nil, err
		}
		for _, tx := range block.Transactions() {
			if tx.To() == nil {
				// Skip on a contract creation tx.
				continue
			}
			address := tx.To().Hex()
			if _, ok := unsettled[address]; !ok {
				continue
			}
			level.Info(ew.logger).Log("msg",
				"found transaction",
				"txHash", tx.Hash().String(),
			)

			amount := new(big.Float)
			amount.SetString(tx.Value().String())
			ethValue := new(big.Float).Quo(amount, big.NewFloat(math.Pow10(18)))
			ethDecimal, err := decimal.NewFromString(ethValue.String())
			if err != nil {
				level.Error(ew.logger).Log("msg",
					"convert big float to decimal",
					"err", err,
				)
				return nil, err
			}

			msg, err := tx.AsMessage(types.NewEIP155Signer(tx.ChainId()))
			if err != nil {
				level.Error(ew.logger).Log("msg",
					"getting tx from",
					"err", err,
				)
				return nil, err
			}
			coin, _ := wallet.ETH.Symbol()
			tx := db.Transaction{
				InvoiceID: unsettled[address],
				Amount:    ethDecimal,
				BlockHash: block.Hash().String(),
				BlockNo:   block.NumberU64(),
				TxHash:    tx.Hash().String(),
				Coin:      coin,
				From:      msg.From().Hex(),
			}
			txs = append(txs, tx)
		}
	}
	// metrics.AddDeposits(trade_types.ETHUSDT, trade_types.Sell, totalDeposits, float64(len(deposits)))
	// if err != nil {
	// 	//	log.Fatal("we coudn't commit deposits to the system")
	// }
	// log.Printf("successfuly add new deposits to the system")
	return txs, nil
}

func (ew *EthereumWatcher) updateState(tx *types.Transaction) error {
	appender := ew.db.Appender(self.ctx)
	defer func() { // An appender always needs to be committed or rolled back.
		if err != nil {
			if err := appender.Rollback(); err != nil {
				level.Error(logger).Log("msg", "db rollback failed", "err", err)
			}
			return
		}
		if errC := appender.Commit(); errC != nil {
			err = errors.Wrap(err, "db append commit failed")
		}
	}()

	lbls := labels.Labels{
		labels.Label{Name: "__name__", Value: IntervalMetricName},
		labels.Label{Name: "source", Value: dataSource.Source()},
		labels.Label{Name: "domain", Value: source.Host},
		labels.Label{Name: "symbol", Value: format.SanitizeMetricName(symbol)},
	}

	sort.Sort(lbls) // This is important! The labels need to be sorted to avoid creating the same series with duplicate reference.

	_, err = appender.Append(0, lbls, ts, float64(interval))
	if err != nil {
		return errors.Wrap(err, "append values to the DB")
	}
}

func (ew *EthereumWatcher) recordTx(tx *types.Transaction) error {
	appender := ew.db.Appender(self.ctx)
	defer func() { // An appender always needs to be committed or rolled back.
		if err != nil {
			if err := appender.Rollback(); err != nil {
				level.Error(logger).Log("msg", "db rollback failed", "err", err)
			}
			return
		}
		if errC := appender.Commit(); errC != nil {
			err = errors.Wrap(err, "db append commit failed")
		}
	}()

	lbls := labels.Labels{
		labels.Label{Name: "__name__", Value: IntervalMetricName},
		labels.Label{Name: "source", Value: dataSource.Source()},
		labels.Label{Name: "domain", Value: source.Host},
		labels.Label{Name: "symbol", Value: format.SanitizeMetricName(symbol)},
	}

	sort.Sort(lbls) // This is important! The labels need to be sorted to avoid creating the same series with duplicate reference.

	_, err = appender.Append(0, lbls, ts, float64(interval))
	if err != nil {
		return errors.Wrap(err, "append values to the DB")
	}
}
