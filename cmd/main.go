package main

import (
	"context"
	"fmt"
	stdlog "log"
	"os"
	"syscall"

	"github.com/IoTube-analytics/go-iotube-analytics/pkg/bridge"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/bridge/bsc/bsciotex"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/bridge/bsc/iotexbsc"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/bridge/eth/ethiotex"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/bridge/eth/iotexeth"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/bridge/poly/iotexpoly"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/bridge/poly/polyiotex"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/config"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/ethereum"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/logging"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/price"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/web"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kit/kit/log/level"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/oklog/run"
	"github.com/pkg/errors"
)

func main() {
	logger := logging.NewLogger()

	cfg, err := config.ParseConfig(logger, "")
	if err != nil {
		ExitOnErr(err, "creating config")

	}
	globalCtx := context.Background()

	// Ethereum client.
	client, err := ethereum.NewClient(globalCtx, logger)
	if err != nil {
		ExitOnErr(err, "creating client")

	}

	// IoTeX babel api client.
	babelClient, err := ethclient.DialContext(globalCtx, os.Getenv(iotexeth.NodeUrlKey))
	if err != nil {
		ExitOnErr(err, "creating iotex client")

	}

	// Polygon api client.
	polygonClient, err := ethclient.DialContext(globalCtx, os.Getenv(polyiotex.NodeUrlKey))
	if err != nil {
		ExitOnErr(err, "creating polygon client")

	}

	// Bsc api client.
	bscClient, err := ethclient.DialContext(globalCtx, os.Getenv(bsciotex.NodeUrlKey))
	if err != nil {
		ExitOnErr(err, "creating polygon client")

	}

	// Influxdb client.
	tsdb := influxdb2.NewClient(os.Getenv("INFLUXDB_URL"), os.Getenv("INFLUXDB_TOKEN"))
	// always close client at the end
	defer client.Close()
	if err != nil {
		ExitOnErr(err, "creating influxdb client")

	}

	var g run.Group
	{
		g.Add(run.SignalHandler(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM))
		store, err := bridge.NewSore(globalCtx, logger, cfg.Bridge, tsdb)
		if err != nil {
			ExitOnErr(err, "creating bridge store")
		}

		// web api component.
		web, err := web.New(logger, globalCtx, tsdb, cfg.Web)
		if err != nil {
			ExitOnErr(err, "creating web controller")
		}
		g.Add(func() error {
			return web.Start()
		}, func(error) {
			web.Stop()
		})

		// Price tracker component.
		price, err := price.New(logger, globalCtx, store, cfg.Price)
		if err != nil {
			ExitOnErr(err, "creating price tracker")
		}
		g.Add(func() error {
			return price.Start()
		}, func(error) {
			price.Stop()
		})

		// Ethereum bridge
		{
			// Ethereum part.
			{
				{
					// ethereum tx tracker.
					ethTXTracker, err := ethiotex.NewTransactionTracker(globalCtx, client, logger, cfg.EthIoTeX, store)
					if err != nil {
						ExitOnErr(err, "creating ethTXTracker")
					}
					g.Add(func() error {
						level.Info(logger).Log("msg", "ethiotex tx tracker started")
						return ethTXTracker.Start()
					}, func(error) {
						ethTXTracker.Stop()
						level.Info(logger).Log("msg", "ethiotex tx tracker shutdown complete")

					})

					// ethereum tvl tracker.
					if true {
						ethTVLTracker, err := ethiotex.NewTVLTracker(globalCtx, client, logger, cfg.EthIoTeX, store)
						if err != nil {
							ExitOnErr(err, "creating ethTVLTracker")
						}
						g.Add(func() error {
							level.Info(logger).Log("msg", "ethiotex tvl tracker started")
							return ethTVLTracker.Start()
						}, func(error) {
							ethTVLTracker.Stop()
							level.Info(logger).Log("msg", "ethiotex tvl tracker shutdown complete")

						})
					}
				}
			}
			// iotex part.
			{
				// ethereum tx tracker.
				iotexEthTXTracker, err := iotexeth.NewTransactionTracker(globalCtx, babelClient, logger, cfg.IoTeXEth, store)
				if err != nil {
					ExitOnErr(err, "creating iotexEthTXTracker")
				}
				g.Add(func() error {

					level.Info(logger).Log("msg", "iotexeth tx tracker started")
					return iotexEthTXTracker.Start()
				}, func(error) {
					iotexEthTXTracker.Stop()
					level.Info(logger).Log("msg", "iotexeth tx tracker shutdown complete")
				})

			}
		}

		// Polygon bridge
		{
			// Polygon part.
			{
				{
					// Polygon tx tracker.
					polyTXTracker, err := polyiotex.NewTransactionTracker(globalCtx, polygonClient, logger, cfg.PolyIoTeX, store)
					if err != nil {
						ExitOnErr(err, "creating polyTXTracker")
					}
					g.Add(func() error {
						level.Info(logger).Log("msg", "polyiotex tx tracker started")
						return polyTXTracker.Start()
					}, func(error) {
						polyTXTracker.Stop()
						level.Info(logger).Log("msg", "polyiotex tx tracker shutdown complete")
					})

					// Polygon tvl tracker.
					if true {
						polyTVLTracker, err := polyiotex.NewTVLTracker(globalCtx, polygonClient, logger, cfg.PolyIoTeX, store)
						if err != nil {
							ExitOnErr(err, "creating polyTVLTracker")
						}
						g.Add(func() error {

							level.Info(logger).Log("msg", "polyiotex tvl tracker started")
							return polyTVLTracker.Start()
						}, func(error) {
							polyTVLTracker.Stop()
							level.Info(logger).Log("msg", "polyiotex tvl tracker shutdown complete")
						})
					}
				}
			}
			// iotex part.
			{
				// ethereum tx tracker.
				iotexPolyTXTracker, err := iotexpoly.NewTransactionTracker(globalCtx, babelClient, logger, cfg.IoTeXPoly, store)
				if err != nil {
					ExitOnErr(err, "creating iotexPolyTXTracker")
				}
				g.Add(func() error {
					level.Info(logger).Log("msg", "iotexpoly tx tracker started")
					return iotexPolyTXTracker.Start()
				}, func(error) {
					iotexPolyTXTracker.Stop()
					level.Info(logger).Log("msg", "iotexpoly tx tracker shutdown complete")
				})
			}
		}

		// Bsc bridge
		{
			// BSC part.
			{
				{
					// BSC tx tracker.
					bscTXTracker, err := bsciotex.NewTransactionTracker(globalCtx, bscClient, logger, cfg.BscIoTeX, store)
					if err != nil {
						ExitOnErr(err, "creating bscTXTracker")
					}
					g.Add(func() error {
						level.Info(logger).Log("msg", "bsciotex tx tracker started")
						return bscTXTracker.Start()
					}, func(error) {
						bscTXTracker.Stop()
						level.Info(logger).Log("msg", "bsciotex tx tracker shutdown complete")
					})

					// Bsc tvl tracker.
					{
						bscTVLTracker, err := bsciotex.NewTVLTracker(globalCtx, bscClient, logger, cfg.BscIoTeX, store)
						if err != nil {
							ExitOnErr(err, "creating bscTVLTracker")
						}
						g.Add(func() error {

							level.Info(logger).Log("msg", "bsciotex tvl tracker started")
							return bscTVLTracker.Start()
						}, func(error) {
							bscTVLTracker.Stop()
							level.Info(logger).Log("msg", "bsciotex tvl tracker shutdown complete")
						})
					}
				}
			}
			// iotex part.
			{
				// ethereum tx tracker.
				iotexBscTXTracker, err := iotexbsc.NewTransactionTracker(globalCtx, babelClient, logger, cfg.IoTeXBsc, store)
				if err != nil {
					ExitOnErr(err, "creating iotexBscTXTracker")
				}
				g.Add(func() error {
					level.Info(logger).Log("msg", "iotexbsc tx tracker started")
					return iotexBscTXTracker.Start()
				}, func(error) {
					iotexBscTXTracker.Stop()
					level.Info(logger).Log("msg", "iotexbsc tx tracker shutdown complete")
				})
			}
		}
	}

	if err := g.Run(); err != nil {
		stdlog.Println(fmt.Sprintf("%+v", errors.Wrapf(err, "run group stacktrace")))
	}

}

func ExitOnErr(err error, msg string) {
	if err != nil {
		stdlog.Fatalf("root execution error:%+v msg:%+v", err, msg)
	}
}
