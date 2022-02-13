package broker

type Order struct {
	time       int
	strategyId int
	price      float32
	takeProfit float32
	stopLoss   float32
}

func MakeOrder(time int, strategyId int, price, takeProfit, stopLoss float32) Order {
	return Order{time, strategyId, price, takeProfit, stopLoss}
}

func (o Order) Time() int {
	return o.time
}

func (o Order) StrategyId() int {
	return o.strategyId
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
	return o.time == other.time && o.strategyId == other.strategyId
}
