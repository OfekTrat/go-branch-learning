package condition

import (
	stream "branch_learning/candle_stream"
)

const (
	GreenConditionType = "GreenCondition"
)

type GreenCondition struct {
	CandleIndex int `json:"GreenCondition"`
}

func (c GreenCondition) DoesApply(stream *stream.CandleStream) bool {
	candle := stream.Get(c.CandleIndex)
	return candle.Close() > candle.Open()
}

func (c GreenCondition) ConditionType() string {
	return GreenConditionType
}
