package iotexpoly

import "github.com/ethereum/go-ethereum/common"

const ComponentName = "iotexpoly"

var TokenCashierAddress = common.HexToAddress("0x540a92dd951407ee6c94b997a43ecf30ea6d04cd")
var StandardTokenListAddress = common.HexToAddress("0x2F8768cD292E94A0Da78671974B89B87a398356E")
var ProxyTokenListAddress = common.HexToAddress("0xD757adFF0eC4060e2c4A15f9777767f5Ca738Ca9")

// First deposit to the iotex blockchain.
const TokenCashierStartBlockNo = 11426143
const StandardTokenListAddressStartBlockNo = 11426024
const ProxyTokenListAddressStartBlockNo = 11461992

// We will track at most `ethBridgeTVLTracker` block before save the tx data to the tsdb.
const blockLimitBeforeCommit = uint64(10000)

type Config struct {
	LogLevel string
	Timeout  uint
}
