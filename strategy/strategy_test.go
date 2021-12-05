package strategy

import (
	"branch_learning/candle"
	candle_stream "branch_learning/candle_stream"
	condition "branch_learning/condition"
	conditions "branch_learning/condition/condition_list"
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

	redCond := conditions.RedCondition{CandleIndex: 0}
	greenCond := conditions.GreenCondition{CandleIndex: 1}

	strategy := CreateStrategy(10, 1, 1, []condition.ICondition{redCond, greenCond})
	answer := strategy.MeetsConditions(candleStream)

	if !answer {
		t.Logf("Expected %v\tGot: %v", true, answer)
		t.Error("AssertionError")
	}
}

func TestStrategy_GetExit(t *testing.T) {
	strategy := CreateStrategy(10, 1, 1, []condition.ICondition{})
	exit := strategy.GetExit(100)
	if exit.StopLossPercentage() != 99 || exit.TakeProfitPercentage() != 101 {
		t.Error("AssertionError")
	}
}
