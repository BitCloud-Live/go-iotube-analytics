package config

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/IoTube-analytics/go-iotube-analytics/pkg/bridge"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/bridge/eth/ethiotex"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/bridge/eth/iotexeth"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/db"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/format"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/web"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

// Config is the top-level configuration that holds configs for all components.

type Config struct {
	Web      web.Config
	EthIoTeX ethiotex.Config
	IoTeXEth iotexeth.Config
	Db       db.Config
	Bridge   bridge.Config
	// EnvFile location that include all private details like private key etc.
	EnvFile string `json:"envFile"`
}

var DefaultConfig = Config{
	Web: web.Config{
		LogLevel:    "info",
		ListenHost:  "", // Listen on all addresses.
		ListenPort:  9090,
		ReadTimeout: format.Duration{Duration: 10 * time.Second},
	},
	Db: db.Config{
		LogLevel:      "info",
		Path:          "db",
		RemoteTimeout: format.Duration{Duration: 5 * time.Second},
	},
	EthIoTeX: ethiotex.Config{
		LogLevel: "info",
		Timeout:  3000,
	},
	IoTeXEth: iotexeth.Config{
		LogLevel: "info",
		Timeout:  3000,
	},
	Bridge: bridge.Config{
		LogLevel: "info",
		Timeout:  3000,
	},
	EnvFile: ".env",
}

func ParseConfig(logger log.Logger, path string) (*Config, error) {
	if path == "" {
		path = filepath.Join("configs", "config.json")
	}

	cfg := &Config{}
	// DeepCopy the default config into the final config.
	{
		b, err := json.Marshal(DefaultConfig)
		if err != nil {
			return nil, errors.Wrap(err, "marshal default config")
		}

		if err := json.Unmarshal(b, cfg); err != nil {
			return nil, errors.Wrap(err, "copy default config")
		}
	}

	f, err := os.Open(path)
	var noConfigFile bool
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, errors.Wrap(err, "open config file")
		}
		noConfigFile = true
		level.Warn(logger).Log("msg", "no config file on disk so using defaults", "path", path)
	}

	if !noConfigFile {
		dec := json.NewDecoder(f)
		dec.DisallowUnknownFields()
		for {
			// Override defaults with the custom configs.
			if err := dec.Decode(cfg); err == io.EOF {
				break
			} else if err != nil {
				return nil, errors.Wrap(err, "parse config")
			}

		}
	}

	if err := godotenv.Load(cfg.EnvFile); err != nil && !os.IsNotExist(err) {
		return nil, errors.Wrap(err, "loading env vars from env file")
	}

	return cfg, nil
}
