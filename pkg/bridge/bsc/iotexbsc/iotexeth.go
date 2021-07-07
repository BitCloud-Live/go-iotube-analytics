package iotexeth

import "github.com/ethereum/go-ethereum/common"

const ComponentName = "iotexeth"

var TokenCashierAddressIoTeX = "io1zjlng7je02kxyvjq4eavswp6uxvfvcnh2a0a3d"
var TokenCashierAddress = common.HexToAddress("0x14bf347a597aac623240ae7ac8383ae198966277")
var StandardTokenListAddressIoTeX = "io1h2d3r0d20t58sv6h707ppc959kvs8wjsurrtnk"
var StandardTokenListAddress = common.HexToAddress("0xba9b11bdaa7ae8783357f3fc10e0b42d9903ba50")
var ProxyTokenListAddressIoTeX = "io17r9ehjstwj4gfqzwpm08fjnd606h04h2m6r92f"
var ProxyTokenListAddress = common.HexToAddress("0xf0cb9bca0b74aa84804e0ede74ca6dd3f577d6ea")

// First deposit to the iotex blockchain.
const TokenCashierStartBlockNo = 9780237

const BabelHost = "https://babel-api.mainnet.iotex.io"

// We will track at most `ethBridgeTVLTracker` block before save the tx data to the tsdb.
const blockLimitBeforeCommit = uint64(10000)

type Config struct {
	LogLevel string
	Timeout  uint
}
