package candle

import (
	"fmt"
	"os"
)

type Candle struct {
	data map[string]float32
}

func CreateCandle(candleMap map[string]float32) Candle {
	return Candle{candleMap}
}

func (c *Candle) Get(key string) float32 {
	val, ok := c.data[key]
	if !ok {
		fmt.Printf("Missing value of %s\n", key)
		os.Exit(1)

	}
	return val
}
