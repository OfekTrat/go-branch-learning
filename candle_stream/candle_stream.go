package candleStream

import (
	"branch_learning/candle"
)

type CandleStream struct {
	candles []candle.Candle
}

func CreateCandleStream(candles []candle.Candle) *CandleStream {
	return &CandleStream{candles: candles}
}

func LoadStreamFromCsv(csvPath string) *CandleStream {
	return &CandleStream{} // TODO: Implement Function
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
