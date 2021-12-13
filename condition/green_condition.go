package condition

import (
	stream "branch_learning/candle_stream"
	"math/rand"
)

const (
	GreenConditionType = "GreenCondition"
)

type GreenCondition struct {
	CandleIndex int `json:"GreenCondition"`
}

func (c GreenCondition) MeetsCondition(stream *stream.CandleStream) bool {
	candle := stream.Get(c.CandleIndex)
	return candle.Get("close") > candle.Get("open")
}

func (c GreenCondition) ConditionType() string {
	return GreenConditionType
}

func (c GreenCondition) IsValidStreamSize(streamsize int) bool {
	return c.CandleIndex <= streamsize
}

func (c GreenCondition) Mutate(streamsize int) ICondition {
	c.CandleIndex = rand.Intn(streamsize)
	return c
}
