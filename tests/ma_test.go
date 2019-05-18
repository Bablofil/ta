package tests

import (
	"math"
	"strings"
	"testing"

	"../../ta"
)

func TestEmptyData(t *testing.T) {
	var data []float64 = make([]float64, 0)
	var period int = 0

	res, err := ta.SMA(data, period)
	if err != nil {
		if !strings.Contains(err.Error(), "'data' is empty") {
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
		if !strings.Contains(err.Error(), "Invalid period") {
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
		if !strings.Contains(err.Error(), "Invalid period") {
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
		if !strings.Contains(err.Error(), "Invalid period") {
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

func TestMACD(t *testing.T) {

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
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), -0.00716444, 0.01465165, 0.07390009, 0.16929050, 0.29790203,
				0.29632517, 0.22851044, 0.13816764, 0.05408012, -0.00584755, -0.03152305, -0.01835773, 0.03445836, 0.12523397, 0.25071224,
				0.24720252, 0.17842060, 0.08788076, 0.00420383, -0.05484023, -0.07927030, -0.06458923, -0.01006204, 0.08255888, 0.20996711,
				0.20843210, 0.14163779, 0.05307328, -0.02866021, -0.08580783, -0.10839972, -0.09194697, -0.03572029, 0.05852424, 0.18747821,
				0.18741035, 0.12200494, 0.03475222, -0.04574484, -0.10172910, -0.12322816, -0.10575027, -0.04856315, 0.04658020, 0.17637445,
				0.17709145, 0.11241854, 0.02584895, -0.05401144, -0.10940269, -0.13034970, -0.11235813, -0.05469327, 0.04089424, 0.17110127,
				0.17220177, 0.10788504, 0.02164617, -0.05790721, -0.11301353, -0.13369616, -0.11545932, -0.05756695, 0.03823155, 0.16863424,
				0.16991614, 0.10576758, 0.01968460, -0.05972429, -0.11469670, -0.13525523, -0.11690339, -0.05890447, 0.03699275, 0.16748690,
			},
			ExpectedHist: []float64{
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), -0.05731555, 0.08726436, 0.23699379, 0.38156163, 0.51444613,
				-0.00630745, -0.27125893, -0.36137117, -0.33635008, -0.23971070, -0.10270200, 0.05266126, 0.21126438, 0.36310243, 0.50191308,
				-0.01403888, -0.27512767, -0.36215937, -0.33470770, -0.23617625, -0.09772027, 0.05872427, 0.21810877, 0.37048366, 0.50963291,
				-0.00614003, -0.26717723, -0.35425803, -0.32693400, -0.22859045, -0.09036757, 0.06581099, 0.22490670, 0.37697814, 0.51581587,
				-0.00027143, -0.26162165, -0.34901089, -0.32198821, -0.22393706, -0.08599623, 0.06991158, 0.22874846, 0.38057338, 0.51917701,
				0.00286802, -0.25869165, -0.34627837, -0.31944154, -0.22156501, -0.08378802, 0.07196627, 0.23065946, 0.38235002, 0.52082814,
				0.00440200, -0.25726692, -0.34495548, -0.31821352, -0.22042530, -0.08273049, 0.07294736, 0.23156948, 0.38319399, 0.52161075,
				0.00512761, -0.25659424, -0.34433193, -0.31763556, -0.21988966, -0.08223411, 0.07340733, 0.23199569, 0.38358889, 0.52197661,
			},
		},
	}

	for _, test := range MACDTests {
		macd, macdsignal, macdhist, err := ta.MACD(test.Input, test.FastPeriod,
			test.SlowPeriod, test.SignalPeriod)

		if err != nil {
			t.Errorf(err.Error())
		}
		compare_inp_exp(t, macd, test.ExpectedMACD, 8, "MACD")
		compare_inp_exp(t, macdsignal, test.ExpectedMACDSignal, 8, "MACDSignal")
		compare_inp_exp(t, macdhist, test.ExpectedHist, 8, "MACDHIST")
	}
}
