package iotexeth

import "github.com/ethereum/go-ethereum/common"

const ComponentName = "iotexeth"

var TokenCashierAddressIoTeX = "io1gsr52ahqzklaf7flqar8r0269f2utkw9349qg8"
var TokenCashierAddress = common.HexToAddress("0x44074576e015bfd4f93f074671bd5a2a55c5d9c5")
var TokenSafeAddressIoTeX = "io1cj3f498390srqv765nnvaxuk0rpxyadzpfjz75"
var TokenSafeAddress = common.HexToAddress("0xc4a29a94f12be03033daa4e6ce9b9678c26275a2")
var StandardTokenListAddressIoTeX = "io1t89whrwyfr0supctsqcx9n7ex5dd8yusfqhyfz"
var StandardTokenListAddress = common.HexToAddress("0x59caeb8dc448df0e070b803062cfd9351ad39390")
var ProxyTokenListAddressIoTeX = "io1dn8nqk3pmmll990xz6a94fpradtrljxmmx5p8j"
var ProxyTokenListAddress = common.HexToAddress("0x6ccf305a21defff295e616ba5aa423eb563fc8db")

// First deposit to the iotex blockchain.
const TokenCashierStartBlockNo = 9529096
const TokenSafeStartBlockNo = 9509443

const NodeUrlKey = "IOTEX_BABEL_URL"

// We will track at most `ethBridgeTVLTracker` block before save the tx data to the tsdb.
const blockLimitBeforeCommit = uint64(10000)

type Config struct {
	LogLevel string
	Timeout  uint
}
