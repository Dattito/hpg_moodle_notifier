package utils

import (
	"math/rand"
	"time"
)

func GenerateVerificationCode() int {
	rand.Seed(time.Now().UnixNano() + rand.Int63())
	return 111111 + rand.Intn(888888)
}
