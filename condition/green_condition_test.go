package condition

import (
	candle "branch_learning/candle"
	candle_stream "branch_learning/candle_stream"
	"testing"
)

func TestGreenCondition_GreenCandle(t *testing.T) {
	candleMap := make(map[string]float32)
	candleMap["close"] = 1.1
	candleMap["open"] = 0.4
	c := candle.CreateCandle(candleMap)
	cstream := candle_stream.CreateCandleStream([]candle.Candle{c})
	greenCondition := GreenCondition{CandleIndex: 0}

	if answer := greenCondition.MeetsCondition(cstream); !answer {
		t.Logf("Expected: %v\tGot: %v", true, answer)
		t.Error("AssertionError")
	}
}

func TestGreenCondition_RedCandle(t *testing.T) {
	candleMap := make(map[string]float32)
	candleMap["close"] = 0.4
	candleMap["open"] = 1.1
	c := candle.CreateCandle(candleMap)
	cstream := candle_stream.CreateCandleStream([]candle.Candle{c})
	greenCondition := GreenCondition{CandleIndex: 0}

	if answer := greenCondition.MeetsCondition(cstream); answer {
		t.Logf("Expected: %v\tGot: %v", false, answer)
		t.Error("AssertionError")
	}
}

func TestGreenCondition_ConditionType(t *testing.T) {
	expectedType := "GreenCondition"
	greenCondition := GreenCondition{CandleIndex: 0}
	if condType := greenCondition.ConditionType(); condType != expectedType {
		t.Logf("Expected: %v\tGot: %v", expectedType, condType)
		t.Error("AssertionError")
	}
}

func TestGreenCondition_IsValidStreamSize(t *testing.T) {
	gc1 := GreenCondition{CandleIndex: 4}
	gc2 := GreenCondition{CandleIndex: 2}

	valid1 := gc1.IsValidStreamSize(3)
	valid2 := gc2.IsValidStreamSize(3)

	if valid1 {
		t.Error("AssertionError")
	}
	if !valid2 {
		t.Error("AssertionError")
	}
}

func TestGreenCondition_Mutate(t *testing.T) {
	rc := GreenCondition{CandleIndex: 4}
	rcMutation := rc.Mutate(1000).(GreenCondition)

	if rcMutation.CandleIndex == 4 {
		t.Error("AssertionError")
	}
	if rc.CandleIndex != 4 {
		t.Log("Changed the type itself")
		t.Error("AssertionError")
	}
}
