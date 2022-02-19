package backtester

import (
	b "branch_learning/broker"
	cst "branch_learning/candle_stream"
	st "branch_learning/strategy"
	"math"
)

type BackTester struct {
	broker   *b.Broker
	strategy *st.Strategy
}

func CreateBackTester(strategy *st.Strategy) *BackTester {
	broker := b.CreateBroker()
	return &BackTester{broker: broker, strategy: strategy}
}

func (bt *BackTester) Strategy() *st.Strategy {
	return bt.strategy
}

func (bt *BackTester) Broker() *b.Broker {
	return bt.broker
}

func (bt *BackTester) Score() float64 {
	sumOrders := bt.broker.ScanResults().Wins() + bt.broker.ScanResults().Losses()

	if sumOrders == 0 {
		return 0
	}
	numberOfConditionWeight := calcConditionLengthWeight(bt.strategy.Conditions().Length())
	winRate := float32(bt.broker.ScanResults().Wins()) / float32(sumOrders)
	lossRate := float32(bt.broker.ScanResults().Losses()) / float32(sumOrders)
	totalEstimatedEarningsForHundredOrders := (winRate * bt.strategy.TakeProfit()) - (lossRate * bt.strategy.StopLoss())
	sumOrdersWeight := calcSumOrdersWeight(sumOrders)

	return float64(totalEstimatedEarningsForHundredOrders) * sumOrdersWeight * numberOfConditionWeight
}

func calcSumOrdersWeight(sumOrders int) float64 {
	/*
		kind of sigmoid function aims for 3% of number of orders
		   4
		1 + e^(-0.005*sumOrders)
		minus 2

		It should be changed to be relative to the given data size (if possible)
	*/

	return 4/(1+math.Pow(float64(math.E), -0.005*float64(sumOrders))) - 2
}

func calcConditionLengthWeight(numberOfConditions int) float64 {
	/*
		This function gives wait to the number of conditions of a strategy.
		The reason for doing that is to keep the strategies simple and not get too much crazy.
	*/
	threshold := 100.0
	slope := -0.02
	if float64(numberOfConditions) <= threshold {
		return float64(1)
	} else {
		return slope*float64(numberOfConditions) + (1 - (slope * threshold))
	}
}

func (bt *BackTester) Test(stream *cst.CandleStream) {
	windowSize := bt.strategy.WindowSize()

	for i := 0; i < stream.Length()-windowSize; i++ {
		slicedStream := stream.GetSlice(i, i+windowSize)
		lastCandle := slicedStream.Get(windowSize - 1)

		bt.broker.ScanOrders(lastCandle.Get("low"), lastCandle.Get("high"))

		if bt.strategy.MeetsConditions(slicedStream) {
			bt.broker.AddOrder(b.MakeOrderFromCandleAndStrategy(bt.strategy, lastCandle))
		}
	}
}
