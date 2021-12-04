package conditionset

import (
	candle_stream "branch_learning/candle_stream"
	condition "branch_learning/condition"
)

type ConditionSet struct {
	conditions []condition.ICondition
}

func CreateConditionSet(conditions []condition.ICondition) ConditionSet {
	return ConditionSet{conditions: conditions}
}

func (cp *ConditionSet) DoesApply(stream *candle_stream.CandleStream) bool {
	for _, cond := range cp.conditions {
		if isApplied := cond.DoesApply(stream); !isApplied {
			return false
		}
	}
	return true
}
