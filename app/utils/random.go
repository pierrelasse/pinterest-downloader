package utils

import (
	"math/rand"
	"time"
)

func Random_new(seed int64) *rand.Rand {
	return rand.New(rand.NewSource(seed))
}

func Random_newFromTime() *rand.Rand {
	return Random_new(time.Now().UnixNano())
}

func Random_randInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func Random_randFloat(min, max float32, decimals int) float32 {
	factor := float32(1)
	for i := 0; i < decimals; i++ {
		factor *= 10
	}
	value := rand.Float32()*(max-min) + min
	return float32(int(value*factor)) / factor
}

func Random_randBool() bool {
	return rand.Intn(2) == 1
}

func Random_float64(args ...float64) float64 {
	var min, max float64
	switch len(args) {
	case 0:
		min, max = 0, 1
	case 2:
		min, max = args[0], args[1]
	default:
		panic("Random_float64 accepts either no arguments or exactly 2 arguments (min, max)")
	}
	return rand.Float64()*(max-min) + min
}

func Random_getRandomData(choices []string) string {
	return choices[rand.Intn(len(choices))]
}
