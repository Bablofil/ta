package ta

import (
	"math"
	"testing"
)

type testItem struct {
	Input    []float64
	Expected []float64
	Period   int
}

type ma_func func([]float64, int) ([]float64, error)

func TestEmptyData(t *testing.T) {
	var data []float64 = make([]float64, 0)
	var period int = 0

	res, err := SMA(data, period)
	if err != nil {
		if err.Error() != "input parameter 'data' is empty" {
			t.Errorf("Unexpected error %s", err.Error())
		}
	} else {
		t.Errorf("Must be error, got %#v as res instead", res)
	}

}

func TestZeroPeriod(t *testing.T) {
	var data []float64 = []float64{
		1.0, 2.0, 3.0,
	}
	var period int = 0

	res, err := SMA(data, period)
	if err != nil {
		if err.Error() != "Invalid period" {
			t.Errorf("Unexpected error %s", err.Error())
		}
	} else {
		t.Errorf("Must be error, got %#v as res instead", res)
	}
}

func TestNegativePeriod(t *testing.T) {
	var data []float64 = []float64{
		1.0, 2.0, 3.0,
	}
	var period int = -1

	res, err := SMA(data, period)
	if err != nil {
		if err.Error() != "Invalid period" {
			t.Errorf("Unexpected error %s", err.Error())
		}
	} else {
		t.Errorf("Must be error, got %#v as res instead", res)
	}
}

func test_ma(t *testing.T, tests []*testItem, f ma_func, round_floats float64) {
	for _, test := range tests {

		res, err := f(test.Input, test.Period)

		if err != nil {
			t.Errorf("Unexpected error %s", err.Error())
		} else {
			for k, v := range res {
				rounded_v := math.Round(v*(math.Pow(10.0, round_floats))) / math.Pow(10.0, round_floats)
				rounded_e := math.Round(test.Expected[k]*(math.Pow(10.0, round_floats))) / math.Pow(10.0, round_floats)

				if ((rounded_v != rounded_e) && !math.IsNaN(v) && !math.IsNaN(test.Expected[k])) ||
					(math.IsNaN(rounded_e) && !math.IsNaN(rounded_v)) ||
					(math.IsNaN(rounded_v) && !math.IsNaN(rounded_e)) {
					t.Errorf(`
							Expected result %#v
							Got result		%#v
							%0.8f instead of %0.8f
							(%0.8f instead of %0.8f)
						`, test.Expected, res, rounded_v, rounded_e, v, test.Expected[k])
					break
				}
			}
		}
	}
}
func TestSMACalculation(t *testing.T) {

	tests := []*testItem{
		&testItem{
			Input: []float64{
				1.0, 2.0, 3.0, 4.0, 5.0,
			},
			Expected: []float64{
				1.0, 2.0, 3.0, 4.0, 5.0,
			},
			Period: 1,
		},

		&testItem{
			Input: []float64{
				1.0, 2.0, 3.0, 4.0, 5.0,
			},
			Expected: []float64{
				math.NaN(), 1.5, 2.5, 3.5, 4.5,
			},
			Period: 2,
		},

		&testItem{
			Input: []float64{
				1.0, 2.0, 3.0, 4.0, 5.0,
			},
			Expected: []float64{
				math.NaN(), math.NaN(), 2, 3, 4,
			},
			Period: 3,
		},
		&testItem{
			Input: []float64{
				1.0, 2.0, 3.0, 4.0, 5.0,
			},
			Expected: []float64{
				math.NaN(), math.NaN(), math.NaN(), 2.5, 3.5,
			},
			Period: 4,
		},
		&testItem{
			Input: []float64{
				1.0, 2.0, 3.0, 4.0, 5.0,
			},
			Expected: []float64{
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), 3,
			},
			Period: 5,
		},
		&testItem{
			Input: []float64{
				1.0, 2.0, 3.0, 4.0, 5.0,
			},
			Expected: []float64{
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
			},
			Period: 6,
		},
	}
	test_ma(t, tests, SMA, 8)
}

func TestEMACalculation(t *testing.T) {

	tests := []*testItem{
		// praise this guy https://stockcharts.com/school/doku.php?id=chart_school:technical_indicators:moving_averages
		&testItem{
			Input: []float64{
				22.2734, 22.1940, 22.0847, 22.1741,
				22.1840, 22.1344, 22.2337, 22.4323,
				22.2436, 22.2933, 22.1542, 22.3926,
				22.3816, 22.6109, 23.3558, 24.0519,
				23.7530, 23.8324, 23.9516, 23.6338,
				23.8225, 23.8722, 23.6537, 23.1870,
				23.0976, 23.3260, 22.6805, 23.0976,
				22.4025, 22.1725,
			},
			Expected: []float64{
				math.NaN(), math.NaN(), math.NaN(), math.NaN(),
				math.NaN(), math.NaN(), math.NaN(), math.NaN(),
				math.NaN(), 22.22, 22.21, 22.24,
				22.27, 22.33, 22.52, 22.80,
				22.97, 23.13, 23.28, 23.34,
				23.43, 23.51, 23.54, 23.47,
				23.40, 23.39, 23.26, 23.23,
				23.08, 22.92,
			},
			Period: 10,
		},
		&testItem{
			Input: []float64{
				1.0, 2.0, 3.0, 4.0, 5.0,
			},
			Expected: []float64{
				math.NaN(), 1.5, 2.5, 3.5, 4.5,
			},
			Period: 2,
		},
	}
	test_ma(t, tests, EMA, 2)
}
