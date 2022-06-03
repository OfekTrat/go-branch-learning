package broker

import (
	c "branch_learning/candle"
	st "branch_learning/strategy"
)

type Order struct {
	ticker     string
	time       int
	price      float64
	takeProfit float64
	stopLoss   float64
	closeTime  int // Default will be 0 which means there was not close time
}

func MakeOrderFromCandleAndStrategy(ticker string, strategy *st.Strategy, candle c.Candle) Order {
	time := int(candle.Get("mts"))
	price := candle.Get("close")
	takeProfit := price * (1 + (strategy.TakeProfit() / 100))
	stopLoss := price * (1 - (strategy.StopLoss() / 100))

	return MakeOrder(ticker, time, price, takeProfit, stopLoss)
}
func MakeOrder(ticker string, time int, price, takeProfit, stopLoss float64) Order {
	return Order{ticker: ticker, time: time, price: price, takeProfit: takeProfit, stopLoss: stopLoss}
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

func (o Order) Ticker() string {
	return o.ticker
}

func (o Order) Equals(other Order) bool {
	return o.time == other.time
}
