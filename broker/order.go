package broker

import (
	c "branch_learning/candle"
	st "branch_learning/strategy"
)

type Order struct {
	time       int
	price      float64
	takeProfit float64
	stopLoss   float64
}

func MakeOrderFromCandleAndStrategy(strategy *st.Strategy, candle c.Candle) Order {
	time := int(candle.Get("mts"))
	price := candle.Get("close")
	takeProfit := price * (1 + (strategy.TakeProfit() / 100))
	stopLoss := price * (1 - (strategy.StopLoss() / 100))

	return MakeOrder(time, price, takeProfit, stopLoss)
}
func MakeOrder(time int, price, takeProfit, stopLoss float64) Order {
	return Order{time: time, price: price, takeProfit: takeProfit, stopLoss: stopLoss}
}

func (o Order) Time() int {
	return o.time
}

func (o Order) Price() float64 {
	return o.price
}

func (o Order) TakeProfit() float64 {
	return o.takeProfit
}

func (o Order) StopLoss() float64 {
	return o.stopLoss
}

func (o Order) Equals(other Order) bool {
	return o.time == other.time
}
