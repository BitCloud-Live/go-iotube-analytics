package main

import (
	"context"
	"log"
	"math/big"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-antenna-go/v2/account"
	"github.com/iotexproject/iotex-antenna-go/v2/iotex"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
)

const (
	host = "api.testnet.iotex.one:443"
)

func main() {
	// Create grpc connection
	conn, err := iotex.NewDefaultGRPCConn(host)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Add account by private key
	acc, err := account.HexStringToAccount("...")
	if err != nil {
		log.Fatal(err)
	}

	// create client
	c := iotex.NewAuthedClient(iotexapi.NewAPIServiceClient(conn), acc)

	// transfer
	to, err := address.FromString("to...")
	if err != nil {
		log.Fatal(err)
	}
	hash, err := c.Transfer(to, big.NewInt(10)).Call(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
