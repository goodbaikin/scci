package compute

import (
	"math"
)

func length(values []float64) float64 {
	return float64(len(values))
}

func sum(values []float64) float64 {
	result := 0.0
	for _, v := range values {
		result += v
	}

	return result
}

func Average(values []float64) float64 {
	return sum(values) / length(values)
}

func variance(values []float64) float64 {
	v0 := Average(values)
	result := 0.0

	for _, v := range values {
		result += (v - v0) * (v - v0)
	}
	result /= length(values)

	return result
}

func std(values []float64) float64 {
	return math.Sqrt(variance(values))
}

func Error(values []float64, r float64) float64 {
	return r * std(values)
}
