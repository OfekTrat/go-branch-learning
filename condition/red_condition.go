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

func (c RedCondition) Mutate() ICondition {
	c.CandleIndex += int((float32(rand.Intn(2)) - 0.5) * 2)
	return c
}

func CreateRandomRedCondition(streamsize int) ICondition {
	return RedCondition{CandleIndex: rand.Intn(streamsize)}
}
