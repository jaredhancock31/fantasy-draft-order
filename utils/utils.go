package utils

import (
	"math/rand/v2"
)


func RandRangeSlice(min, max int) []int {
	randSlice := make([]int, max-min)
	for i := min; i < max; i++ {
		randSlice[i] = i
	}
	// shuffle shuffle shuffle
	rand.Shuffle(len(randSlice), func(i, j int) {
		randSlice[i], randSlice[j] = randSlice[j], randSlice[i]
	})
	return randSlice
}
