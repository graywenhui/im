package util

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var letters = []byte("asdfgasjkdajslfaskdjasldSKDALSDKAJLSDKJASDLJASLDKALSH")
	result := make([]byte, n)

	rand.Seed(time.Now().UnixNano())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}
