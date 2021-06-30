package ethiotex

import (
	"context"
	"time"

	"github.com/IoTube-analytics/go-iotube-analytics/pkg/bridge"
	eth_bridge "github.com/IoTube-analytics/go-iotube-analytics/pkg/bridge/eth"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/logging"
	typ "github.com/IoTube-analytics/go-iotube-analytics/pkg/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	log "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type TVLTracker struct {
	logger log.Logger
	cfg    Config
	ctx    context.Context
	cncl   context.CancelFunc
	client *ethclient.Client
	store  *bridge.Store
	// Map: token address ->  token symbol.
	tokens map[string]eth_bridge.ERC20
}

func NewTVLTracker(ctx context.Context, client *ethclient.Client, logger log.Logger, cfg Config, store *bridge.Store) (*TVLTracker, error) {
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
	ctx, cncl := context.WithCancel(ctx)
	return &TVLTracker{
		logger: logger,
		cfg:    cfg,
		ctx:    ctx,
		cncl:   cncl,
		store:  store,

		client: client,
		tokens: tokens,
	}, nil
}

func (self *TVLTracker) Stop() {
	self.cncl()
	level.Debug(self.logger).Log("msg", "tvl tracker stopped")
}

func (self *TVLTracker) Start() error {
	level.Debug(self.logger).Log("msg", "tvl tracker started")

	// Ethereum blocktime ticker.
	// Update tvl every 10 minutes by default.
	ticker := time.NewTicker(10 * time.Minute)
	for {
		tvlData := make([]typ.TVLData, 0)
		for addr, erc20 := range self.tokens {
			//ctx, _ := context.WithTimeout(self.ctx, 10*time.Second)
			tvl, err := eth_bridge.GetTVL(self.ctx, self.client, common.HexToAddress(addr), TokenSafeAddress)
			if err != nil {
				level.Error(self.logger).Log("msg", "getting tvl", "token", erc20.Symbol, "err", err)
			}
			tvlData = append(tvlData, typ.TVLData{
				Value:   tvl,
				Network: typ.NetEthereum,
				Symbol:  erc20.Symbol,
			})

		}
		err := self.store.UpdateTVL(tvlData)
		if err != nil {
			level.Error(self.logger).Log("msg", "recording tvl", "err", err)
		}
		level.Info(self.logger).Log("msg", "tvl data updated")
		select {
		case <-self.ctx.Done():
			return nil
		case <-ticker.C:
		}
	}
}
