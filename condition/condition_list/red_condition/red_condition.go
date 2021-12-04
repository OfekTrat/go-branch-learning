package condition

import candleStream "branch_learning/candle_stream"

const (
	ConditionType = "RedCondition"
)

type RedCondition struct {
	CandleIndex int `json:"candle_index"`
}

func (condition RedCondition) DoesApply(stream *candleStream.CandleStream) bool {
	candle := stream.Get(condition.CandleIndex)
	return candle.Close() < candle.Open()
}

func (condition RedCondition) ConditionType() string {
	return ConditionType
}
