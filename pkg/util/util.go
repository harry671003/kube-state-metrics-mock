package util

import (
	"crypto/sha1"
	"encoding/hex"
	"hash/fnv"
	"math"
	"math/rand"
)

func RandBetween(max, min float64) float64 {
	return math.Floor(float64(min) + rand.Float64()*float64(max-min))
}

func RandFloatBetween(max, min float64) float64 {
	return float64(min) + rand.Float64()*float64(max-min)
}

func FNV32a(text string) uint32 {
	algorithm := fnv.New32a()
	algorithm.Write([]byte(text))
	return algorithm.Sum32()
}

func GetFixedAssignment(item string, items []string) string {
	h := FNV32a(item)
	numNodes := uint32(len(items))
	return items[h%numNodes]
}

func SHA1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
