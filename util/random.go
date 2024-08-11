package util

import "math/rand"

func RandomInteger(min, max int) int {
	return min + rand.Intn(max - min + 1) // min -> n
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))] // n -> len(letters)
	}
	return string(s)
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return int64(RandomInteger(0, 1000))
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)] // n -> len(currencies)
}