package iotexbsc

import "github.com/ethereum/go-ethereum/common"

const ComponentName = "iotexeth"

var TokenCashierAddress = common.HexToAddress("0x14bf347a597aac623240ae7ac8383ae198966277")

// First deposit to the iotex blockchain.
const TokenCashierStartBlockNo = 9780237

// We will track at most `blockLimitBeforeCommit` block before save the tx data to the tsdb.
const blockLimitBeforeCommit = uint64(10000)

type Config struct {
	LogLevel string
	Timeout  uint
}
