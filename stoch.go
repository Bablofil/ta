package ta

import (
	"math"
)

func STOCH(high, low, closes []float64, fastk_period, slowk_period, slowd_period int) (fastk, slowd []float64, err error) {

	fastk_dummy := make([]float64, fastk_period-1)
	fastk_vals := make([]float64, len(closes)-fastk_period+1)

	for i := 0; i < len(closes); i++ {

		if (i + 1) < fastk_period {
			fastk_dummy[i] = math.NaN()
		} else {
			lower_bound := i + 1 - fastk_period
			upper_bound := i + 1
			curr_low := SliceMin(low[lower_bound:upper_bound])
			curr_high := SliceMax(high[lower_bound:upper_bound])

			fastk_vals[i+1-fastk_period] = ((closes[i] - curr_low) / (curr_high - curr_low)) * 100
		}
	}

	fastk_ema, err := EMA(fastk_vals, slowk_period)

	if err != nil {
		return
	}

	fastk = append(fastk_dummy, fastk_ema...)
	slowd_dummy := make([]float64, 0)
	var offset int
	for _, v := range fastk {
		if math.IsNaN(v) {
			slowd_dummy = append(slowd_dummy, math.NaN())
			offset++
		} else {
			break
		}
	}

	slowd_ema, err := EMA(fastk[offset:], slowd_period)
	if err != nil {
		return
	}
	slowd = append(slowd_dummy, slowd_ema...)

	return
}
