package types

type Bridge string

const (
	EthereumIoteX Bridge = "ethereum"
)

type BridgeSide string

const (
	FromLeft  BridgeSide = "left"
	FromRight BridgeSide = "right"
)

type Transaction struct {
	From       string
	To         string
	Hash       string
	BlockNo    uint64
	Bridge     Bridge
	Amount     float64
	BridgeSide BridgeSide
	Symbol     string
	Deposit    bool
	Timestamp  uint64
}
