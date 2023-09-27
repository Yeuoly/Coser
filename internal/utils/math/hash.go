package math

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func Md5(src string) string {
	md5ctx := md5.New()
	md5ctx.Write([]byte(src))
	return hex.EncodeToString(md5ctx.Sum(nil))
}

func Sha1(src string) string {
	sha1ctx := sha1.New()
	sha1ctx.Write([]byte(src))
	return hex.EncodeToString(sha1ctx.Sum(nil))
}

func Sha256(src string) string {
	sha256ctx := sha256.New()
	sha256ctx.Write([]byte(src))
	return hex.EncodeToString(sha256ctx.Sum(nil))
}

func Sha512(src string) string {
	sha512ctx := sha512.New()
	sha512ctx.Write([]byte(src))
	return hex.EncodeToString(sha512ctx.Sum(nil))
}

func Hash2Int(src string) int {
	result := 0
	for _, v := range src {
		if v >= '0' && v <= '9' {
			result = result*16 + int(v-'0')
		} else if v >= 'a' && v <= 'b' {
			result = result*16 + int(v-'a'+10)
		}
	}

	return result
}
