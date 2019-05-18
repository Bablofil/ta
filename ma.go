package ta

import (
	"errors"
	"math"
)

// Data expected to be older->newer, fair for result as well
func SMA(data []float64, period int) (result []float64, err error) {
	if len(data) == 0 {
		return result, errors.New("SMA: input parameter 'data' is empty")
	}

	if period <= 0 {
		return result, errors.New("SMA: Invalid period")
	}

	var interm float64

	for i := 0; i < len(data); i++ {
		if math.IsNaN(data[i]) {
			result = append(result, math.NaN())
			interm = 0
		} else {
			interm += data[i]
			if (i + 1) < period {
				result = append(result, math.NaN())
			} else {
				result = append(result, interm/float64(period))
				if !math.IsNaN(data[i+1-period]) {
					interm -= data[i+1-period]
				}
			}
		}
	}
	return result, nil
}

func generalEMA(data []float64, period int, multiplier float64) (result []float64, err error) {
	if period <= 1 {
		return result, errors.New("Invalid period")
	}

	sma, err := SMA(data, period)
	if err != nil {
		return
	}

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

func EMA(data []float64, period int) (result []float64, err error) {
	return generalEMA(data, period, 2/(float64(period)+1))
}

/* SMMA, MMA and RMA are synonyms (https://en.wikipedia.org/wiki/Moving_average) */
func SMMA(data []float64, period int) (result []float64, err error) {
	return generalEMA(data, period, 1/float64(period))
}
func MMA(data []float64, period int) (result []float64, err error) {
	return generalEMA(data, period, 1/float64(period))
}
func RMA(data []float64, period int) (result []float64, err error) {
	return generalEMA(data, period, 1/float64(period))
}

func MACD(data []float64, fastperiod, slowperiod, signalperiod int) (macd, macdsignal, macdhist []float64, err error) {
	fast_ema, err := EMA(data, fastperiod)
	if err != nil {
		return
	}

	slow_ema, err := EMA(data, slowperiod)
	if err != nil {
		return
	}

	macd = make([]float64, len(fast_ema))
	//macdsignal = make([]float64, 0)
	macdsignal = make([]float64, len(fast_ema))
	//diff := make([]float64, 0)

	for k, fast := range fast_ema {
		if math.IsNaN(fast) || math.IsNaN(slow_ema[k]) {
			macd[k] = math.NaN()
			//macdsignal = append(macdsignal, math.NaN())
			macdsignal[k] = math.NaN()
		} else {
			macd[k] = fast - slow_ema[k]
			//diff = append(diff, macd[k])
			macdsignal[k] = macd[k]
		}
	}

	/*diff_ema, err := EMA(diff, signalperiod)

	if err != nil {
		return
	}
	macdsignal = append(macdsignal, diff_ema...)
	*/
	macdsignal, err = EMA(macdsignal, signalperiod)
	if err != nil {
		return
	}
	macdhist = make([]float64, len(macd))

	for k, ms := range macdsignal {
		if math.IsNaN(ms) || math.IsNaN(macd[k]) {
			macdhist[k] = math.NaN()
		} else {
			macdhist[k] = macd[k] - macdsignal[k]
		}
	}

	return
}
