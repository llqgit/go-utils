package go_utils

import (
	"crypto/rand"
	"math/big"
)

// 获取一个随机数 [0, max)
func GetRandomNum(max int64) int64 {
	num, _ := rand.Int(rand.Reader, big.NewInt(max))
	return num.Int64()
}
