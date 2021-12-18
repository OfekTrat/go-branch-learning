package backtester

import (
	"branch_learning/candle"
	candle_stream "branch_learning/candle_stream"
	"branch_learning/condition"
	"branch_learning/strategy"
	"testing"
)

func TestBackTester_TestStream(t *testing.T) {
	redCondition := condition.CandleTypeCondition{CandleIndex: 0, IsGreen: false}
	greenCondition := condition.CandleTypeCondition{CandleIndex: 1, IsGreen: true}
	conditions := []condition.ICondition{redCondition, greenCondition}

	redCandleMap := make(map[string]float32)
	redCandleMap["close"] = 6
	redCandleMap["open"] = 10
	redCandleMap["high"] = 200

	greenCandleMap := make(map[string]float32)
	greenCandleMap["close"] = 10
	greenCandleMap["open"] = 6
	greenCandleMap["high"] = 11
	redCandle := candle.CreateCandle(redCandleMap)
	greenCandle := candle.CreateCandle(greenCandleMap)
	candles := []candle.Candle{redCandle, greenCandle, redCandle, greenCandle}
	stream := candle_stream.CreateCandleStream(candles)

	strategy := strategy.CreateStrategy(2, 100, 100, conditions)
	backtester := CreateBackTester(strategy)

	backtester.Test(stream)

	stats := backtester.Stats()
	exits := backtester.orderMananger.Exits()

	if stats.Losses() != 0 || stats.Wins() != 1 || stats.Losses()+stats.Wins() != 1 {
		t.Log("Error in the addition to the stats")
		t.Error("AssertionError")
	}
	if len(exits) != 0 {
		t.Log("Should have had zero order exit")
		t.Error("AssertionError")
	}
}
