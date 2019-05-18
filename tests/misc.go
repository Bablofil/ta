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
	for test_num, test := range tests {
		comment := "TEST_" + string(test_num)
		res, err := f(test.Input, test.Period)

		if err != nil {
			t.Errorf("%s: Unexpected error %s", comment, err.Error())
		} else {
			compare_inp_exp(t, res, test.Expected, round_floats, comment)
		}
	}
}

func compare_inp_exp(t *testing.T, res, expected []float64, round_floats float64, comment string) {
	if len(res) != len(expected) {
		t.Errorf(`
			TEST COMMENT %s: 
			Different sizes: result %d, expected %d
			res: %v,
			expected: %v
		`, comment, len(res), len(expected), res, expected,
		)
	}

	for k, v := range res {
		rounded_v := math.Round(v*(math.Pow(10.0, round_floats))) / math.Pow(10.0, round_floats)
		rounded_e := math.Round(expected[k]*(math.Pow(10.0, round_floats))) / math.Pow(10.0, round_floats)

		if ((rounded_v != rounded_e) && !math.IsNaN(v) && !math.IsNaN(expected[k])) ||
			(math.IsNaN(rounded_e) && !math.IsNaN(rounded_v)) ||
			(math.IsNaN(rounded_v) && !math.IsNaN(rounded_e)) {
			t.Errorf(`
							TEST COMMENT %s
							Expected result %#v
							Got result		%#v
							%0.8f instead of %0.8f
							(%0.8f instead of %0.8f)
						`, comment, expected, res, rounded_v, rounded_e, v, expected[k])
			break
		}
	}
}
