package service

import (
	"crypto/md5"
	"math/big"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Hash(url string) string {
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

	return string(result)
}
