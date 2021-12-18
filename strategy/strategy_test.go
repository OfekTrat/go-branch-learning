package strategy

import (
	"branch_learning/candle"
	candle_stream "branch_learning/candle_stream"
	condition "branch_learning/condition"
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

	redCond := condition.CandleTypeCondition{CandleIndex: 0, IsGreen: false}
	greenCond := condition.CandleTypeCondition{CandleIndex: 1, IsGreen: true}

	s := CreateStrategy(10, 1, 1, []condition.ICondition{redCond, greenCond})
	answer := s.MeetsConditions(candleStream)

	if !answer {
		t.Logf("Expected %v\tGot: %v", true, answer)
		t.Error("AssertionError")
	}
}

func TestStrategy_GetExit(t *testing.T) {
	s := CreateStrategy(10, 1, 1, []condition.ICondition{})
	exit := s.GetExit(100)
	if exit.StopLossPercentage() != 99 || exit.TakeProfitPercentage() != 101 {
		t.Error("AssertionError")
	}
}

func TestStrategy_GettingConditions(t *testing.T) {
	redCondition := condition.CandleTypeCondition{CandleIndex: 1, IsGreen: false}
	redCondition2 := condition.CandleTypeCondition{CandleIndex: 4, IsGreen: false}
	s := CreateStrategy(10, 1, 1, []condition.ICondition{redCondition, redCondition2})

	conditions1 := s.Conditions()
	conditions2 := s.Conditions()

	conditions1[0] = condition.CandleTypeCondition{CandleIndex: 2, IsGreen: true}

	if conditions1[0] == conditions2[0] {
		t.Error("AssertionError")
	}
}
