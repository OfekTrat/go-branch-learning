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
	sumOrders := bt.tradeStats.Wins() + bt.tradeStats.Losses()

	if sumOrders == 0 {
		return 0
	}
	numberOfConditionWeight := calcConditionLengthWeight(bt.strategy.Conditions().Length())
	winRate := float32(bt.tradeStats.Wins()) / float32(sumOrders)
	lossRate := float32(bt.tradeStats.Losses()) / float32(sumOrders)
	totalEstimatedEarningsForHundredOrders := (winRate * bt.strategy.TakeProfit()) - (lossRate * bt.strategy.StopLoss())
	sumOrdersWeight := calcSumOrdersWeight(sumOrders)

	return float64(totalEstimatedEarningsForHundredOrders) * sumOrdersWeight * numberOfConditionWeight
}

func calcSumOrdersWeight(sumOrders int) float64 {
	// kind of sigmoid function aims for 3% of number of orders
	//    4
	// 1 + e^(-0.005*sumOrders)
	// minus 2

	// It should be changed to be relative to the given data size (if possible)

	return 4/(1+math.Pow(float64(math.E), -0.005*float64(sumOrders))) - 2
}

func calcConditionLengthWeight(numberOfConditions int) float64 {
	threshold := 100.0
	slope := -0.02
	if float64(numberOfConditions) <= threshold {
		return float64(1)
	} else {
		return slope*float64(numberOfConditions) + (1 - (slope * threshold))
	}
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
