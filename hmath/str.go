package hmath

import (
	"fmt"
	"math/big"
)

func Base10ToBase35(numStr string) (string, error) {
	n := new(big.Int)
	_, success := n.SetString(numStr, 10)
	if !success {
		return "", fmt.Errorf("invalid number string: %s", numStr)
	}

	const charset = "0123456789ABCDEFGHIJKLMNPQRSTUVWXYZ"

	if n.Sign() == 0 {
		return "0", nil
	}

	base := big.NewInt(35)
	zero := big.NewInt(0)
	mod := new(big.Int)

	result := make([]byte, 0, len(numStr)/3+1)

	for n.Cmp(zero) > 0 {
		n.DivMod(n, base, mod)
		result = append(result, charset[mod.Int64()])
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result), nil
}
