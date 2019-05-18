package tests

import (
	"math"
	"testing"

	"../../ta"
)

type TestRollingItem struct {
	Input    []float64
	Expected float64
}

func TestRollingMax(t *testing.T) {
	test_cases := []*TestRollingItem{
		&TestRollingItem{Input: []float64{1.0, 2.0, 3.0}, Expected: 3},
		&TestRollingItem{Input: []float64{3.0, 2.0, 3.0}, Expected: 3},
		&TestRollingItem{Input: []float64{3.0, 3.0, 3.0}, Expected: 3},
		&TestRollingItem{Input: []float64{}, Expected: math.NaN()},
		&TestRollingItem{Input: []float64{-1, 0, -1}, Expected: 0},
		&TestRollingItem{Input: []float64{-1, math.NaN(), -1}, Expected: math.NaN()},
		&TestRollingItem{Input: []float64{-1, math.Inf(-1), -1}, Expected: math.NaN()},
		&TestRollingItem{Input: []float64{-1, -2, -5}, Expected: -1},
	}
	for _, t_case := range test_cases {
		res := ta.SliceMax(t_case.Input)
		if (res != t_case.Expected && !math.IsNaN(res) && !math.IsNaN(t_case.Expected)) || (math.IsNaN(res) && !math.IsNaN(t_case.Expected)) || (math.IsNaN(t_case.Expected) && !math.IsNaN(res)) {
			t.Errorf("expected %f , got %f", t_case.Expected, res)
		}
	}
}

func TestRollingMin(t *testing.T) {
	test_cases := []*TestRollingItem{
		&TestRollingItem{Input: []float64{1.0, 2.0, 3.0}, Expected: 1},
		&TestRollingItem{Input: []float64{3.0, 2.0, 3.0}, Expected: 2},
		&TestRollingItem{Input: []float64{3.0, 3.0, 3.0}, Expected: 3},
		&TestRollingItem{Input: []float64{}, Expected: math.NaN()},
		&TestRollingItem{Input: []float64{-1, 0, -1}, Expected: -1},
		&TestRollingItem{Input: []float64{-1, math.NaN(), -1}, Expected: math.NaN()},
		&TestRollingItem{Input: []float64{-1, math.Inf(-1), -1}, Expected: math.NaN()},
		&TestRollingItem{Input: []float64{-1, -2, -5}, Expected: -5},
	}
	for _, t_case := range test_cases {
		res := ta.SliceMin(t_case.Input)
		if (res != t_case.Expected && !math.IsNaN(res) && !math.IsNaN(t_case.Expected)) || (math.IsNaN(res) && !math.IsNaN(t_case.Expected)) || (math.IsNaN(t_case.Expected) && !math.IsNaN(res)) {
			t.Errorf("expected %f , got %f", t_case.Expected, res)
		}
	}
}
