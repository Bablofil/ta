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

func TestD2Calculation(t *testing.T) {

	tests := []*testItem{
		&testItem{
			Input: []float64{
				1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 14, 13, 12, 11, 10, 9, 8, 7, 8, 9, 10,
			},
			Expected: []float64{
				math.NaN(), math.NaN(), 3.00000000, 4.00000000, 5.00000000, 6.00000000, 7.00000000, 8.00000000, 9.00000000, 10.00000000,
				11.00000000, 12.00000000, 13.00000000, 14.00000000, 15.00000000, 14.22222222, 13.14814815, 12.07407407, 11.03292181, 10.01371742,
				9.00548697, 8.00213382, 7.00081288, 7.77808261, 8.85196475, 9.92596732,
			},
			Period: 2,
		},
	}
	test_worker(t, tests, ta.D2, 8)
}

func TestT3Calculation(t *testing.T) {

	tests := []*testItem{
		&testItem{
			Input: []float64{
				1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 14, 13, 12, 11, 10, 9, 8, 7, 8, 9, 10,
			},
			Expected: []float64{
				math.NaN(), math.NaN(), math.NaN(), 4.00000000, 5.00000000, 6.00000000, 7.00000000, 8.00000000, 9.00000000, 10.00000000,
				11.00000000, 12.00000000, 13.00000000, 14.00000000, 15.00000000, 14.07407407, 13.00000000, 11.97530864, 10.97805213, 9.98628258,
				8.99268404, 7.99644363, 6.99837423, 7.92521465, 8.99969893, 10.02456717,
			},
			Period: 2,
		},
	}
	test_worker(t, tests, ta.T3, 8)
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
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
				math.NaN(), math.NaN(), math.NaN(), 0.16148065747445678, 0.08007287621765422, 0.03721329120823753, 0.036631338408597405,
				0.07844961907601512, 0.16042697468135714, 0.2788666424333089, 0.2697260408737835, 0.19643941933554376, 0.10229581177313908,
				0.015735876018179967, -0.045614595592292834, -0.07188978952215715, -0.05868482251281618, -0.00533851113976145, 0.08633769892179652,
				0.21299016336284352, 0.21085054372150316, 0.14357254820636867, 0.054621089405418294, -0.0274219703997097, -0.08481723061562085,
				-0.10760724159480003, -0.09131298810852691, -0.03521310981103186, 0.058929988388582294, 0.1878028056220495, 0.18767002771550023,
				0.12221268046565888, 0.034918409539180184, -0.045611882331584364, -0.10162273753982914, -0.12314306855543031, -0.10568219197013977,
				-0.04850869060820274, 0.04662376225883174, 0.17640930113316233, 0.1771193344958429, 0.11244084565330154, 0.0258667916627012,
				-0.053997163104184356, -0.10939127078442816, -0.13034056092334664, -0.11235082187028105, -0.054687419338137686, 0.04089891565633101,
				0.17110501545219098, 0.17220476769164184, 0.10788743811408653, 0.021648088348481045, -0.057905674170294436, -0.11301230573754767,
				-0.13369517469864028, -0.11545853168822755, -0.0575663182285551, 0.0382320542788724, 0.1686346412961186, 0.16991646443303948,
				0.1057678409291133, 0.01968480756405483, -0.05972412381121238, -0.11469657064486122, -0.13525512324087177, -0.11690331059859257,
				-0.058904404941257274, 0.0369928030612927, 0.16748694452894614,
			},
			ExpectedHist: []float64{
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
				math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
				-0.44375061609414473, -0.32563112502721026, -0.17143834003766673, -0.0023278111985604727, 0.16727312266967087, 0.3279094224213681,
				0.47375867100780694, -0.036562406238101475, -0.2931464861529588, -0.3765744302496187, -0.3462397430198364, -0.24540188644189123,
				-0.10510077571945724, 0.05281986803736386, 0.2133852454922189, 0.36670484024623184, 0.506609857764188, -0.008558478565361433,
				-0.269111982060538, -0.3558058352038015, -0.32817223922051186, -0.22958104086364461, -0.09116004391671675, 0.06517701394509252,
				0.2243995131899802, 0.37657239279845656, 0.5154912689338689, -0.0005311116261971227, -0.2618293889993654, -0.34917708370591477,
				-0.32212116748305814, -0.22404342083297907, -0.08608132406240468, 0.06984350634116214, 0.2286940054477481, 0.38052981146813786,
				0.5191421554973222, 0.002840133450722343, -0.25871395537016545, -0.34629621596240134, -0.31945581906754217, -0.22157643072097516,
				-0.0837971605556739, 0.07195895621226238, 0.23065361012857344, 0.38234533997787473, 0.5208243991834398, 0.00439900895780343,
				-0.25726931831022115, -0.3449573990624219, -0.31821505007510187, -0.220426526269013, -0.08273147584437052, 0.07294657204165093,
				0.23156885383868978, 0.38319349002971, 0.5216103480689849, 0.005127292547683571, -0.25659449401570467, -0.3443321334602339,
				-0.3176357255010688, -0.21988978733459535, -0.08223421038404219, 0.07340725056911676, 0.23199562262934115, 0.3835888320101999,
				0.5219765658706137,
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
