package app

import "math/rand/v2"

func RandomStringWithCharset(length int, charset string) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.IntN(len(charset))]
	}
	return string(result)
}
