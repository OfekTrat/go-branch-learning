package condition

import (
	candle "branch_learning/candle"
	candle_stream "branch_learning/candle_stream"
	"testing"
)

func TestRedCondition_GreenCandle(t *testing.T) {
	candleMap := make(map[string]float32)
	candleMap["close"] = 1.1
	candleMap["open"] = 0.4
	c := candle.CreateCandle(candleMap)
	clst := []candle.Candle{c}
	cstream := candle_stream.CreateCandleStream(clst)

	condition := RedCondition{CandleIndex: 0}
	answer := condition.MeetsCondition(cstream)

	if answer {
		t.Logf("Expected: %v\tGot: %v", false, answer)
		t.Error("AssertionError")
	}
}

func TestRedCondition_RedCandle(t *testing.T) {
	candleMap := make(map[string]float32)
	candleMap["close"] = 0.4
	candleMap["open"] = 1.1
	c := candle.CreateCandle(candleMap)
	clst := []candle.Candle{c}
	cstream := candle_stream.CreateCandleStream(clst)

	condition := RedCondition{CandleIndex: 0}
	answer := condition.MeetsCondition(cstream)

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

func TestRedCondition_ValidStreamSize(t *testing.T) {
	rc1 := RedCondition{CandleIndex: 4}
	rc2 := RedCondition{CandleIndex: 2}

	valid1 := rc1.IsValidStreamSize(3)
	valid2 := rc2.IsValidStreamSize(3)

	if valid1 {
		t.Error("AssertionError")
	}
	if !valid2 {
		t.Error("AssertionError")
	}
}
