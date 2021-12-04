package strategy

type Exit struct {
	takeProfitPrice float32
	stopLossPrice   float32
}

func (e Exit) TakeProfit() float32 {
	return e.takeProfitPrice
}

func (e Exit) StopLoss() float32 {
	return e.stopLossPrice
}

func (e Exit) Take(price float32) bool {
	return e.takeProfitPrice < price
}
func (e Exit) Stop(price float32) bool {
	return e.stopLossPrice > price
}
