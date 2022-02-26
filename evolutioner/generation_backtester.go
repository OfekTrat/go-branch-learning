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
			bt.TestMultipleBacktesters(backtesters[i*jumpSize:(i+1)*jumpSize], streams)
		}(i)
	}
	wg.Wait()
	for i, backtester := range backtesters {
		scores[i] = backtester.Score()
	}
	return scores
}
