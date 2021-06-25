package web

import "github.com/IoTube-analytics/go-iotube-analytics/pkg/format"

// import (
// 	"github.com/IoTube-analytics/go-iotube-analytics/pkg/db"
// 	"github.com/IoTube-analytics/go-iotube-analytics/pkg/format"
// 	"github.com/IoTube-analytics/go-iotube-analytics/pkg/logging"
// 	"github.com/IoTube-analytics/go-iotube-analytics/pkg/openapi/swagger/models"
// 	"github.com/IoTube-analytics/go-iotube-analytics/pkg/openapi/swagger/restapi"
// 	"github.com/IoTube-analytics/go-iotube-analytics/pkg/openapi/swagger/restapi/operations"
// 	"github.com/IoTube-analytics/go-iotube-analytics/pkg/openapi/swagger/restapi/operations/data"
// 	"github.com/go-kit/kit/log"
// 	"github.com/go-kit/kit/log/level"
// 	"github.com/go-openapi/loads"
// 	"github.com/go-openapi/runtime/middleware"
// 	"github.com/pkg/errors"
// 	"github.com/prometheus/tsdb"
// )

const ComponentName = "web"

type Config struct {
	LogLevel    string
	ListenHost  string
	ListenPort  uint
	ReadTimeout format.Duration
}

// type resp struct {
// 	Code    int
// 	Message string
// }

// type Web struct {
// 	db     *tsdb.DB
// 	cfg    Config
// 	logger log.Logger
// 	api    *operations.PolydefiAPI
// 	server *restapi.Server
// }

// func New(cfg *Config, db db.DB, logger log.Logger) (*Web, error) {
// 	// Creating the component logger.
// 	filterLog, err := logging.ApplyFilter(*cfg, ComponentName, logger)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "apply filter logger")
// 	}
// 	logger = log.With(filterLog, "component", ComponentName)
// 	api, err := newApi()
// 	if err != nil {
// 		return nil, err
// 	}
// 	c := &Web{db: db,
// 		logger: logger,
// 		api:    api,
// 		cfg:    cfg,
// 	}
// 	// Register the api.
// 	c.register(api)
// 	c.server = newServer(cfg, api)
// 	return c, nil
// }

// func (c *Web) register(api *operations.PolydefiAPI) {
// 	api.DataGetAllDataHandler = data.GetAllDataHandlerFunc(c.GetAllData)
// 	api.DataGetChartDataHandler = data.GetChartDataHandlerFunc(c.GetChartData)
// }

// func newApi() (*operations.PolydefiAPI, error) {
// 	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return operations.NewPolydefiAPI(swaggerSpec), nil

// }
// func newServer(cfg *config.Config, api *operations.PolydefiAPI) *restapi.Server {
// 	server := restapi.NewServer(api)
// 	server.ConfigureAPI()
// 	server.EnabledListeners = []string{"http"}
// 	server.Port = cfg.GetDefaultInt(config.PORT, 9876)
// 	return server
// }

// func (c *Web) Start() error {
// 	return c.server.Serve()
// }

// func (c *Web) Stop() {
// 	level.Debug(c.logger).Log("msg", "shutting down the controller")
// 	c.server.Shutdown()
// }

// func (c *Web) GetAllData(params data.GetAllDataParams) middleware.Responder {
// 	level.Debug(c.logger).Log("msg", "getting all data")
// 	ds, err := c.db.GetLatestDefiData()
// 	if err != nil {
// 		level.Error(c.logger).Log("msg", "getting latest defi data", "err", err)
// 		return data.NewGetAllDataNotFound().WithPayload(&models.APIResponse{
// 			Message: "defi data not found",
// 			Status:  "error",
// 		})
// 	}

// 	return data.NewGetAllDataOK().WithPayload(convert(ds))
// }

// func (c *Web) GetChartData(params data.GetChartDataParams) middleware.Responder {
// 	level.Debug(c.logger).Log("msg", "getting all chart data")
// 	chartData, err := c.db.GetChartData(params.Days)
// 	if err != nil {
// 		level.Error(c.logger).Log("msg", "getting all chart data from db", "err", err)

// 		return data.NewGetAllDataNotFound().WithPayload(&models.APIResponse{
// 			Message: "chart data not found",
// 			Status:  "error",
// 		})
// 	}

// 	return data.NewGetChartDataOK().WithPayload(chartData)
// }

// func convert(in []db.DefiData) models.AllData {
// 	m := make(models.AllData, 0)
// 	for _, i := range in {
// 		m = append(m, &models.DefiData{
// 			Category:            i.Category,
// 			Chain:               i.Chain,
// 			ContractNum:         i.ContractNum,
// 			Holders:             i.Holders,
// 			HoldersChange24hNum: i.HoldersChange24hNum,
// 			LastUpdated:         i.CreatedAt.Unix(),
// 			LockedUsd:           i.LockedUsd,
// 			MarketCap:           i.MarketCap,
// 			MarketCapChange24h:  i.MarketCapChange24h,
// 			Name:                i.Name,
// 			Price:               i.Price,
// 			PriceChange24h:      i.PriceChange24h,
// 			Token:               i.Token,
// 			TvlPercentChange24h: i.TvlPercentChange24h,
// 			Verified:            i.Verified,
// 			Volume:              i.Volume,
// 		})
// 	}
// 	return m
// }
