package ethereum

import "github.com/ethereum/go-ethereum/common"

const ComponentName = "ethereum"

var TokenCashierAddress = common.HexToAddress("0xa0fd7430852361931b23a31f84374ba3314e1682")
var TokenSafeAddress = common.HexToAddress("0xc2e0f31d739cb3153ba5760a203b3bd7c27f0d7a")
var StandardTokenListAddress = common.HexToAddress("0x7c0bef36e1b1cbeb1f1a5541300786a7b608aede")
var ProxyTokenListAddress = common.HexToAddress("0x73ffdfc98983ad59fb441fc5fe855c1589e35b3e")

type Config struct {
	LogLevel string
	Timeout  uint
}
