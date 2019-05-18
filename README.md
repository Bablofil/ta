# ta
Functions for technical analysis written in go

## Install
`go get github.com/bablofil/ta`

## Usage

```
package main

import (
	"fmt"

	"github.com/bablofil/ta"
)

func main() {
	var closes []float64 = []float64{
		1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0,
		1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0,
		1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0,
	}

	// run SMA with period 2
	fmt.Println(ta.SMA(closes, 2))

	// run SMA with period 1
	fmt.Println(ta.SMA(closes, 1))

	// run EMA with period 10 (the same is for SMMA, MMA, RMA)
	fmt.Println(ta.EMA(closes, 10))
	
	// MACD fast=12, slow=26, smoothing=9
	macd, macdsignal, macdhist, err := ta.MACD(closes, 12, 26, 9)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println("macd", macd)
	fmt.Println("macdsignal", macdsignal)
	fmt.Println("macdhist", macdhist)
	
	// RSI period 9
	fmt.Println(RSI(closes, 9))
	
	// STOCH (high, low, close, fastk_period=5, slowk_period=3, slowd_period=3
	slowk, slowd, err := ta.STOCH(high, low, closes, 5, 3, 3)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println("slowk", slowk)
	fmt.Println("slowd", slowd)
	
	// STOCHRSI (close, period, fastk_period, fastd_period )
	slowk, slowd, err := ta.STOCHRSI(close, 14, 3, 3)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println("slowk", slowk)
	fmt.Println("slowd", slowd)
}
```
## Benchmarks

```
goos: windows
goarch: amd64
BenchmarkSMA-4           1000000              1122 ns/op
BenchmarkEMA-4            500000              2260 ns/op
BenchmarkMACD-4           200000              7605 ns/op
BenchmarkRSI-4            200000              6170 ns/op
BenchmarkSTOCH-4          200000             10305 ns/op
BenchmarkSTOCHRSI-4       100000             21041 ns/op
PASS
ok      _/D_/golang/ta/tests    9.738s
```
