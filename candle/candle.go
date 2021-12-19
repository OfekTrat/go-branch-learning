package candle

import "strings"

type Candle struct {
	data map[string]float32
}

func CreateCandle(candleMap map[string]float32) Candle {
	return Candle{candleMap}
}

func (c *Candle) Get(key string) float32 {
	return c.data[strings.ToLower(key)]
}
