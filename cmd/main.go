package main

import (
	"context"
	"fmt"
	stdlog "log"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/oklog/run"
	"github.com/pkg/errors"
	"github.com/polystation/polydefi-api/pkg/config"
	"github.com/polystation/polydefi-api/pkg/controller"
	"github.com/polystation/polydefi-api/pkg/db"
	"github.com/polystation/polydefi-api/pkg/logging"
)

func main() {
	logger := logging.NewLogger()
	cfg := config.GetConfig()
	ExitOnErr(godotenv.Load(), "loading .env file")
	db, err := db.OpenDB(cfg, logger)
	if err != nil {
		ExitOnErr(err, "open db")
	}
	var g run.Group
	{
		g.Add(run.SignalHandler(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM))

		// Open the TSDB database.
		tsdbOptions := tsdb.DefaultOptions()
		// tsdbOptions.RetentionDuration = int64(2 * 24 * time.Hour / time.Millisecond)
		if err := os.MkdirAll(cfg.Db.Path, 0777); err != nil {
			return errors.Wrap(err, "creating tsdb DB folder")
		}
		tsDB, err := tsdb.Open(cfg.Db.Path, nil, nil, tsdbOptions)
		if err != nil {
			return errors.Wrap(err, "creating tsdb DB")
		}
		level.Info(logger).Log("msg", "opened local db", "path", cfg.Db.Path)

		defer func() {
			if err := tsDB.Close(); err != nil {
				level.Error(logger).Log("msg", "closing the tsdb", "err", err)
			}
		}()

		// web Controller component.
		{
			controller, err := controller.NewController(cfg, db, logger)
			if err != nil {
				ExitOnErr(err, "creating controller")
			}
			g.Add(func() error {
				return controller.Start()
			}, func(error) {
				controller.Stop()
			})
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
