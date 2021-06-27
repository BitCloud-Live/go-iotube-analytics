package ethereum

import (
	"math"
	"math/big"

	"github.com/IoTube-analytics/go-iotube-analytics/pkg/contracts/erc20"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/contracts/tokenList"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type ERC20 struct {
	Symbol   string
	Decimals uint8
}

// getTokenList gathers a map of: token address -> token symbol.
func getTokenList(client *ethclient.Client) (map[string]ERC20, error) {
	// Getting standard token list.
	tokenListCaller, err := tokenList.NewTokenListCaller(StandardTokenListAddress, client)
	if err != nil {
		return nil, err
	}
	standardTokens, err := tokenListCaller.GetActiveItems(&bind.CallOpts{}, nil, 0)
	if err != nil {
		return nil, err
	}
	// Getting proxy token list.
	proxyTokenListCaller, err := tokenList.NewTokenListCaller(ProxyTokenListAddress, client)
	if err != nil {
		return nil, err
	}
	proxyTokens, err := proxyTokenListCaller.GetActiveItems(&bind.CallOpts{}, nil, 0)
	if err != nil {
		return nil, err
	}
	out := make(map[string]ERC20)
	for _, t := range proxyTokens.Items {
		symbol, err := getTokenSymbol(client, t)
		if err != nil {
			return nil, errors.Wrap(err, "can't fetch token symbol")
		}
		decimals, err := getTokenDecimals(client, t)
		if err != nil {
			return nil, errors.Wrap(err, "can't fetch token decimals")
		}
		out[t.Hash().Hex()] = ERC20{Symbol: symbol, Decimals: decimals}
	}
	for _, t := range standardTokens.Items {
		symbol, err := getTokenSymbol(client, t)
		if err != nil {
			return nil, errors.Wrap(err, "can't fetch token symbol")
		}
		decimals, err := getTokenDecimals(client, t)
		if err != nil {
			return nil, errors.Wrap(err, "can't fetch token decimals")
		}
		out[t.Hash().Hex()] = ERC20{Symbol: symbol, Decimals: decimals}
	}
	return out, nil

}

func getTVL(client *ethclient.Client, token common.Address) (float64, error) {
	// Getting standard token list.
	erc20Caller, err := erc20.NewErc20Caller(token, client)
	if err != nil {
		return 0, err

	}
	balance, err := erc20Caller.BalanceOf(&bind.CallOpts{}, TokenSafeAddress)
	if err != nil {
		return 0, errors.Wrap(err, "can't fetch token balance")
	}
	decimals, err := getTokenDecimals(client, token)
	if err != nil {
		return 0, errors.Wrap(err, "can't fetch token decimals")
	}
	transferValue := big.NewFloat(0).SetInt(balance)
	// Apply decimals.
	amount, _ := big.NewFloat(0).Quo(transferValue, big.NewFloat(math.Pow10(int(decimals)))).Float64()
	return amount, nil
}

func getTokenSymbol(client *ethclient.Client, token common.Address) (string, error) {
	// Getting token symbol.
	erc20Caller, err := erc20.NewErc20Caller(token, client)
	if err != nil {
		return "", err

	}
	return erc20Caller.Symbol(&bind.CallOpts{})
}

func getTokenDecimals(client *ethclient.Client, token common.Address) (uint8, error) {
	// Getting token decimals.
	erc20Caller, err := erc20.NewErc20Caller(token, client)
	if err != nil {
		return 0, err

	}
	return erc20Caller.Decimals(&bind.CallOpts{})
}
