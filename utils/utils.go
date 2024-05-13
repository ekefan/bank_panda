package utils

import (
	"math/rand"
	"time"
	"strings"
)

var alphabets = "abcdefghijklmnopkrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

//Random int generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func randomString(n int) string {
	var sb  strings.Builder
	lenAlphabets := len(alphabets)
	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(lenAlphabets)]
		sb.WriteByte(c)
	}
	return sb.String()
}

//Util  to generate random account owner name
func RandomOwner() string {
	return randomString(5)
}

func RandomBalance() int64 {
	return RandomInt(200, 1000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	lenCurrencies := len(currencies)
	return currencies[rand.Intn(lenCurrencies)] 
}

