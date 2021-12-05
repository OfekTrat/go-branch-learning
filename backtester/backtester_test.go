package backtester

import (
	"branch_learning/candle"
	candle_stream "branch_learning/candle_stream"
	"branch_learning/condition"
	condition_list "branch_learning/condition/condition_list"
	"branch_learning/strategy"
	"testing"
)

func TestBackTester_TestStream(t *testing.T) {
	redCondition := condition_list.RedCondition{CandleIndex: 0}
	greenCondition := condition_list.GreenCondition{CandleIndex: 1}
	conditions := []condition.ICondition{redCondition, greenCondition}

	redCandle := candle.CreateCandle(10, 200, 6, 7)
	greenCandle := candle.CreateCandle(7, 11, 6, 10)
	candles := []candle.Candle{redCandle, greenCandle, redCandle, greenCandle}
	stream := candle_stream.CreateCandleStream(candles)

	strategy := strategy.CreateStrategy(2, 100, 100, conditions)
	backtester := CreateBackTester(strategy)

	backtester.Test(stream)

	stats := backtester.Stats()
	exits := backtester.orderMananger.Exits()

	if stats.Losses() != 0 || stats.Wins() != 1 || stats.Matches() != 1 {
		t.Log("Error in the addition to the stats")
		t.Error("AssertionError")
	}
	if len(exits) != 1 {
		t.Log("Should have had one order exit")
		t.Error("AssertionError")
	}
}
