package strategytester

import (
	b "branch_learning/broker"
	candlestream "branch_learning/candle_stream"
	"branch_learning/configuration"
	l "branch_learning/logger"
	st "branch_learning/strategy"
)

var logger *l.Logger = l.CreateLogger()

type StrategyTester struct {
	strategy *st.Strategy
	results  *TestResults
}

func NewStrategyTester(strategy *st.Strategy) *StrategyTester {
	results := CreateTestResultsFromStrategy(strategy)
	return &StrategyTester{strategy: strategy, results: results}
}

func NewStrategyFromConfiguration(testConfiguration configuration.TestConfiguration) *StrategyTester {
	strategy := st.CreateStrategyFromFile(testConfiguration.Strategy)
	results := CreateTestResultsFromStrategy(strategy)
	return &StrategyTester{strategy: strategy, results: results}
}

func (st *StrategyTester) Strategy() *st.Strategy {
	return st.strategy
}

func (st *StrategyTester) Results() *TestResults {
	return st.results
}

func (st *StrategyTester) Test(streams []*candlestream.CandleStream) {
	for _, stream := range streams {
		st.testSingleCandleStream(stream)
	}

	st.results.CalcScore()
}

func (st *StrategyTester) testSingleCandleStream(stream *candlestream.CandleStream) {
	window := st.strategy.WindowSize()
	broker := b.CreateBroker()

	for i := 0; i < stream.Length()-window; i++ {
		streamSlice := stream.GetSlice(i, i+window)
		lastCandle := streamSlice.Get(window - 1)
		broker.ScanOrders(lastCandle.Get("low"), lastCandle.Get("high"))

		if st.strategy.MeetsConditions(streamSlice) {
			order := b.MakeOrderFromCandleAndStrategy(st.strategy, lastCandle)
			broker.AddOrder(order)
		}
	}
	st.results.AddLosses(broker.ScanResults().Losses())
	st.results.AddWins(broker.ScanResults().Wins())
}
