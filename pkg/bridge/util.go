package bridge

import (
	"context"
	"math"
	"math/big"

	"strings"

	"github.com/IoTube-analytics/go-iotube-analytics/pkg/contracts/erc20"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/contracts/tokenList"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/go-kit/kit/log"
	"github.com/pkg/errors"
)

// CanonicalSymbolName unifies symbol names.
func CanonicalSymbolName(symbol string) string {
	symbol = strings.TrimPrefix(symbol, "io")
	if symbol == "WETH" {
		symbol = "ETH"
	}
	return symbol
}

type ERC20 struct {
	Symbol   string
	Decimals uint8
}

// getTokenList gathers a map of: token address -> token symbol.
func GetTokenList(ctx context.Context, client *ethclient.Client, logger log.Logger, standardTokenListAddress, proxyTokenListAddress common.Address) (map[string]ERC20, error) {
	out := make(map[string]ERC20)
	// Getting standard token list.
	tokenListCaller, err := tokenList.NewTokenListCaller(standardTokenListAddress, client)
	if err != nil {
		return nil, errors.Wrap(err, "getting token list caller")
	}
	count, err := tokenListCaller.Count(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "getting standard token count")
	}
	standardTokens, err := tokenListCaller.GetActiveItems(&bind.CallOpts{}, big.NewInt(0), uint8(count.Uint64()))
	if err != nil {
		return nil, errors.Wrap(err, "getting standard token")
	}
	// Getting proxy token list.
	proxyTokenListCaller, err := tokenList.NewTokenListCaller(proxyTokenListAddress, client)
	if err != nil {
		return nil, errors.Wrap(err, "getting proxy token caller")
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
	count, err = tokenListCaller.Count(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "getting proxy token count")
	}
	if count.Uint64() != 0 {
		proxyTokens, err := proxyTokenListCaller.GetActiveItems(&bind.CallOpts{}, big.NewInt(0), uint8(count.Uint64()))
		if err != nil {
			// return nil, errors.Wrap(err, "getting proxy token")
			// Skipping on proxy tokens.
			return out, nil
		}
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
