package types

type Network string

const (
	NetEthereum Network = "ethereum"
	NetIoTeX    Network = "iotex"
	NetPolygon  Network = "polygon"
	NetBsc      Network = "bsc"
)

type TVLData struct {
	Value   float64
	Network Network
	Symbol  string
}
