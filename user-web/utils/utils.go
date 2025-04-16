package utils

import (
	"math/rand"
	"time"
)

func GenerateNumericCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	const digits = "0123456789"
	var code string
	for i := 0; i < length; i++ {
		code += string(digits[rand.Intn(len(digits))])
	}
	return code
}
