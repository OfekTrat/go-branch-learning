package candle

type Candle struct {
	open  float32
	high  float32
	low   float32
	close float32
}

func CreateCandle(open float32, high float32, low float32, close float32) Candle {
	return Candle{
		open:  open,
		high:  high,
		low:   low,
		close: close,
	}
}

func (c Candle) Open() float32 {
	return c.open
}

func (c Candle) Close() float32 {
	return c.close
}

func (c Candle) High() float32 {
	return c.high
}

func (c Candle) Low() float32 {
	return c.low
}
