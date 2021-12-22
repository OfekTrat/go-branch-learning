package exit

type Exit struct {
	takeProfitPrice float32
	stopLossPrice   float32
}

func CreateExit(takeProfitPrice float32, stopLossPrice float32) Exit {
	return Exit{takeProfitPrice: takeProfitPrice, stopLossPrice: stopLossPrice}
}

func (e Exit) TakeProfitPercentage() float32 {
	return e.takeProfitPrice
}

func (e Exit) StopLossPercentage() float32 {
	return e.stopLossPrice
}

func (e Exit) IsTake(price float32) bool {
	return e.takeProfitPrice <= price
}
func (e Exit) IsStop(price float32) bool {
	return e.stopLossPrice >= price
}
