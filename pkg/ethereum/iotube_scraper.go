package ethereum

import (
	"context"
	"math"
	"math/big"
	"time"

	"github.com/pkg/errors"

	log "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/shopspring/decimal"
)

const ComponentName = "ethWatcher"

type EthereumWatcher struct {
	logger log.Logger
	cfg    *config.Config
	ctx    context.Context
	cncl   context.CancelFunc
	client *ethclient.Client
	db     db.DB
}

func NewEthereumWatcher(ctx context.Context, client *ethclient.Client, logger log.Logger, cfg *config.Config, db db.DB) (*EthereumWatcher, error) {
	filterLog, err := logging.ApplyFilter(*cfg, ComponentName, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	logger = log.With(filterLog, "component", ComponentName)
	ctx, cncl := context.WithCancel(ctx)
	return &EthereumWatcher{
		logger: logger,
		cfg:    cfg,
		ctx:    ctx,
		cncl:   cncl,
		db:     db,
		client: client,
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

		// Get last checked block no from the db.
		coin, _ := wallet.ETH.Symbol()
		lastCheckedBlockNo, err := ew.db.GetLastCheckedBlockNo(coin)

		if err != nil {
			fromBlockNo = new(big.Int).Sub(toBlockNo, big.NewInt(18))
			level.Info(ew.logger).Log("msg", "watching ethereum blockchain for the first time")

		} else {
			// Look ahead one block to make sure we didn't miss any new invoices.
			fromBlockNo = big.NewInt(int64(lastCheckedBlockNo))
			if toBlockNo.Uint64() < lastCheckedBlockNo {
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
		err = ew.db.UpdateState(wallet.ETH, txs, toBlockNo)
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

func (ew *EthereumWatcher) traverse(fromBlockNo, toBlockNo *big.Int) ([]db.Transaction, error) {
	unsettled, err := ew.db.GetUnsettledInvoices(wallet.ETH)
	if err != nil {
		level.Error(ew.logger).Log("msg",
			"loading unsettled invoice from the kv",
			"err", err)
		return nil, err
	}
	txs := []db.Transaction{}
	if len(unsettled) == 0 {
		level.Info(ew.logger).Log("msg",
			"there is no unsettled invoice, skipping on-chain calls",
		)
		return txs, nil

	}

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
