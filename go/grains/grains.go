package grains

import (
	"math"
)

func Square(number int) uint64 {
	return uint64(math.Pow(2, float64(number-1)))
}

func Total() uint64 {
	return uint64(math.Pow(2, float64(63))) + uint64(math.Pow(2, float64(63))) - 1
}
