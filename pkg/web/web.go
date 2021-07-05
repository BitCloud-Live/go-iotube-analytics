// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package web

import (
	"context"
	"fmt"
	"net/http"

	"github.com/IoTube-analytics/go-iotube-analytics/pkg/format"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/logging"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/web/api"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/pkg/errors"
	"github.com/prometheus/common/route"
	"github.com/rs/cors"
)

const ComponentName = "web"

type Config struct {
	LogLevel    string
	ListenHost  string
	ListenPort  uint
	ReadTimeout format.Duration
}

type Web struct {
	logger log.Logger
	cfg    Config
	ctx    context.Context
	stop   context.CancelFunc
	srv    *http.Server
}

func New(logger log.Logger, ctx context.Context, tsDB influxdb2.Client, cfg Config) (*Web, error) {
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	router := route.New()

	api := api.New(logger, ctx, tsDB)
	api.Register(router.WithPrefix("/api/v1"))

	mux := http.NewServeMux()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	mux.Handle("/", c.Handler(router))
	srv := &http.Server{
		Handler:     mux,
		ReadTimeout: cfg.ReadTimeout.Duration,
		Addr:        fmt.Sprintf("%s:%d", cfg.ListenHost, cfg.ListenPort),
	}

	ctx, stop := context.WithCancel(ctx)

	return &Web{
		logger: log.With(logger, "component", ComponentName),
		cfg:    cfg,
		ctx:    ctx,
		stop:   stop,
		srv:    srv,
	}, nil

}

func (self *Web) Start() error {
	level.Info(self.logger).Log("msg", "starting", "addr", self.srv.Addr)
	if err := self.srv.ListenAndServe(); err != http.ErrServerClosed {
		return errors.Wrapf(err, "ListenAndServe")
	}
	return nil
}

func (self *Web) Stop() {
	self.stop()
	if err := self.srv.Close(); err != nil {
		level.Error(self.logger).Log("msg", "closing srv", "err", err)
	}
}
