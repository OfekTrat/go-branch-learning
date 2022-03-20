package strategytester

import (
	b "branch_learning/broker"
	candlestream "branch_learning/candle_stream"
	"branch_learning/configuration"
	l "branch_learning/logger"
	st "branch_learning/strategy"
	"time"
)

var logger *l.Logger = l.CreateLogger()

const TIME_FORMAT = "2006-01-02 15:04:05"

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

	logger.Results.Printf(
		"generation=%d, strategyId=%d, conditionCount=%d, wins=%d, losses=%d, winRate=%f, takeProfit=%f, stopLoss=%f, Score=%f\n",
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
	window := st.strategy.WindowSize()
	broker := b.CreateBroker()

	for i := 0; i < stream.Length()-window; i++ {
		streamSlice := stream.GetSlice(i, i+window)
		lastCandle := streamSlice.Get(window - 1)
		ordersLost, ordersWon := broker.ScanOrders(lastCandle.Get("low"), lastCandle.Get("high"))

		closeTime := int(lastCandle.Get("mts"))
		broker.CloseOrders(closeTime, st.strategy, ordersLost, false)
		broker.CloseOrders(closeTime, st.strategy, ordersWon, true)

		if st.strategy.MeetsConditions(streamSlice) {
			order := b.MakeOrderFromCandleAndStrategy(st.strategy, lastCandle)
			broker.AddOrder(order)

			logger.Orders.Printf(
				"Open [%s] Generation=%d, Strategy=%d, Price=%f\n",
				time.UnixMilli(int64(order.Time())).Format(TIME_FORMAT),
				st.strategy.Generation(),
				st.strategy.Id(),
				order.Price(),
			)
		}
	}
	st.results.AddLosses(broker.ScanResults().Losses())
	st.results.AddWins(broker.ScanResults().Wins())
}
