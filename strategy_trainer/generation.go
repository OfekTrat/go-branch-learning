package strategytrainer

import (
	candlestream "branch_learning/candle_stream"
	st "branch_learning/strategy"
	tester "branch_learning/strategy_tester"
	"sync"
	"time"
)

const (
	WORKERS = 4
)

type generation struct {
	id      int
	testers []*tester.StrategyTester
}

func newGeneration(id int, strategies []*st.Strategy) *generation {
	testers := make([]*tester.StrategyTester, len(strategies))

	for i, strategy := range strategies {
		testers[i] = tester.NewStrategyTester(strategy)
	}
	return &generation{id: id, testers: testers}
}

func (g *generation) test(streams []*candlestream.CandleStream) *generationTestResults {
	var wg sync.WaitGroup
	testersBatch := g.getTesterBatch()

	start := time.Now()

	for i, batch := range testersBatch {
		logger.Info.Printf("Testing Generation=%d, Batch=%d, BatchSize=%d\n", g.id, i, len(batch))

		wg.Add(1)
		go func(i int, batch []*tester.StrategyTester) {
			g.testBatch(batch, streams)
			wg.Done()
			logger.Info.Printf("Done Testing Generation=%d, Batch=%d\n", g.id, i)
		}(i, batch)
	}
	wg.Wait()

	elapsed := time.Since(start).Seconds()
	logger.Info.Printf("Done Testing Generation %d. time=%f\n", g.id, elapsed)

	return newGenerationTestResultsFromStrategyTesters(g.testers)
}

func (g *generation) getTesterBatch() [][]*tester.StrategyTester {
	batches := make([][]*tester.StrategyTester, WORKERS)

	generationSize := len(g.testers)
	jumpSize := generationSize / WORKERS

	for i := 0; i < WORKERS; i++ {
		batches[i] = g.testers[i*jumpSize : (i+1)*jumpSize]
	}
	return batches
}

func (*generation) testBatch(batch []*tester.StrategyTester, streams []*candlestream.CandleStream) {
	for _, strategyTester := range batch {
		strategyTester.Test(streams)
	}
}

func (g *generation) GetStrategyByChance(change float64) *st.Strategy {
	return nil
}

func (g *generation) CalcSumScore() float64 {
	return 0
}
