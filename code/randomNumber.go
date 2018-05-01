package code

import (
	"math/rand"
	"time"
)

//random return random number
func random(min, max int) (number int) {
	rand.Seed(time.Now().UnixNano())
	number = rand.Intn(max-min) + min
	return
}
