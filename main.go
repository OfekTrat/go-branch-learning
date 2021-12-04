package main

import (
	candle "branch_learning/candle"
	candleStream "branch_learning/candle_stream"
	"fmt"
)

func main() {
	c := candle.CreateCandle(1.1, 2.2, 3.3, 4.4)
	c2 := candle.CreateCandle(4.4, 3.3, 2.2, 1.1)
	cs := candleStream.CreateCandleStream([]candle.Candle{c, c2})
	fmt.Println(cs.Get(0))
	fmt.Println(cs.Get(1))

}
