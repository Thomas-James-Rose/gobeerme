package truerand

import (
	"math/rand"
	"time"
)

func Int(min int, max int) int {
	seed := rand.NewSource(time.Now().UnixNano())
	nonDeterministicRand := rand.New(seed)
	return nonDeterministicRand.Intn(max-min+1) + min
}
