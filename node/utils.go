package node

import (
	"math/big"
	"strconv"
)

// Uint64ToHex converts a uint64 to a hex string
func Uint64ToHex(v uint64) string {
	return "0x" + strconv.FormatUint(v, 16)
}

// BigIntToHex converts a big.Int to a hex string
func BigIntToHex(v *big.Int) string {
	return "0x" + v.Text(16)
}

// StringToHex converts a string to a hex string
func StringToHex(s string) string {
	return "0x" + s
}
