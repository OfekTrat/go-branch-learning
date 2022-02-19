package broker

import (
	c "branch_learning/candle"
	st "branch_learning/strategy"
)

type Order struct {
	time       int
	price      float32
	takeProfit float32
	stopLoss   float32
}

func MakeOrderFromCandleAndStrategy(strategy *st.Strategy, candle c.Candle) Order {
	time := int(candle.Get("mts"))
	price := candle.Get("close")
	takeProfit := price * (1 + (strategy.TakeProfit() / 100))
	stopLoss := price * (1 - (strategy.StopLoss() / 100))

	return MakeOrder(time, price, takeProfit, stopLoss)
}
func MakeOrder(time int, price, takeProfit, stopLoss float32) Order {
	return Order{time: time, price: price, takeProfit: takeProfit, stopLoss: stopLoss}
}

func (o Order) Time() int {
	return o.time
}

func (o Order) Price() float32 {
	return o.price
}

func (o Order) TakeProfit() float32 {
	return o.takeProfit
}

func (o Order) StopLoss() float32 {
	return o.stopLoss
}

func (o Order) Equals(other Order) bool {
	return o.time == other.time
}