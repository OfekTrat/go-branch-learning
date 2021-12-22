package evolutioner

import (
	bt "branch_learning/backtester"
	candle_stream "branch_learning/candle_stream"
)

func backtestGeneration(data []*candle_stream.CandleStream, backtesters []*bt.BackTester) []float64 {
	scores := make([]float64, len(backtesters))
	for i, backtester := range backtesters {
		for _, stream := range data {
			backtester.Test(stream)
		}
		scores[i] = backtester.Score()
	}
	return scores
}
