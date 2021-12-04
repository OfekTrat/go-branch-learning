package condition

import candleStream "branch_learning/candle_stream"

const (
	RedConditionType = "RedCondition"
)

type RedCondition struct {
	CandleIndex int `json:"candle_index"`
}

func (condition RedCondition) MeetsCondition(stream *candleStream.CandleStream) bool {
	candle := stream.Get(condition.CandleIndex)
	return candle.Close() < candle.Open()
}

func (condition RedCondition) ConditionType() string {
	return RedConditionType
}
