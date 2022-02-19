package strategy

import (
	candle_stream "branch_learning/candle_stream"
	"branch_learning/condition"
	exit "branch_learning/exit"
)

type Strategy struct {
	takeProfit float32 //In Percentage
	stopLoss   float32 //In Percentage
	windowSize int
	conditions *condition.Conditions
}

func (strategy *Strategy) WindowSize() int {
	return strategy.windowSize
}

func (strategy *Strategy) TakeProfit() float32 {
	return strategy.takeProfit
}
func (strategy *Strategy) StopLoss() float32 {
	return strategy.stopLoss
}

func (strategy *Strategy) Conditions() *condition.Conditions {
	return strategy.conditions.Clone()
}

func CreateStrategy(windowSize int, takeProfit, stopLoss float32, conditions *condition.Conditions) *Strategy {
	return &Strategy{
		windowSize: windowSize,
		takeProfit: takeProfit,
		stopLoss:   stopLoss,
		conditions: conditions,
	}
}

func (strategy *Strategy) MeetsConditions(stream *candle_stream.CandleStream) bool {
	return strategy.conditions.MeetsConditions(stream)
}

func (strategy *Strategy) GetExit(price float32) exit.Exit {
	takeProfit := (1 + strategy.takeProfit/100) * price
	stopLoss := (1 - strategy.stopLoss/100) * price
	return exit.CreateExit(takeProfit, stopLoss)
}
