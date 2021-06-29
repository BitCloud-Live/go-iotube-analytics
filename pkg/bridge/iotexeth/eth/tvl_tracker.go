package iotexeth

import (
	"context"
	"sort"
	"time"

	"github.com/IoTube-analytics/go-iotube-analytics/pkg/logging"
	typ "github.com/IoTube-analytics/go-iotube-analytics/pkg/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/pkg/timestamp"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/tsdb"

	log "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type TVLTracker struct {
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

func NewTVLTracker(ctx context.Context, client *ethclient.Client, logger log.Logger, cfg Config, db *tsdb.DB) (*TVLTracker, error) {
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
	tokens, err := getTokenList(client, logger)
	if err != nil {
		return nil, errors.Wrap(err, "getting token list")
	}

	return &TVLTracker{
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

func (ew *TVLTracker) Stop() {
	ew.cncl()
	level.Debug(ew.logger).Log("msg", "tvl tracker stopped")
}

func (ew *TVLTracker) Start() error {
	level.Debug(ew.logger).Log("msg", "tvl tracker started")

	// Ethereum blocktime ticker.
	// Update tvl every 10 minutes by default.
	ticker := time.NewTicker(10 * time.Minute)
	for {
		tvlData := make([]typ.TVLData, 0)
		for addr, erc20 := range ew.tokens {
			tvl, err := getTVL(ew.client, common.HexToAddress(addr))
			if err != nil {
				level.Error(ew.logger).Log("msg", "getting tvl", "token", erc20.Symbol)
			}
			tvlData = append(tvlData, typ.TVLData{
				Value:   tvl,
				Network: typ.NetEthereum,
				Symbol:  erc20.Symbol,
			})

		}
		err := ew.updateTVL(tvlData)
		if err != nil {
			level.Error(ew.logger).Log("msg", "recording tvl", "err", err)
		}
		level.Info(ew.logger).Log("msg", "tvl data updated")
		select {
		case <-ew.ctx.Done():
			return nil
		case <-ticker.C:
			return nil
		}
	}
}

func (ew *TVLTracker) updateTVL(tvls []typ.TVLData) error {
	var err error

	appender := ew.db.Appender(ew.ctx)
	defer func() { // An appender always needs to be committed or rolled back.
		if err != nil {
			if err := appender.Rollback(); err != nil {
				level.Error(ew.logger).Log("msg", "db rollback failed", "err", err)
			}
			return
		}
		if errC := appender.Commit(); errC != nil {
			err = errors.Wrap(err, "db append commit failed")
		}
	}()

	for _, tvl := range tvls {
		ts := timestamp.FromTime(time.Now().Round(5 * time.Second))
		lbls := labels.Labels{
			labels.Label{Name: "__name__", Value: "tvl"},
			labels.Label{Name: "network", Value: string(tvl.Network)},
			labels.Label{Name: "symbol", Value: tvl.Symbol},
		}

		sort.Sort(lbls) // This is important! The labels need to be sorted to avoid creating the same series with duplicate reference.

		_, err = appender.Append(0, lbls, ts, tvl.Value)
		if err != nil {
			return errors.Wrap(err, "append values to the DB")
		}
	}
	return nil
}
