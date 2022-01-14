package parser

import "testing"

func TestStrategyParser(t *testing.T) {
	strategyString := `{
		"stop_loss": 0.768,
		"take_profit": 1.2455,
		"window": 12,
		"conditions": [
			{
				"type": "CandleComparison",
				"CandleIndex1": 2,
				"CandlePart1": "close",
				"CandleIndex2": 8,
				"CandlePart2": "high",
				"Percentage": 28.325
			},
			{
				"type": "IndicatorCondition",
				"Indicator": "rsi",
				"CandleIndex": 7,
				"IndicatorValue": -1.5,
				"Percentage": 50,
				"GreaterThan": false
			}	
		]
	}`

	strategy := ParseStrategy([]byte(strategyString))
	conditions := strategy.Conditions().ToList()

	if conditions[0].ConditionType() != "CandleComparison" && conditions[1].ConditionType() != "IndicatorCondition" {
		t.Log(conditions[0].ConditionType())
		t.Log(conditions[1].ConditionType())
		t.Error()
	}
}
