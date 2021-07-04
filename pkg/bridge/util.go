package bridge

import "strings"

// CanonicalSymbolName unifies symbol names.
func CanonicalSymbolName(symbol string) string {
	symbol = strings.TrimPrefix(symbol, "io")
	if symbol == "WETH" {
		symbol = "ETH"
	}
	return symbol
}
