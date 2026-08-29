package service

import (
	"context"
	"crypto/md5"
	"math/big"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Hash(ctx context.Context, url string) string {
	hasher := md5.New()
	hasher.Write([]byte(url))
	hashBytes := hasher.Sum(nil)

	var num big.Int
	num.SetBytes(hashBytes)

	result := make([]byte, 5)
	base := big.NewInt(62)
	rem := new(big.Int)

	for i := 4; i >= 0; i-- {
		num.DivMod(&num, base, rem)
		result[i] = alphabet[rem.Int64()]
	}

	// запись в бд

	return string(result)
}
