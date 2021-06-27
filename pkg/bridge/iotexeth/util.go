package ethereum

import (
	"math/big"

	"github.com/IoTube-analytics/go-iotube-analytics/pkg/contracts/erc20"
	"github.com/IoTube-analytics/go-iotube-analytics/pkg/contracts/tokenList"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getTokenList(client *ethclient.Client) ([]common.Address, error) {
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
	for _, t := range proxyTokens.Items {
		standardTokens.Items = append(standardTokens.Items, t)
	}
	return standardTokens.Items, nil

}

func getTVL(client *ethclient.Client, token common.Address) (*big.Int, error) {
	// Getting standard token list.
	erc20Caller, err := erc20.NewErc20Caller(token, client)
	if err != nil {
		return nil, err

	}
	return erc20Caller.BalanceOf(&bind.CallOpts{}, TokenSafeAddress)
}
