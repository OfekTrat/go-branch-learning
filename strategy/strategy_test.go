package strategy

import (
	"branch_learning/candle"
	candle_stream "branch_learning/candle_stream"
	"branch_learning/condition"
	"testing"
)

func TestStrategy_MeetsConditions(t *testing.T) {
	redCandleMap := make(map[string]float32)
	redCandleMap["close"] = 9
	redCandleMap["open"] = 10
	greenCandleMap := make(map[string]float32)
	greenCandleMap["close"] = 10
	greenCandleMap["open"] = 9

	redCandle := candle.CreateCandle(redCandleMap)
	greenCandle := candle.CreateCandle(greenCandleMap)
	candleStream := candle_stream.CreateCandleStream([]candle.Candle{redCandle, greenCandle})

	redCond := condition.DummyCondition{CandleIndex: 0, IsGreen: false}
	greenCond := condition.DummyCondition{CandleIndex: 1, IsGreen: true}
	conditions := condition.CreateConditions([]condition.ICondition{redCond, greenCond})

	s := CreateStrategy(10, 1, 1, conditions)
	answer := s.MeetsConditions(candleStream)

	if !answer {
		t.Logf("Expected %v\tGot: %v", true, answer)
		t.Error("AssertionError")
	}
}

func TestStrategy_GetExit(t *testing.T) {
	s := CreateStrategy(10, 1, 1, condition.CreateConditions([]condition.ICondition{}))
	exit := s.GetExit(100)
	if exit.StopLossPercentage() != 99 || exit.TakeProfitPercentage() != 101 {
		t.Error("AssertionError")
	}
}

func TestStrategy_GettingConditions(t *testing.T) {
	redCondition := condition.DummyCondition{CandleIndex: 1, IsGreen: false}
	redCondition2 := condition.DummyCondition{CandleIndex: 4, IsGreen: false}
	s := CreateStrategy(10, 1, 1, condition.CreateConditions([]condition.ICondition{redCondition, redCondition2}))

	conditions1 := s.Conditions().ToList()
	conditions2 := s.Conditions().ToList()

	conditions1[0] = condition.DummyCondition{CandleIndex: 2, IsGreen: true}

	if conditions1[0] == conditions2[0] {
		t.Error("AssertionError")
	}
}
