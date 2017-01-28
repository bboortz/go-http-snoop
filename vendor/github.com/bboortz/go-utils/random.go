package utils

import (
	"math"
	"math/rand"
	"time"
)

func init() {
	Seed()
}

func Seed() {
	rand.Seed(GenSeed())
	rand.Seed(int64(RandomInt64(0, math.MaxInt64)))
}

func GenSeed() int64 {
	return time.Now().UTC().UnixNano() + int64(RandomInt(0, 9999999))
}

func RandomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(RandomInt(65, 90))
	}
	return string(bytes)
}

func RandomInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func RandomInt64(min int64, max int64) int64 {
	return min + rand.Int63n(max-min)
}
