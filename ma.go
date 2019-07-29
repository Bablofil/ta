package ta

import (
	"errors"
	"math"
)

type ma_func func([]float64, int) ([]float64, error)

// Simple moving average
// https://en.wikipedia.org/wiki/Moving_average
// Data expected to be older->newer, fair for result as well
func SMA(data []float64, period int) (result []float64, err error) {
	if len(data) == 0 {
		return result, errors.New("SMA: input parameter 'data' is empty")
	}

	if period <= 0 {
		return result, errors.New("SMA: Invalid period")
	}

	var interm float64
	var nan_inp int = 0

	for i := 0; i < len(data); i++ {
		if math.IsNaN(data[i]) {
			result = append(result, math.NaN())
			interm = 0
			nan_inp++
		} else {
			interm += data[i]
			if (i + 1 - nan_inp) < period {
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

// Calculates various EMA with different smoothing multipliers, see lower
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

// Exponential moving average
// https://en.wikipedia.org/wiki/Moving_average#Exponential_moving_average
func EMA(data []float64, period int) (result []float64, err error) {
	return generalEMA(data, period, 2/(float64(period)+1))
}

// Synonym to EMA
func EWMA(data []float64, period int) (result []float64, err error) {
	return EMA(data, period)
}

// Modified moving average
// https://en.wikipedia.org/wiki/Moving_average
// SMMA, MMA and RMA are synonyms according to wiki
func SMMA(data []float64, period int) (result []float64, err error) {
	return generalEMA(data, period, 1/float64(period))
}
func MMA(data []float64, period int) (result []float64, err error) {
	return generalEMA(data, period, 1/float64(period))
}
func RMA(data []float64, period int) (result []float64, err error) {
	return generalEMA(data, period, 1/float64(period))
}

// Double exponential moving average
// https://en.wikipedia.org/wiki/Double_exponential_moving_average
func D2(data []float64, period int) (result []float64, err error) {
	ema, err := EMA(data, period)
	if err != nil {
		return
	}
	ema_ema, err := EMA(ema, period)
	if err != nil {
		return
	}
	for k, v := range ema {
		result = append(result, v*2-ema_ema[k])
	}
	return
}

// Synonym for double exponential moving average
func DEMA(data []float64, period int) (result []float64, err error) {
	return D2(data, period)
}

// Triple Exponential Moving Average
// https://en.wikipedia.org/wiki/Triple_exponential_moving_average
func T3(data []float64, period int) (result []float64, err error) {
	e1, err := EMA(data, period)
	if err != nil {
		return
	}
	e2, err := EMA(e1, period)
	if err != nil {
		return
	}
	e3, err := EMA(e2, period)
	if err != nil {
		return
	}
	for k, _ := range e1 {
		result = append(result, e1[k]*3-e2[k]*3+e3[k])
	}
	return
}

// Synonym for Triple Exponential Moving Average
func TEMA(data []float64, period int) (result []float64, err error) {
	return T3(data, period)
}

// Synonym for Triple Exponential Moving Average
func TMA(data []float64, period int) (result []float64, err error) {
	return T3(data, period)
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
	macdsignal = make([]float64, len(fast_ema))

	for k, fast := range fast_ema {
		if math.IsNaN(fast) || math.IsNaN(slow_ema[k]) {
			macd[k] = math.NaN()
			macdsignal[k] = math.NaN()
		} else {
			macd[k] = fast - slow_ema[k]
			macdsignal[k] = macd[k]
		}
	}

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
