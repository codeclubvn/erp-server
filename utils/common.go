package utils

import (
	"crypto/rand"
	"math/big"
	"time"
)

func GenerateCode(n int) string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return ""
		}
		ret[i] = letters[num.Int64()]
	}
	return string(ret)
}

func IsBetweenDate(startDate, endDate, date time.Time) bool {
	if date.Before(startDate) || date.After(endDate) {
		return false
	}
	return true
}
