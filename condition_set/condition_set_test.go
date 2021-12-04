package conditionset

import (
	candle "branch_learning/candle"
	candle_stream "branch_learning/candle_stream"
	condition "branch_learning/condition"
	condition_list "branch_learning/condition/condition_list"
	"testing"
)

func TestConditionSet_AllConditionsApplied(t *testing.T) {
	candle1 := candle.CreateCandle(1.1, 2.2, 3.3, 4.4)
	candle2 := candle.CreateCandle(1.1, 2.2, 3.3, 4.4)
	candles := []candle.Candle{candle1, candle2}
	cstream := candle_stream.CreateCandleStream(candles)
	greenCondition1 := condition_list.GreenCondition{CandleIndex: 1}
	greenCondition2 := condition_list.GreenCondition{CandleIndex: 0}
	conditionsLst := []condition.ICondition{greenCondition2, greenCondition1}
	set := CreateConditionSet(conditionsLst)

	answer := set.DoesApply(&cstream)
	if !answer {
		t.Logf("Expected: %v\tGot: %v", true, answer)
		t.Error("AssertionError")
	}
}
