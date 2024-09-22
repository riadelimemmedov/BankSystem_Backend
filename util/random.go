package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// ! Init generate random numbers from unix time
func init() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(r.Uint32())
}

// ! RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// ! RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	random_string := sb.String()
	if len(random_string) > 0 {
		first_char := unicode.ToUpper(rune(random_string[0]))
		return string(first_char) + random_string[1:]
	}
	return random_string
}

// ! RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// ! RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// ! RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "GBP", "JPY", "CAD", "TL", "AZN", "LR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// ! RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(6))
}
