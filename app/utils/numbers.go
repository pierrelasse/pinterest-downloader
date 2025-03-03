package utils

func Numbers_roundTo7Decimals(value float64) float64 {
	return float64(int(value*1e7+0.5)) / 1e7
}
