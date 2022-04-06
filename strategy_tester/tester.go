package strategytester

import (
	b "branch_learning/broker"
	candlestream "branch_learning/candle_stream"
	c "branch_learning/configuration"
	l "branch_learning/logger"
	st "branch_learning/strategy"
)

var logger *l.Logger = l.CreateLogger()
var configuration *c.Configuration = c.GetConfiguration()

const TIME_FORMAT = "2006-01-02 15:04:05"

type StrategyTester struct {
	strategy *st.Strategy
	results  *TestResults
}

func NewStrategyTester(strategy *st.Strategy) *StrategyTester {
	results := CreateTestResultsFromStrategy(strategy)
	return &StrategyTester{strategy: strategy, results: results}
}

func NewStrategyFromConfiguration() *StrategyTester {
	strategy := st.CreateStrategyFromFile(configuration.Strategy())
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

	logger.Results.Printf(
		"%d,%d,%d,%d,%d,%f,%f,%f,%f\n",
		st.strategy.Generation(),
		st.strategy.Id(),
		st.results.ConditionCount,
		st.results.Wins,
		st.results.Losses,
		float32(st.results.Wins)/float32(st.results.Wins+st.results.Losses),
		st.results.TakeProfit,
		st.results.StopLoss,
		st.results.Score,
	)
}

func (st *StrategyTester) testSingleCandleStream(stream *candlestream.CandleStream) {
	ticker := stream.Name()
	window := st.strategy.WindowSize()
	broker := b.CreateBroker()

	for i := 0; i < stream.Length()-window; i++ {
		streamSlice := stream.GetSlice(i, i+window)
		lastCandle := streamSlice.Get(window - 1)
		closeTime := int(lastCandle.Get("mts"))
		ordersLost, ordersWon := broker.ScanOrders(lastCandle.Get("low"), lastCandle.Get("high"))

		broker.CloseLossOrders(closeTime, st.strategy, ordersLost)
		broker.CloseWinOrders(closeTime, st.strategy, ordersWon)

		if st.strategy.MeetsConditions(streamSlice) {
			order := b.MakeOrderFromCandleAndStrategy(ticker, st.strategy, lastCandle)
			broker.AddOrder(order)

			logger.LogOrder(
				"%s,%d,%d,%d,%d,%f,%d\n",
				order.Ticker(),
				order.Time(),
				st.strategy.Generation(),
				st.strategy.Id(),
				0,
				order.Price(),
				0,
			)
		}
	}
	st.results.AddLosses(broker.ScanResults().Losses())
	st.results.AddWins(broker.ScanResults().Wins())
}
