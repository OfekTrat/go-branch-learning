package candleStream

import (
	"branch_learning/candle"
)

type CandleStream struct {
	candles []candle.Candle
}

func CreateCandleStream(candles []candle.Candle) CandleStream {
	return CandleStream{candles: candles}
}

func (candleStream *CandleStream) Get(candleIndex int) candle.Candle {
	return candleStream.candles[candleIndex]
}
