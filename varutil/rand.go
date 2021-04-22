package varutil

import (
	"math/rand"
	"time"
)

const (
	// NumericBytes is set of numeric characters for RandString function
	NumericBytes = "1234567890"
	// LowerAlphaBytes is set of lower alphabetic characters for RandString function
	LowerAlphaBytes = "abcdefghijklmnopqrstuvwxyz"
	// UpperAlphaBytes is set of upper alphabetic characters for RandString function
	UpperAlphaBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// AlphaBytes is set of alphabetic characters for RandString function
	AlphaBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// AlphaNumericBytes is set of numeric and alphabetic characters for RandString function
	AlphaNumericBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	// LowerAlphaNumericBytes is set of numeric and only lower alphabetic characters for RandString function
	LowerAlphaNumericBytes = "abcdefghijklmnopqrstuvwxyz1234567890"
	// UpperAlphaNumericBytes is set of numeric and only upper alphabetic characters for RandString function
	UpperAlphaNumericBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	// StrongBytes is set of numeric, alphabetic and special characters for RandString function
	StrongBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()`~<>?:[]{}-=_+|\\/,."

	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var (
	src = rand.NewSource(time.Now().UnixNano())
)

// RandString create new random string
func RandString(n int, pool string) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(pool) {
			b[i] = pool[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}
