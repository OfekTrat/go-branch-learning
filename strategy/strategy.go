package strategy

import (
	candle_stream "branch_learning/candle_stream"
	"branch_learning/condition"
)

type Strategy struct {
	takeProfit float32 //In Percentage
	stopLoss   float32 //In Percentage
	conditions []condition.ICondition
}

func CreateStrategy(takeProfit, stopLoss float32, conditions []condition.ICondition) Strategy {
	return Strategy{takeProfit: takeProfit, stopLoss: stopLoss, conditions: conditions}
}

func (strategy *Strategy) MeetsConditions(stream candle_stream.CandleStream) bool {
	for _, condition := range strategy.conditions {
		if !condition.MeetsCondition(&stream) {
			return false
		}
	}
	return true
}

func (strategy *Strategy) GetExit(price float32) Exit {
	takeProfit := (1 + strategy.takeProfit/100) * price
	stopLoss := (1 - strategy.stopLoss/100) * price
	return Exit{takeProfitPrice: takeProfit, stopLossPrice: stopLoss}
}
