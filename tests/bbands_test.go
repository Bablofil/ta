package tests

import (
	"math"
	"testing"

	"../../ta"
)

func TestBBANDSCalculation(t *testing.T) {
	var data []float64 = []float64{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12,
		13, 14, 15, 14, 13, 12, 11, 10, 9,
		8, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7,
		8, 9, 10, 11, 12, 13, 14, 15, 14,
		13, 12, 11, 10, 9, 8, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
		12, 13, 14, 15, 14, 13, 12, 11,
		10, 9, 8, 7, 8, 9, 10, 1, 2, 3, 4,
		5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
		15, 14, 13, 12, 11, 10, 9, 8, 7, 8, 9, 10,
	}

	var expected_upper []float64 = []float64{
		math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
		math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), 9.00000000,
		9.40000000, 9.70000000, 9.90000000, 10.10000000, 10.30000000, 10.50000000, 10.20000000, 9.90000000, 9.60000000, 9.30000000,
		9.00000000, 8.70000000, 8.40000000, 8.10000000, 7.80000000, 7.60000000, 7.50000000, 7.50000000, 7.60000000, 9.88851622,
		10.20864886, 10.48353066, 10.72333882, 10.82922264, 10.79120071, 10.58552830, 10.85227408, 11.11629799, 11.36778745, 11.68921364,
		12.05371035, 12.43522608, 12.23319945, 11.95958734, 11.61720103, 11.21739406, 10.77817322, 10.32662227, 9.90129944, 9.54830936,
		9.30129944, 9.22662227, 9.27817322, 9.41611586, 9.62136093, 9.88851622, 10.20864886, 10.48353066, 10.72333882, 10.82922264,
		10.79120071, 10.58552830, 10.85227408, 11.11629799, 11.36778745, 11.68921364, 12.05371035, 12.43522608, 12.23319945, 11.95958734,
		11.61720103, 11.21739406, 10.77817322, 10.32662227, 9.90129944, 9.54830936, 9.30129944, 9.22662227, 9.27817322, 9.41611586,
		9.62136093, 9.88851622, 10.20864886, 10.48353066, 10.72333882, 10.82922264, 10.79120071, 10.58552830, 10.85227408, 11.11629799,
		11.36778745, 11.68921364, 12.05371035, 12.43522608,
	}

	var expected_middle []float64 = []float64{
		math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
		math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), 9.00000000,
		9.40000000, 9.70000000, 9.90000000, 10.10000000, 10.30000000, 10.50000000, 10.20000000, 9.90000000, 9.60000000, 9.30000000,
		9.00000000, 8.70000000, 8.40000000, 8.10000000, 7.80000000, 7.60000000, 7.50000000, 7.50000000, 7.60000000, 7.80000000,
		8.10000000, 8.40000000, 8.70000000, 8.90000000, 9.00000000, 9.00000000, 9.40000000, 9.70000000, 9.90000000, 10.10000000,
		10.30000000, 10.50000000, 10.20000000, 9.90000000, 9.60000000, 9.30000000, 9.00000000, 8.70000000, 8.40000000, 8.10000000,
		7.80000000, 7.60000000, 7.50000000, 7.50000000, 7.60000000, 7.80000000, 8.10000000, 8.40000000, 8.70000000, 8.90000000,
		9.00000000, 9.00000000, 9.40000000, 9.70000000, 9.90000000, 10.10000000, 10.30000000, 10.50000000, 10.20000000, 9.90000000,
		9.60000000, 9.30000000, 9.00000000, 8.70000000, 8.40000000, 8.10000000, 7.80000000, 7.60000000, 7.50000000, 7.50000000,
		7.60000000, 7.80000000, 8.10000000, 8.40000000, 8.70000000, 8.90000000, 9.00000000, 9.00000000, 9.40000000, 9.70000000,
		9.90000000, 10.10000000, 10.30000000, 10.50000000,
	}

	var expected_lower []float64 = []float64{
		math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
		math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), 9.00000000,
		9.40000000, 9.70000000, 9.90000000, 10.10000000, 10.30000000, 10.50000000, 10.20000000, 9.90000000, 9.60000000, 9.30000000,
		9.00000000, 8.70000000, 8.40000000, 8.10000000, 7.80000000, 7.60000000, 7.50000000, 7.50000000, 7.60000000, 5.71148378,
		5.99135114, 6.31646934, 6.67666118, 6.97077736, 7.20879929, 7.41447170, 7.94772592, 8.28370201, 8.43221255, 8.51078636,
		8.54628965, 8.56477392, 8.16680055, 7.84041266, 7.58279897, 7.38260594, 7.22182678, 7.07337773, 6.89870056, 6.65169064,
		6.29870056, 5.97337773, 5.72182678, 5.58388414, 5.57863907, 5.71148378, 5.99135114, 6.31646934, 6.67666118, 6.97077736,
		7.20879929, 7.41447170, 7.94772592, 8.28370201, 8.43221255, 8.51078636, 8.54628965, 8.56477392, 8.16680055, 7.84041266,
		7.58279897, 7.38260594, 7.22182678, 7.07337773, 6.89870056, 6.65169064, 6.29870056, 5.97337773, 5.72182678, 5.58388414,
		5.57863907, 5.71148378, 5.99135114, 6.31646934, 6.67666118, 6.97077736, 7.20879929, 7.41447170, 7.94772592, 8.28370201,
		8.43221255, 8.51078636, 8.54628965, 8.56477392,
	}

	upper, middle, lower, err := ta.BBANDS(data, ta.SMA, 20, 2)
	if err != nil {
		t.Errorf(err.Error())
	}
	compare_inp_exp(t, upper, expected_upper, 8, "UPPER")
	compare_inp_exp(t, middle, expected_middle, 8, "MIDDLE")
	compare_inp_exp(t, lower, expected_lower, 8, "LOWER")
}