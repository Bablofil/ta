// ta project ta.go
package ta

import (
	"errors"
	"math"
)

// Data expected to be older->newer, fair for result as well
func SMA(data []float64, period int) (result []float64, err error) {

	if len(data) == 0 {
		return result, errors.New("input parameter 'data' is empty")
	}

	if period <= 0 {
		return result, errors.New("Invalid period")

	}

	var interm float64
	for i := 0; i < len(data); i++ {
		interm += data[i]
		if (i + 1) < period {
			result = append(result, math.NaN())
		} else {
			result = append(result, interm/float64(period))
			interm -= data[i+1-period]
		}
	}
	return result, nil
}

func EMA(data []float64, period int) (result []float64, err error) {

	sma, err := SMA(data, period)
	if err != nil {
		return
	}

	var multiplier float64 = 2 / (float64(period) + 1)

	for k, v := range sma {
		if math.IsNaN(v) {
			result = append(result, math.NaN())
		} else {
			prev := result[k-1]
			if math.IsNaN(prev) {
				result = append(result, v)
				continue
			}
			ema := (data[k]-prev)*multiplier + prev
			result = append(result, ema)
		}
	}
	return result, nil
}
