package condition

import (
	candleStream "branch_learning/candle_stream"
	"math/rand"
)

const (
	RedConditionType = "RedCondition"
)

type RedCondition struct {
	CandleIndex int `json:"candle_index"`
}

func (c RedCondition) MeetsCondition(stream *candleStream.CandleStream) bool {
	candle := stream.Get(c.CandleIndex)
	return candle.Get("close") < candle.Get("open")
}

func (c RedCondition) ConditionType() string {
	return RedConditionType
}

func (c RedCondition) IsValidStreamSize(streamsize int) bool {
	return c.CandleIndex <= streamsize
}

func (c RedCondition) Mutate(streamsize int) ICondition {
	c.CandleIndex = rand.Intn(streamsize)
	return c
}
