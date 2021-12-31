package tester

import "testing"

func TestStrategyLoader(t *testing.T) {
	strategyString := `
	StopLoss: 0.768
	TakeProfit: 1.2455
	Window: 12
	{
		"type": "CandleComparison",
		"CandleIndex1": 2,
		"CandlePart1": "close",
		"CandleIndex2": 8,
		"CandlePart2": "high",
		"Percentage": 28.325
	}
	{
		"type": "MACD",
		"CandleIndex": 7,
		"MacdValue": -1.5,
		"GreaterThan": false
	}
	`
	strategy := parseStrategy(strategyString)
	if strategy == nil {
		t.Error("AssertionError - Strategy is nil")
	}

	if strategy.StopLoss() != 0.768 || strategy.TakeProfit() != 1.2455 || strategy.WindowSize() != 12 {
		t.Error("AssertionError - Did not parse stoploss|takeprofit|windowsize correctly")
	}
	conditions := strategy.Conditions().ToList()
	if len(conditions) != 2 {
		t.Error("AssertionError - Did not parse conditions correctly")
	}
	if conditions[0].ConditionType() != "MACD" && conditions[1].ConditionType() != "MACD" {
		t.Error("AssertionError - Did not parse by condition type correctly")
	}
}
