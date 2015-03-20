package shortener

import (
	"math/rand"
	"regexp"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func IsURL(str string) bool {
	if str == "" || len(str) > 2082 {
		return false
	}
	var rxURL = regexp.MustCompile("^(https?://)?([0-9a-z.-]+).([a-z.]{2,6})([/a-z .-]*)*/?$")
	return rxURL.MatchString(str)
}
