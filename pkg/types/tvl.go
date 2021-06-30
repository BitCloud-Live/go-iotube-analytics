package types

type Network string

const (
	NetEthereum Network = "ethereum"
	NetIoTeX    Network = "iotex"
)

type TVLData struct {
	Value   float64
	Network Network
	Symbol  string
}
