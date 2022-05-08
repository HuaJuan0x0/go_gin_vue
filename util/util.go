package util

import (
	"math/rand"
	"time"
)

func RandomString(i int) string {
	var letters = []byte("qwertyuuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	result := make([]byte, i)
	rand.Seed(time.Now().Unix())
	for n := range result {
		result[n] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}
