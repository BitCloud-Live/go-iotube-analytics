package main

import (
	"context"
	"fmt"
	stdlog "log"
	"os"
	"syscall"

	"github.com/IoTube-analytics/go-iotube-analytics/pkg/bridge"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/bridge/eth/ethiotex"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/bridge/eth/iotexeth"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/config"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/ethereum"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/logging"
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
	babelClient, err := ethclient.DialContext(globalCtx, iotexeth.BabelHost)
	if err != nil {
		ExitOnErr(err, "creating client")

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

		// Ethereum bridge
		{
			// Ethereum part.
			{
				if true {
					// ethereum tx tracker.
					ethTXTracker, err := ethiotex.NewTransactionTracker(globalCtx, client, logger, cfg.EthIoTeX, store)
					if err != nil {
						ExitOnErr(err, "creating ethTXTracker")
					}
					g.Add(func() error {
						ethTXTracker.Start()
						level.Info(logger).Log("msg", "ethiotex tx tracker shutdown complete")
						return nil
					}, func(error) {
						ethTXTracker.Stop()
					})

					// ethereum tvl tracker.
					if true {
						ethTVLTracker, err := ethiotex.NewTVLTracker(globalCtx, client, logger, cfg.EthIoTeX, store)
						if err != nil {
							ExitOnErr(err, "creating ethTVLTracker")
						}
						g.Add(func() error {
							ethTVLTracker.Start()
							level.Info(logger).Log("msg", "ethiotex tvl tracker shutdown complete")
							return nil
						}, func(error) {
							ethTVLTracker.Stop()
						})
					}
				}
			}
			// iotex part.
			if true {
				// ethereum tx tracker.
				iotexTXTracker, err := iotexeth.NewTransactionTracker(globalCtx, babelClient, logger, cfg.IoTeXEth, store)
				if err != nil {
					ExitOnErr(err, "creating ethBridgeTXTracker")
				}
				g.Add(func() error {
					iotexTXTracker.Start()
					level.Info(logger).Log("msg", "iotexeth tx tracker shutdown complete")
					return nil
				}, func(error) {
					iotexTXTracker.Stop()
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
