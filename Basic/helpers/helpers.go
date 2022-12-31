package helpers

import (
	"math/rand"
	"time"
)

type SomeType struct {
	TypeName   string
	TypeNumber int
}

func RandomNumber(n int) int {
	rand.Seed(time.Now().Unix()) //give different number
	value := rand.Intn(n)
	return value
}
