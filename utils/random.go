package utils

import (
	rand2 "crypto/rand"
	"math/big"
	"math/rand"
)

func Random(min, max int) int {
	return rand.Intn(max-min) + min
}

func Random2(min int64, max int64) int64 {
	r, _ := rand2.Int(rand2.Reader, big.NewInt(max-min+1))
	return min + r.Int64()
}

func Random3(max int64) int64 {
	r, _ := rand2.Int(rand2.Reader, big.NewInt(max+1))
	return r.Int64()
}
