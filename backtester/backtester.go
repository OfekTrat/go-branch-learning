package backtester

import (
	cst "branch_learning/candle_stream"
	om "branch_learning/order_manager"
	st "branch_learning/strategy"
	ts "branch_learning/trade_stats"
)

type BackTester struct {
	orderMananger *om.OrderManager
	tradeStats    *ts.TradeStats
	strategy      *st.Strategy
}

func CreateBackTester(strategy *st.Strategy) *BackTester {
	manager := om.OrderManager{}
	stats := ts.TradeStats{}
	return &BackTester{strategy: strategy, orderMananger: &manager, tradeStats: &stats}
}

func (bt *BackTester) Stats() *ts.TradeStats {
	return bt.tradeStats
}
func (bt *BackTester) Strategy() *st.Strategy {
	return bt.strategy
}

func (bt *BackTester) Score() float64 {
	return bt.tradeStats.Score()
}

func (bt *BackTester) Test(stream *cst.CandleStream) {
	var metCondition bool
	windowSize := bt.strategy.WindowSize()

	logger.Debugf("Testing Strategy: Window=%v, TakeProfit/StopLoss=%v/%v", bt.strategy.WindowSize(),
		bt.strategy.TakeProfit(), bt.strategy.StopLoss())

	for i := 0; i < stream.Length()-windowSize; i++ {
		slicedStream := stream.GetSlice(i, i+windowSize)
		lastCandle := slicedStream.Get(windowSize - 1)

		wins, losses := bt.orderMananger.CheckExits(lastCandle.Get("high"))

		for k := 0; k < wins; k++ {
			bt.tradeStats.AddWin()
		}
		for k := 0; k < losses; k++ {
			bt.tradeStats.AddLoss()
		}

		metCondition = bt.strategy.MeetsConditions(slicedStream)

		if !metCondition {
			continue
		}

		exit := bt.strategy.GetExit(lastCandle.Get("close"))
		bt.orderMananger.AddExit(exit)
	}

	logger.Debugf("Finished Testing: Wins=%v, Loss=%v", bt.tradeStats.Wins(), bt.tradeStats.Losses())
}
