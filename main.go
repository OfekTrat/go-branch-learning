package main

import (
	candle_stream "branch_learning/candle_stream"
)

func main() {
	candle_stream.LoadCandleStreamFromCsv("./candle_stream/test_data/test.csv")
}
