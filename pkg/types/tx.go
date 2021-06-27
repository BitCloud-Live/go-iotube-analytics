package types

type Network string

const (
	NetEthereum Network = "ethereum"
)

type Transaction struct {
	From      string
	To        string
	Hash      string
	BlockNo   uint64
	Network   Network
	Amount    float64
	Symbol    string
	Deposit   bool
	Timestamp uint64
}
