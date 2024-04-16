package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomString(length int) string {
	rand.New(rand.NewSource((time.Now().UnixNano())))

	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	randomString := make([]byte, length)

	for i := 0; i < length; i++ {
		randomString[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return string(randomString)
}
