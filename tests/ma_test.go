package tests

import (
	"math"
	"testing"

	"../../ta"
)

func TestEmptyData(t *testing.T) {
	var data []float64 = make([]float64, 0)
	var period int = 0

	res, err := ta.SMA(data, period)
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

	res, err := ta.SMA(data, period)
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

	res, err := ta.SMA(data, period)
	if err != nil {
		if err.Error() != "Invalid period" {
			t.Errorf("Unexpected error %s", err.Error())
		}
	} else {
		t.Errorf("Must be error, got %#v as res instead", res)
	}
}

func TestEmaPeriod(t *testing.T) {
	var data []float64 = []float64{
		1.0, 2.0, 3.0,
	}
	var period int = 1

	res, err := ta.EMA(data, period)
	if err != nil {
		if err.Error() != "Invalid period" {
			t.Errorf("Unexpected error %s", err.Error())
		}
	} else {
		t.Errorf("Must be error, got %#v as res instead", res)
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
	test_worker(t, tests, ta.SMA, 8)
}

func TestEMACalculation(t *testing.T) {

	tests := []*testItem{
		// numbers taken here https://stockcharts.com/school/doku.php?id=chart_school:technical_indicators:moving_averages
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
	test_worker(t, tests, ta.EMA, 2)
}

func test_workerCD(t *testing.T) {

	type MACDTestItem struct {
		Input        []float64
		FastPeriod   int
		SlowPeriod   int
		SignalPeriod int

		ExpectedMACD       []float64
		ExpectedMACDSignal []float64
		ExpectedHist       []float64
	}

	MACDTests := []*MACDTestItem{

		&MACDTestItem{
			Input: []float64{
				1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0,
				1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0,
				1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0,
				1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0,
				1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0,
				1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0,
				1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0,
				1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0,
				1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0,
				1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0,
			},
			FastPeriod:   12,
			SlowPeriod:   26,
			SignalPeriod: 9,

			ExpectedMACD: []float64{
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
				-0.06448, 0.101916, 0.31089389, 0.55085213, 0.81234816,
				0.29001772, -0.04274849, -0.22320352, -0.28226996, -0.24555825,
				-0.13422505, 0.03430353, 0.24572274, 0.4883364, 0.75262531,
				0.23316363, -0.09670707, -0.27427862, -0.33050387, -0.29101648,
				-0.17699057, -0.00586495, 0.20804673, 0.45304254, 0.71960002,
				0.20229207, -0.12553943, -0.30118475, -0.35559421, -0.31439827,
				-0.19876729, -0.02613597, 0.1891864, 0.43550238, 0.70329407,
				0.18713892, -0.13961671, -0.31425867, -0.36773305, -0.32566616,
				-0.20922439, -0.03583869, 0.18018531, 0.42715357, 0.69555146,
				0.17995947, -0.14627311, -0.32042942, -0.37345298, -0.3309677,
				-0.21413772, -0.04039187, 0.17596619, 0.42324426, 0.69192941,
				0.17660378, -0.14938188, -0.32330931, -0.37612072, -0.33343883,
				-0.21642665, -0.04251196, 0.17400254, 0.42142554, 0.69024499,
				0.17504376, -0.15082665, -0.32464733, -0.37735985, -0.33458636,
				-0.21748933, -0.04349606, 0.17309122, 0.42058164, 0.68946351,
			},
			ExpectedMACDSignal: []float64{
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
				math.NaN(), math.NaN(), math.NaN(), 0.16148066, 0.08007288,
				0.03721329, 0.03663134, 0.07844962, 0.16042697, 0.27886664,
				0.26972604, 0.19643942, 0.10229581, 0.01573588, -0.0456146,
				-0.07188979, -0.05868482, -0.00533851, 0.0863377, 0.21299016,
				0.21085054, 0.14357255, 0.05462109, -0.02742197, -0.08481723,
				-0.10760724, -0.09131299, -0.03521311, 0.05892999, 0.18780281,
				0.18767003, 0.12221268, 0.03491841, -0.04561188, -0.10162274,
				-0.12314307, -0.10568219, -0.04850869, 0.04662376, 0.1764093,
				0.17711933, 0.11244085, 0.02586679, -0.05399716, -0.10939127,
				-0.13034056, -0.11235082, -0.05468742, 0.04089892, 0.17110502,
				0.17220477, 0.10788744, 0.02164809, -0.05790567, -0.11301231,
				-0.13369517, -0.11545853, -0.05756632, 0.03823205, 0.16863464,
				0.16991646, 0.10576784, 0.01968481, -0.05972412, -0.11469657,
				-0.13525512, -0.11690331, -0.0589044, 0.0369928, 0.16748694,
			},
			ExpectedHist: []float64{
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
				math.NaN(), math.NaN(), math.NaN(), -0.44375062, -0.32563113,
				-0.17143834, -0.00232781, 0.16727312, 0.32790942, 0.47375867,
				-0.03656241, -0.29314649, -0.37657443, -0.34623974, -0.24540189,
				-0.10510078, 0.05281987, 0.21338525, 0.36670484, 0.50660986,
				-0.00855848, -0.26911198, -0.35580584, -0.32817224, -0.22958104,
				-0.09116004, 0.06517701, 0.22439951, 0.37657239, 0.51549127,
				-0.00053111, -0.26182939, -0.34917708, -0.32212117, -0.22404342,
				-0.08608132, 0.06984351, 0.22869401, 0.38052981, 0.51914216,
				0.00284013, -0.25871396, -0.34629622, -0.31945582, -0.22157643,
				-0.08379716, 0.07195896, 0.23065361, 0.38234534, 0.5208244,
				0.00439901, -0.25726932, -0.3449574, -0.31821505, -0.22042653,
				-0.08273148, 0.07294657, 0.23156885, 0.38319349, 0.52161035,
				0.00512729, -0.25659449, -0.34433213, -0.31763573, -0.21988979,
				-0.08223421, 0.07340725, 0.23199562, 0.38358883, 0.52197657,
			},
		},
	}

	for _, test := range MACDTests {
		macd, macdsignal, macdhist, err := ta.MACD(test.Input, test.FastPeriod,
			test.SlowPeriod, test.SignalPeriod)

		if err != nil {
			t.Errorf(err.Error())
		}
		compare_inp_exp(t, macd, test.ExpectedMACD, 8)
		compare_inp_exp(t, macdsignal, test.ExpectedMACDSignal, 8)
		compare_inp_exp(t, macdhist, test.ExpectedHist, 8)
	}
}
