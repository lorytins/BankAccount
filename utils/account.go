package utils

import (
	"math/rand"
	"time"
)

func GenerateAccountNumber() int {
	rand.Seed(time.Now().UnixNano())
	min := 100000
	max := 999999
	return rand.Intn(max-min) + min
}

func GenerateBranchNumber() int {
	rand.Seed(time.Now().UnixNano())
	min := 1000
	max := 9999
	return rand.Intn(max-min) + min
}