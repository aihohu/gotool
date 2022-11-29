package randomUtil

import (
	"crypto/rand"
	"math/big"
	mrand "math/rand"
)

const letterNumber = "0123456789"
const letterRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomLimitInt 获得指定范围内的随机数 [0,limit)
// @param limit 限制随机数的范围，不包括这个数
// @return 随机数
func RandomLimitInt(limit int64) int64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(limit))
	return n.Int64()
}

// RandomInt 指定范围内的随机数
// @param min 最小数（包含）
// @param max 最大数（不包含）
// @return 指定范围内的随机数
func RandomInt(min, max int64) int64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(max-min))
	return n.Int64() + min
}

// RandomString 随机指定长度字符串（只包含数字和字符）
// @param length 随机字符串长度
// @return 随机字符串
func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterRunes[mrand.Intn(len(letterRunes))]
	}
	return string(b)
}

// RandomNumbers 只包含数字的随机字符串
// @param length 随机字符串长度
// @return 只包含数字的随机字符串
func RandomNumbers(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterNumber[mrand.Intn(len(letterNumber))]
	}
	return string(b)
}
