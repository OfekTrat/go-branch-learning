package condition

import (
	candle "branch_learning/candle"
	candle_stream "branch_learning/candle_stream"
	"testing"
)

func TestRedCondition_GreenCandle(t *testing.T) {
	c := candle.CreateCandle(1.1, 2.2, 3.3, 4.4)
	clst := []candle.Candle{c}
	cstream := candle_stream.CreateCandleStream(clst)

	condition := RedCondition{CandleIndex: 0}
	answer := condition.DoesApply(&cstream)

	if answer {
		t.Logf("Expected: %v\tGot: %v", false, answer)
		t.Error("AssertionError")
	}
}

func TestRedCondition_RedCandle(t *testing.T) {
	c := candle.CreateCandle(1.1, 2.2, 0.3, 0.4)
	clst := []candle.Candle{c}
	cstream := candle_stream.CreateCandleStream(clst)

	condition := RedCondition{CandleIndex: 0}
	answer := condition.DoesApply(&cstream)

	if !answer {
		t.Logf("Expected: %v\tGot: %v", true, answer)
		t.Error("AssertionError")
	}
}

func TestRedCondition_ConditionType(t *testing.T) {
	expectedType := "RedCondition"
	condition := RedCondition{CandleIndex: 2}
	if condType := condition.ConditionType(); condType != expectedType {
		t.Logf("Expected %s\tGot: %s", expectedType, condType)
		t.Errorf("AssertionError")
	}
}
