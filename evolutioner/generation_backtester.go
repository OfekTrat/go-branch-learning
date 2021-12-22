package evolutioner

import (
	bt "branch_learning/backtester"
	candle_stream "branch_learning/candle_stream"
	"sync"
)

func backtestGeneration(streams []*candle_stream.CandleStream, backtesters []*bt.BackTester) []float64 {
	var wg sync.WaitGroup
	nWorkers := 4
	nStrategies := len(backtesters)
	scores := make([]float64, len(backtesters))

	for i := 0; i < nWorkers; i++ {
		wg.Add(1)
		jumpSize := int(nStrategies / nWorkers)
		go func(i int) {
			defer wg.Done()
			worker(streams, backtesters[i*jumpSize:(i+1)*jumpSize])
		}(i)
	}
	wg.Wait()
	for i, backtester := range backtesters {
		scores[i] = backtester.Score()
	}
	return scores
}

func worker(streams []*candle_stream.CandleStream, backtesters []*bt.BackTester) {
	for _, backtester := range backtesters {
		backtestStrategy(backtester, streams)
	}
}

func backtestStrategy(backtester *bt.BackTester, streams []*candle_stream.CandleStream) {
	for _, stream := range streams {
		backtester.Test(stream)
	}
}
