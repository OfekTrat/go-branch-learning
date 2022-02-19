package evolutioner

import (
	broker "branch_learning/broker"
	c "branch_learning/candle"
	st "branch_learning/strategy"
)

func getOrder(strategyId int, strategy *st.Strategy, candle c.Candle) broker.Order {
	mts := int(candle.Get("mts"))
	close := candle.Get("close")
	stopLoss := (100 - strategy.StopLoss()) * close
	takeProfit := (100 + strategy.TakeProfit()) * close
	return broker.MakeOrder(mts, strategyId, close, takeProfit, stopLoss)
}

// TODO: Implement the following things
// 1. Function that gets an order from single strategy on single slice
// 2. Function that runs the function in section 1 in a multiprocessing mechanism and updates broker.
