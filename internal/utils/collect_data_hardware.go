package utils

import (
	"fmt"
	"strconv"
)

func BytesToGB(b uint64) float64 {
	return TruncateFloat(float64(b) / (1024 * 1024 * 1024))
}

func TruncateFloat(value float64) float64 {
	formatted := fmt.Sprintf("%.2f", value)
	valueFinal, _ := strconv.ParseFloat(formatted, 64)
	return valueFinal
}
