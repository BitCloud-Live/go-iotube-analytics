package eth_bridge

import (
	"context"
	"math"
	"math/big"

	"github.com/IoTube-analytics/go-iotube-analytics/pkg/contracts/erc20"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/contracts/tokenList"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/go-kit/kit/log"
	"github.com/pkg/errors"
)

type ERC20 struct {
	Symbol   string
	Decimals uint8
}

// getTokenList gathers a map of: token address -> token symbol.
func GetTokenList(ctx context.Context, client *ethclient.Client, logger log.Logger, standardTokenListAddress, proxyTokenListAddress common.Address) (map[string]ERC20, error) {
	// Getting standard token list.
	tokenListCaller, err := tokenList.NewTokenListCaller(standardTokenListAddress, client)
	if err != nil {
		return nil, err
	}
	standardTokens, err := tokenListCaller.GetActiveItems(&bind.CallOpts{}, big.NewInt(0), 10)
	if err != nil {
		return nil, err
	}
	// Getting proxy token list.
	proxyTokenListCaller, err := tokenList.NewTokenListCaller(proxyTokenListAddress, client)
	if err != nil {
		return nil, err
	}
	proxyTokens, err := proxyTokenListCaller.GetActiveItems(&bind.CallOpts{}, big.NewInt(0), 10)
	if err != nil {
		return nil, err
	}
	out := make(map[string]ERC20)
	for _, t := range proxyTokens.Items {
		// Skip on zero address!
		if t == common.BigToAddress(big.NewInt(0)) {
			continue
		}
		symbol, err := GetTokenSymbol(ctx, client, t)
		if err != nil {
			return nil, errors.Wrap(err, "can't fetch token symbol")
		}
		decimals, err := GetTokenDecimals(ctx, client, t)
		if err != nil {
			return nil, errors.Wrap(err, "can't fetch token decimals")
		}
		out[t.Hash().Hex()] = ERC20{Symbol: symbol, Decimals: decimals}
	}
	for _, t := range standardTokens.Items {
		// Skip on zero address!
		if t == common.BigToAddress(big.NewInt(0)) {
			continue
		}
		symbol, err := GetTokenSymbol(ctx, client, t)
		if err != nil {
			return nil, errors.Wrap(err, "can't fetch token symbol")
		}
		decimals, err := GetTokenDecimals(ctx, client, t)
		if err != nil {
			return nil, errors.Wrap(err, "can't fetch token decimals")
		}
		out[t.Hash().Hex()] = ERC20{Symbol: symbol, Decimals: decimals}
	}
	return out, nil

}

func GetTVL(ctx context.Context, client *ethclient.Client, tokenAddress, tokenSafeAddress common.Address) (float64, error) {
	// Getting standard token list.
	erc20Caller, err := erc20.NewErc20Caller(tokenAddress, client)
	if err != nil {
		return 0, err

	}
	balance, err := erc20Caller.BalanceOf(&bind.CallOpts{Context: ctx}, tokenSafeAddress)
	if err != nil {
		return 0, errors.Wrap(err, "can't fetch token balance")
	}
	decimals, err := GetTokenDecimals(ctx, client, tokenAddress)
	if err != nil {
		return 0, errors.Wrap(err, "can't fetch token decimals")
	}
	transferValue := big.NewFloat(0).SetInt(balance)
	// Apply decimals.
	amount, _ := big.NewFloat(0).Quo(transferValue, big.NewFloat(math.Pow10(int(decimals)))).Float64()
	return amount, nil
}

func GetTokenSymbol(ctx context.Context, client *ethclient.Client, token common.Address) (string, error) {
	// Getting token symbol.
	erc20Caller, err := erc20.NewErc20Caller(token, client)
	if err != nil {
		return "", err

	}
	return erc20Caller.Symbol(&bind.CallOpts{})
}

func GetTokenDecimals(ctx context.Context, client *ethclient.Client, token common.Address) (uint8, error) {
	// Getting token decimals.
	erc20Caller, err := erc20.NewErc20Caller(token, client)
	if err != nil {
		return 0, err

	}
	return erc20Caller.Decimals(&bind.CallOpts{})
}
