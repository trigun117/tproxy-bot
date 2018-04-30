package code

import (
	"math/rand"
	"time"
)

//Random return random number
func Random(min, max int) (number int) {
	rand.Seed(time.Now().UnixNano())
	number = rand.Intn(max-min) + min
	return
}
