package tests

import (
	"math"
	"testing"
)

type testItem struct {
	Input    []float64
	Expected []float64
	Period   int
}

type test_func func([]float64, int) ([]float64, error)

func test_worker(t *testing.T, tests []*testItem, f test_func, round_floats float64) {
	for _, test := range tests {

		res, err := f(test.Input, test.Period)

		if err != nil {
			t.Errorf("Unexpected error %s", err.Error())
		} else {
			compare_inp_exp(t, res, test.Expected, round_floats)
		}
	}
}

func compare_inp_exp(t *testing.T, res, expected []float64, round_floats float64) {
	if len(res) != len(expected) {
		t.Errorf(`
			Different sizes: result %d, expected %d
			res: %v,
			expected: %v
		`, len(res), len(expected), res, expected,
		)
	}

	for k, v := range res {
		rounded_v := math.Round(v*(math.Pow(10.0, round_floats))) / math.Pow(10.0, round_floats)
		rounded_e := math.Round(expected[k]*(math.Pow(10.0, round_floats))) / math.Pow(10.0, round_floats)

		if ((rounded_v != rounded_e) && !math.IsNaN(v) && !math.IsNaN(expected[k])) ||
			(math.IsNaN(rounded_e) && !math.IsNaN(rounded_v)) ||
			(math.IsNaN(rounded_v) && !math.IsNaN(rounded_e)) {
			t.Errorf(`
							Expected result %#v
							Got result		%#v
							%0.8f instead of %0.8f
							(%0.8f instead of %0.8f)
						`, expected, res, rounded_v, rounded_e, v, expected[k])
			break
		}
	}
}
