package price

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/yalp/jsonpath"

	"net/http"
	"time"

	"github.com/IoTube-analytics/go-iotube-analytics/pkg/bridge"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/logging"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
)

// Map: symbol -> API ids.
var symbolToIds = map[string]string{
	"wbtc":   "bitcoin",
	"weth":   "ethereum",
	"uni":    "uniswap",
	"busd":   "busd",
	"usdc":   "usd-coin",
	"link":   "link",
	"usdt":   "tether",
	"iotx":   "iotex",
	"paxg":   "pax-gold",
	"cyc":    "cyclone-protocol",
	"wmatic": "matic-network",
	"sushi":  "sushi",
	"dai":    "dai",
	"aave":   "aave",
	"quick":  "quick",
	"wbnb":   "wbnb",
}
var CoinGeckoAPI = "https://api.coingecko.com/api/v3/simple/price?ids=%v&vs_currencies=usd"

const ComponentName = "price"

type Config struct {
	LogLevel string
}

// Track coin prices and add records in the influxdb.
type PriceTracker struct {
	logger log.Logger
	cfg    Config
	ctx    context.Context
	stop   context.CancelFunc
	store  *bridge.Store
}

func New(logger log.Logger, ctx context.Context, store *bridge.Store, cfg Config) (*PriceTracker, error) {
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	ctx, stop := context.WithCancel(ctx)
	return &PriceTracker{
		logger: logger,
		cfg:    cfg,
		ctx:    ctx,
		stop:   stop,
		store:  store,
	}, nil
}

func (self *PriceTracker) Start() error {
	level.Info(self.logger).Log("msg", "starting price tracker")
	// Update price every 2min.
	ticker := time.NewTicker(120 * time.Second)
	for {
		symbols, err := self.store.GetAllSymbols()
		if err != nil {
			level.Error(self.logger).Log("msg", "getting symbols", "err", err)
		}

		level.Debug(self.logger).Log("msg", "updating prices", "symbols", spew.Sdump(symbols))
		for _, symbol := range symbols {
			select {
			case <-self.ctx.Done():
				return errors.New("context canceled")
			default:
			}
			ctx, cncl := context.WithTimeout(self.ctx, 2*time.Second)
			defer cncl()
			price, err := Fetch(ctx, symbolToIds[strings.ToLower(symbol)])
			if err != nil {
				level.Error(self.logger).Log("msg", "fetching price from api", "symbol", symbol, "err", err)
				continue
			}

			level.Debug(self.logger).Log("msg", "recording price", "price", price)
			err = self.store.RecordPrice(symbol, price)
			if err != nil {
				level.Error(self.logger).Log("msg", "recording price", "err", err)
			}
		}
		select {
		case <-self.ctx.Done():
			return errors.New("context canceled")
		case <-ticker.C:
		}
	}
}

func (self *PriceTracker) Stop() {
	self.stop()
}

func Fetch(ctx context.Context, symbol string) (float64, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Transport: tr}
	ticker := time.NewTicker(1 * time.Second)

	var errFinal error
	for i := 0; i < 5; i++ {
		url := fmt.Sprintf(CoinGeckoAPI, symbol)
		r, err := client.Get(url)
		if err != nil {
			errFinal = errors.Wrap(err, "fetching data")
			select {
			case <-ticker.C:
				continue
			case <-ctx.Done():
				return 0, ctx.Err()
			}
		}

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			errFinal = errors.Wrap(err, "read response body")
			select {
			case <-ticker.C:
				continue
			case <-ctx.Done():
				return 0, ctx.Err()
			}
		}
		r.Body.Close()

		if r.StatusCode/100 != 2 {
			errFinal = errors.Errorf("response status code not OK code:%v, payload:%v", r.StatusCode, string(data))
			select {
			case <-ticker.C:
				continue
			case <-ctx.Done():
				return 0, ctx.Err()
			}
		}
		var inputToParse interface{}

		err = json.Unmarshal(data, &inputToParse)
		if err != nil {
			return 0, errors.Wrapf(err, "json marshal:%v", string(data))
		}

		output, err := jsonpath.Read(inputToParse, fmt.Sprintf(`$["%v"].usd`, symbol))
		if err != nil {
			return 0, errors.Wrapf(err, "json path read:%v", string(data))
		}

		return parseInterface(output)
	}

	return 0, errFinal

}

func parseInterface(data interface{}) (float64, error) {
	// Expect result to be a slice of float or a single float value.
	var resultList []interface{}
	switch result := data.(type) {
	case []interface{}:
		resultList = result
	default:
		resultList = []interface{}{result}
	}
	// Parse each item of slice to a float.
	var value float64
	for i, a := range resultList {
		strValue := fmt.Sprintf("%v", a)
		// Normalize based on american locale.
		strValue = strings.Replace(strValue, ",", "", -1)

		switch i {
		case 0:
			val, err := strconv.ParseFloat(strValue, 64)
			if err != nil {
				return 0, errors.Wrapf(err, "value needs to be a valid float:%v", strValue)
			}
			value = val
		case 1:

		}
	}

	return value, nil
}
