package ta

import (
	"math"
)

func SliceMax(slice []float64) (local_max float64) {
	//local_max == neg. inf at start
	local_max = math.Inf(-1)
	for _, v := range slice {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return math.NaN()
		}
		if v > local_max {
			local_max = v
		}
	}
	//if ended with either infinity
	if math.IsInf(local_max, 0) {
		return math.NaN()
	}
	return
}

func SliceMin(slice []float64) (local_min float64) {
	//local_min == pos. inf at start
	local_min = math.Inf(1)
	for _, v := range slice {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return math.NaN()
		}
		if v < local_min {
			local_min = v
		}
	}
	//if ended with either infinity
	if math.IsInf(local_min, 0) {
		return math.NaN()
	}
	return
}

func SliceSum(slice []float64) (sum float64) {
	for _, v := range slice {
		sum += v
	}
	return
}
