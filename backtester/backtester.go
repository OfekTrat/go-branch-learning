package backtester

import (
	b "branch_learning/broker"
	cst "branch_learning/candle_stream"
	st "branch_learning/strategy"
)

type BackTester struct {
	accountStats *b.AccountStats
	strategy     *st.Strategy
}

func CreateBackTester(strategy *st.Strategy) *BackTester {
	return &BackTester{strategy: strategy, accountStats: b.CreateEmptyAccountStats()}
}

func (bt *BackTester) Strategy() *st.Strategy {
	return bt.strategy
}

func (bt *BackTester) AccountStats() *b.AccountStats {
	return bt.accountStats
}

func (bt *BackTester) Test(stream *cst.CandleStream) {
	windowSize := bt.strategy.WindowSize()
	broker := b.CreateBroker()
	for i := 0; i < stream.Length()-windowSize; i++ {
		slicedStream := stream.GetSlice(i, i+windowSize)
		lastCandle := slicedStream.Get(windowSize - 1)

		broker.ScanOrders(lastCandle.Get("low"), lastCandle.Get("high"))

		if bt.strategy.MeetsConditions(slicedStream) {
			broker.AddOrder(b.MakeOrderFromCandleAndStrategy(bt.strategy, lastCandle))
		}
	}

	bt.accountStats.AddAccountStats(broker.ScanResults())
}
