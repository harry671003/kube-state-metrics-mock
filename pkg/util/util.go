package util

import (
	"math"
	"math/rand"
)

func RandBetween(max, min float64) float64 {
	return math.Floor(float64(min) + rand.Float64()*float64(max-min))
}
