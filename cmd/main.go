package main

import (
	"context"
	"fmt"
	stdlog "log"
	"os"
	"syscall"

	"github.com/IoTube-analytics/go-iotube-analytics/pkg/bridge/iotexeth"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/bridge/iotexeth/eth"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/config"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/logging"
	"github.com/go-kit/kit/log/level"
	"github.com/oklog/run"
	"github.com/pkg/errors"

	"github.com/prometheus/prometheus/tsdb"
)

func main() {
	logger := logging.NewLogger()

	cfg, err := config.ParseConfig(logger, "")
	if err != nil {
		ExitOnErr(err, "creating config")

	}
	globalCtx := context.Background()

	client, err := eth.NewClient(globalCtx, logger)
	if err != nil {
		ExitOnErr(err, "creating client")

	}

	var g run.Group
	{
		g.Add(run.SignalHandler(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM))

		// Open the TSDB database.
		tsdbOptions := tsdb.DefaultOptions()
		// tsdbOptions.RetentionDuration = int64(2 * 24 * time.Hour / time.Millisecond)
		if err := os.MkdirAll(cfg.Db.Path, 0777); err != nil {
			ExitOnErr(err, "creating tsdb DB folder")
		}
		tsDB, err := tsdb.Open(cfg.Db.Path, nil, nil, tsdbOptions)
		if err != nil {
			ExitOnErr(err, "creating tsdb DB")
		}
		level.Info(logger).Log("msg", "opened local db", "path", cfg.Db.Path)

		defer func() {
			if err := tsDB.Close(); err != nil {
				level.Error(logger).Log("msg", "closing the tsdb", "err", err)
			}
		}()

		// web Controller component.
		/*
			controller, err := web.New(cfg, db, logger)
			if err != nil {
				ExitOnErr(err, "creating controller")
			}
			g.Add(func() error {
				return controller.Start()
			}, func(error) {
				controller.Stop()
			})
		*/
		// ethereum <-> iotex bridge.
		ethBridgeTXTracker, err := iotexeth.NewTransactionTracker(globalCtx, client, logger, cfg.EthereumBridge, tsDB)
		if err != nil {
			ExitOnErr(err, "creating tsdb DB")
		}
		g.Add(func() error {
			ethBridgeTXTracker.Start()
			level.Info(logger).Log("msg", "iotexeth tx tracker shutdown complete")
			return nil
		}, func(error) {
			ethBridgeTXTracker.Stop()
		})
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
