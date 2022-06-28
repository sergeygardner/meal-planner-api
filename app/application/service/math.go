package service

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

func MathRandomInt(min int64, max int64) int64 {
	bigInt, errorBigInt := rand.Int(rand.Reader, big.NewInt(max))

	if errorBigInt != nil {
		panic(errorBigInt)
	}

	return bigInt.Int64() + min
}

func MathRandomIntAsString(min int64, max int64) string {
	return strconv.FormatInt(MathRandomInt(min, max), 10)
}
