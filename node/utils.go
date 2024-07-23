package node

import (
	"math/big"
	"strconv"
)

func Uint64ToHex(v uint64) string {
	return "0x" + strconv.FormatUint(v, 16)
}

func BigIntToHex(v *big.Int) string {
	return "0x" + v.Text(16)
}

func StringToHex(s string) string {
	return "0x" + s
}
