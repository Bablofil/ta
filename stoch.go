package ta

import (
	"math"
)

func STOCH(high, low, closes []float64, fastk_period, slowk_period, slowd_period int) (fastk, slowd []float64, err error) {

	fastk = make([]float64, len(closes))
	for i := 0; i < len(closes); i++ {
		if (i + 1) < fastk_period {
			fastk[i] = math.NaN()
		} else {
			lower_bound := i + 1 - fastk_period
			upper_bound := i + 1
			curr_low := SliceMin(low[lower_bound:upper_bound])
			curr_high := SliceMax(high[lower_bound:upper_bound])
			fastk[i] = ((closes[i] - curr_low) / (curr_high - curr_low)) * 100
		}
	}

	fastk, err = EMA(fastk, slowk_period)
	if err != nil {
		return
	}

	slowd, err = EMA(fastk, slowd_period)
	if err != nil {
		return
	}

	return
}

func STOCHRSI(data []float64, period, fastk_period, fastd_period int) (fastk, slowd []float64, err error) {
	rsi, err := RSI(data, period)
	if err != nil {
		return
	}

	return STOCH(rsi, rsi, rsi, period, fastk_period, fastd_period)
}
