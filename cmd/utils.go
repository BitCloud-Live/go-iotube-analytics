package main

import (
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/polystation/polydefi-api/pkg/config"
)

func CreateBTCClient(cfg *config.Config) (*rpcclient.Client, error) {
	// Connect to local bitcoin core RPC server using HTTP POST mode.
	rpcURL := cfg.GetString(config.BTC_NODE_URL)
	rpcUser := cfg.GetString(config.BTC_NODE_USER)
	rpcPass := cfg.GetString(config.BTC_NODE_PASS)
	connCfg := &rpcclient.ConnConfig{
		Host:         rpcURL,
		User:         rpcUser,
		Pass:         rpcPass,
		HTTPPostMode: true,  // Bitcoin core only supports HTTP POST mode
		DisableTLS:   false, // Bitcoin core does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	return rpcclient.New(connCfg, nil)
}
