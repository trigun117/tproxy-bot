package main

import (
	"math/rand"
	"time"
)

func random(min, max int) (number int) {
	rand.Seed(time.Now().UnixNano())
	number = rand.Intn(max-min) + min
	return
}
