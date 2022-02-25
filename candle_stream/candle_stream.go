package candlestream

import (
	"branch_learning/candle"
)

type CandleStream struct {
	name    string
	candles []candle.Candle
}

func CreateCandleStream(name string, candles []candle.Candle) *CandleStream {
	return &CandleStream{name: name, candles: candles}
}

func (stream *CandleStream) Name() string {
	return stream.name
}

func (stream *CandleStream) Get(candleIndex int) candle.Candle {
	return stream.candles[candleIndex]
}

func (stream *CandleStream) Length() int {
	return len(stream.candles)
}

func (stream *CandleStream) GetSlice(firstIndex int, lastIndex int) *CandleStream {
	return &CandleStream{candles: stream.candles[firstIndex:lastIndex]}
}
