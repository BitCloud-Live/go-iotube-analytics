package bsciotex

import "github.com/ethereum/go-ethereum/common"

const ComponentName = "bsciotex"

const NodeUrlKey = "BSC_NODE_URL"

var TokenCashierAddress = common.HexToAddress("0x797f1465796fd89ea7135e76dbc7cdb136bba1ca")
var TokenSafeAddress = common.HexToAddress("0xFBe9A4138AFDF1fA639a8c2818a0C4513fc4CE4B")
var StandardTokenListAddress = common.HexToAddress("0x0d793F4D4287265B9bdA86b7a4083193E8743b34")
var MintableTokenListAddress = common.HexToAddress("0xa6ae9312D0AA3CC74d969Fcd4806d7729A321EE3")

const TokenCashierStartBlockNo = 5179731
const TokenSafeStartBlockNo = 5179717

// We will track at most `blockLimitBeforeCommit` block before save the tx data to the tsdb.
const blockLimitBeforeCommit = uint64(4999)

type Config struct {
	LogLevel string
	Timeout  uint
}
