package pipeline

import (
	candle_stream "branch_learning/candle_stream"
	condition "branch_learning/condition"
)

type CandlePipeline struct {
	conditions []condition.ICondition
}

func CreateCandlePipeline(conditions []condition.ICondition) CandlePipeline {
	return CandlePipeline{conditions: conditions}
}

func (cp *CandlePipeline) DoesApply(stream *candle_stream.CandleStream) bool {
	for _, cond := range cp.conditions {
		if isApplied := cond.DoesApply(stream); !isApplied {
			return false
		}
	}
	return true
}
