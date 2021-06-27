package types

type Network string

const (
	NetEthereum Network = "ethereum"
)

type Transaction struct {
	From    string
	To      string
	Network Network
	Value   string
	Symbol  string
	Deposit bool
}
