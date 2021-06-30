package bridge

import (
	"context"

	"math/big"
	"time"

	"github.com/IoTube-analytics/go-iotube-analytics/pkg/logging"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/types"
	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

const ComponentName = "store"

type Config struct {
	LogLevel string
	Timeout  uint
}
type Store struct {
	ctx      context.Context
	tsdb     influxdb2.Client
	writeAPI api.WriteAPIBlocking
	readAPI  api.QueryAPI
	logger   log.Logger
}

func NewSore(ctx context.Context, logger log.Logger, cfg Config, tsdb influxdb2.Client) (*Store, error) {
	filterLog, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	logger = log.With(filterLog, "component", ComponentName)
	writeAPI := tsdb.WriteAPIBlocking("my-org", "my-bucket")
	readAPI := tsdb.QueryAPI("my-org")
	return &Store{
		tsdb:     tsdb,
		writeAPI: writeAPI,
		readAPI:  readAPI,
		ctx:      ctx,
		logger:   logger,
	}, nil
}

// LastCheckedBlockNo returns last checked block number.
func (self *Store) LastCheckedBlockNo(network types.Network) (*big.Int, error) {
	// Get parser flux query result
	query := `from(bucket: "my-bucket")
	|> range(start: -10d)
	|> filter(fn: (r) => r["_measurement"] == "blockchain")
	|> filter(fn: (r) => r["_field"] == "block_number")
	|> filter(fn: (r) => r["network"] == ` + "\"" + string(network) + "\"" + `)
	|> last()`
	result, err := self.readAPI.Query(context.Background(), query)

	var record interface{}
	if err == nil {
		// Use Next() to iterate over query result lines
		if result.Next() {
			// read result
			record = result.Record().Value()
			return big.NewInt(int64(record.(uint64))), nil
		}
		if result.Err() != nil {
			return nil, result.Err()
		}
	}
	return nil, err
}

func (self *Store) RecordTxs(txs []types.Transaction) error {
	for _, tx := range txs {
		// Create point using fluent style.
		p := influxdb2.NewPointWithMeasurement("tx").
			AddTag("bridge", string(tx.Bridge)).
			AddTag("bridge_side", string(tx.BridgeSide)).
			AddTag("symbol", string(tx.Symbol)).
			AddField("amount", tx.Amount).
			SetTime(time.Unix(int64(tx.Timestamp), 0))
		err := self.writeAPI.WritePoint(context.Background(), p)
		if err != nil {
			return err
		}
	}
	return nil
}

func (self *Store) UpdateLastCheckedBlockNo(blockNo *big.Int, net types.Network) error {
	// Create point using fluent style
	p := influxdb2.NewPointWithMeasurement("blockchain").
		AddTag("network", string(net)).
		AddField("block_number", blockNo.Uint64()).
		SetTime(time.Now())
	err := self.writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		return err
	}
	return nil
}

func (self *Store) UpdateTVL(tvls []types.TVLData) error {
	for _, tvl := range tvls {
		// Create point using fluent style
		p := influxdb2.NewPointWithMeasurement("tvl").
			AddTag("network", string(tvl.Network)).
			AddTag("symbol", string(tvl.Symbol)).
			AddField("tvl", tvl.Value).
			SetTime(time.Now())
		err := self.writeAPI.WritePoint(context.Background(), p)
		if err != nil {
			return err
		}
	}
	return nil
}
