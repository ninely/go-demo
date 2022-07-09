package study

import (
	"math/rand"
	"time"
)

func MathRandRepeat() int64 {
	target, err := time.Parse("2006-06-01 00:00:00", "2022-07-08 00:00:00")
	if err != nil {
		panic(err)
	}
	rand.Seed(target.UnixNano())
	return rand.Int63()
}
