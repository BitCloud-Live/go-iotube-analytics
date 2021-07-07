package polyiotex

import "github.com/ethereum/go-ethereum/common"

const ComponentName = "polyiotex"
const NodeUrlKey = "POLYGON_NODE_URL"

var TokenCashierAddress = common.HexToAddress("0xf72CFb704d49aC7BB7FFa420AE5f084C671A29be")
var TokenSafeAddress = common.HexToAddress("0xA239F03Cda98A7d2AaAA51e7bF408e5d73399e45")
var StandardTokenListAddress = common.HexToAddress("0xDe9395d2f4940aA501f9a27B98592589D14Bb0f7")
var MintableTokenListAddress = common.HexToAddress("0xC8DC8dCDFd94f9Cb953f379a7aD8Da5fdC303F3E")

const TokenCashierStartBlockNo = 15316068
const TokenSafeStartBlockNo = 15254714

// We will track at most `blockLimitBeforeCommit` block before save the tx data to the db.
const blockLimitBeforeCommit = uint64(999)

type Config struct {
	LogLevel string
	Timeout  uint
}
