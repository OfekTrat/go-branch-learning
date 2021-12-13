package evolutioner

import (
	bt "branch_learning/backtester"
	candle_stream "branch_learning/candle_stream"
	st "branch_learning/strategy"
)

func backtestGeneration(data *candle_stream.CandleStream, generation []*st.Strategy) []float64 {
	scores := make([]float64, len(generation))
	backtesters := make([]*bt.BackTester, len(generation))

	for i := 0; i < len(generation); i++ {
		backtesters[i] = bt.CreateBackTester(generation[i])
	}

	for i, backtester := range backtesters {
		backtester.Test(data)
		scores[i] = backtester.Score()
	}
	return scores
}
