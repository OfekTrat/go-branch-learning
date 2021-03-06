package strategy

import (
	"branch_learning/candle"
	candle_stream "branch_learning/candle_stream"
	"branch_learning/condition"
	"testing"
)

func TestStrategy_MeetsConditions(t *testing.T) {
	redCandleMap := make(map[string]float64)
	redCandleMap["close"] = 9
	redCandleMap["open"] = 10
	greenCandleMap := make(map[string]float64)
	greenCandleMap["close"] = 10
	greenCandleMap["open"] = 9

	redCandle := candle.CreateCandle(redCandleMap)
	greenCandle := candle.CreateCandle(greenCandleMap)
	candleStream := candle_stream.CreateCandleStream("test", []candle.Candle{redCandle, greenCandle})

	redCond := condition.DummyCondition{CandleIndex: 0, IsGreen: false}
	greenCond := condition.DummyCondition{CandleIndex: 1, IsGreen: true}
	conditions := condition.CreateConditions([]condition.ICondition{redCond, greenCond})

	s := CreateStrategy(0, 0, 10, 1, 1, conditions)
	answer := s.MeetsConditions(candleStream)

	if !answer {
		t.Logf("Expected %v\tGot: %v", true, answer)
		t.Error("AssertionError")
	}
}

func TestStrategy_GettingConditions(t *testing.T) {
	redCondition := condition.DummyCondition{CandleIndex: 1, IsGreen: false}
	redCondition2 := condition.DummyCondition{CandleIndex: 4, IsGreen: false}
	s := CreateStrategy(0, 0, 10, 1, 1, condition.CreateConditions([]condition.ICondition{redCondition, redCondition2}))

	conditions1 := s.Conditions()
	conditions2 := s.Conditions()

	conditions1.SetInIndex(condition.DummyCondition{CandleIndex: 2, IsGreen: true}, 0)

	if conditions1.GetByIndex(0).Equals(conditions2.GetByIndex(0)) {
		t.Error("AssertionError")
	}
}
