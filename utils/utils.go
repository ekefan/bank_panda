package utils

import (
	"math/rand"
	"time"
	"fmt"
	"strings"
)

var alphabets = "abcdefghijklmnopkrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano()) // rand.Seed set the current seed to ensure different generated values
}


//Random Email generates a random email like address for testing
func RandomEmail() string {
	return fmt.Sprintf("%v@email.com", RandomString(4))
}
//Random int generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
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
	return RandomString(5)
}

func RandomBalance() int64 {
	return RandomInt(200, 1000)
}

func RandomCurrency() string {
	currencies := []string{USD, CAD, EUR}
	lenCurrencies := len(currencies)
	return currencies[rand.Intn(lenCurrencies)] 
}

