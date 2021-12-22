package backtester

import (
	cst "branch_learning/candle_stream"
	st "branch_learning/strategy"
	ts "branch_learning/trade_stats"
	"math"
)

type BackTester struct {
	orderMananger *OrderManager
	tradeStats    *ts.TradeStats
	strategy      *st.Strategy
}

func CreateBackTester(strategy *st.Strategy) *BackTester {
	manager := OrderManager{}
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
	stopLossPercentage := float64(bt.tradeStats.Losses()) * float64(bt.strategy.StopLoss())
	takeProfitPercentage := float64(bt.tradeStats.Wins()) * float64(bt.strategy.TakeProfit())
	power := 4/(1+math.Pow(float64(math.E), -0.005*float64(bt.tradeStats.Losses()+bt.tradeStats.Wins()))) - 2 // kind of sigmoid function
	if bt.tradeStats.Losses() == 0 {
		stopLossPercentage = 0.33
	}

	return (takeProfitPercentage / stopLossPercentage) * power
}

func (bt *BackTester) Test(stream *cst.CandleStream) {
	var metCondition bool
	windowSize := bt.strategy.WindowSize()

	for i := 0; i < stream.Length()-windowSize; i++ {
		slicedStream := stream.GetSlice(i, i+windowSize)
		lastCandle := slicedStream.Get(windowSize - 1)

		wins, losses := bt.orderMananger.CheckExits(lastCandle.Get("high"), lastCandle.Get("low"))

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
}
