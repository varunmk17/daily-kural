package utils

import (
	"math/rand"
	"time"
)

func StringArrayContains(array []string, element string) bool {
	if array == nil || element == "" {
		return false
	}

	for _, val := range array {
		if val == element {
			return true
		}
	}

	return false
}

func RandomNumber() int {
	const MAX_KURAL_NUMBER = 1330

	source := rand.NewSource(time.Now().UnixNano())
	randomNumberGenerator := rand.New(source)
	randomNumber := randomNumberGenerator.Intn(MAX_KURAL_NUMBER) + 1

	return randomNumber
}
