package ta

func RSI(data []float64, period int) (result []float64, err error) {
	u_days := make([]float64, len(data))
	d_days := make([]float64, len(data))

	for i, _ := range data {
		/* can't compare first day and day -1, assert they're equal */
		if i == 0 {
			u_days[0] = 0
			d_days[0] = 0
		} else {
			if data[i] > data[i-1] {
				u_days[i] = data[i] - data[i-1]
			} else if data[i] < data[i-1] {
				d_days[i] = data[i-1] - data[i]
			}
		}
	}

	smma_u, err := SMMA(u_days, period)
	if err != nil {
		return
	}
	smma_d, err := SMMA(d_days, period)
	if err != nil {
		return
	}

	result = make([]float64, len(data))

	for k, _ := range data {
		if smma_d[k] == 0 {
			result[k] = 100
		} else {
			result[k] = 100 - (100 / (1 + smma_u[k]/smma_d[k]))
		}
	}
	return
}


