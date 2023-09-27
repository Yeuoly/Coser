package math

import (
	"crypto/rand"
	"math/big"
)

func Random(min int, max int) int {
	res, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	return min + int(res.Int64())
}

func RandomString(len int) string {
	var result = ""
	for i := 0; i < len; i++ {
		result += string(byte(Random(97, 122)))
	}
	return result
}

func RandomBytes(len int) []byte {
	var result = make([]byte, len)
	for i := 0; i < len; i++ {
		result[i] = byte(Random(0, 255))
	}
	return result
}
