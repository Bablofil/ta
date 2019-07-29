package ta

import (
	"math"
)

func BBANDS(data []float64, ma ma_func, ma_period, dev_val int) (upper, middle, lower []float64, err error) {
	middle, err = ma(data, ma_period)
	if err != nil {
		return
	}

	var stddevs []float64
	var real_data_cnt int = 0

	for i, _ := range data {
		if math.IsNaN(middle[i]) {
			stddevs = append(stddevs, 0)
			real_data_cnt++
			continue
		}

		if i-real_data_cnt >= ma_period {

			sum := SliceSum(middle[i-ma_period+1 : i+1])
			avg := sum / float64(ma_period)

			var s float64
			for _, v := range middle[i-ma_period+1 : i+1] {
				s += math.Pow(v-avg, 2)
			}

			stddev_avg := s / float64(ma_period)
			stddev := math.Sqrt(stddev_avg)

			stddevs = append(stddevs, stddev)
		} else {
			stddevs = append(stddevs, 0)
		}

	}

	for i, _ := range middle {

		if !math.IsNaN(middle[i]) {
			upper = append(upper, middle[i]+stddevs[i]*float64(dev_val))
			lower = append(lower, middle[i]-stddevs[i]*float64(dev_val))
		} else {
			upper = append(upper, math.NaN())
			lower = append(lower, math.NaN())
		}
	}
	return
}
