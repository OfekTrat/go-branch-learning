package strategy

import (
	candle_stream "branch_learning/candle_stream"
	"branch_learning/condition"
)

type Strategy struct {
	id         int
	generation int
	takeProfit float64 //In Percentage
	stopLoss   float64 //In Percentage
	windowSize int
	conditions *condition.Conditions
}

func (strategy *Strategy) WindowSize() int {
	return strategy.windowSize
}

func (strategy *Strategy) TakeProfit() float64 {
	return strategy.takeProfit
}
func (strategy *Strategy) StopLoss() float64 {
	return strategy.stopLoss
}

func (strategy *Strategy) Conditions() *condition.Conditions {
	return strategy.conditions.Clone()
}

func (strategy *Strategy) Id() int {
	return strategy.id
}

func (strategy *Strategy) Generation() int {
	return strategy.generation
}

func CreateStrategy(id, generation, windowSize int, takeProfit, stopLoss float64, conditions *condition.Conditions) *Strategy {
	return &Strategy{
		id:         id,
		generation: generation,
		windowSize: windowSize,
		takeProfit: takeProfit,
		stopLoss:   stopLoss,
		conditions: conditions,
	}
}

func CreateStrategyFromOtherStrategy(id, geneartion int, strategy *Strategy) *Strategy {
	return &Strategy{
		id:         id,
		generation: geneartion,
		takeProfit: strategy.takeProfit,
		stopLoss:   strategy.stopLoss,
		windowSize: strategy.windowSize,
		conditions: strategy.conditions.Clone(),
	}
}

func (strategy *Strategy) MeetsConditions(stream *candle_stream.CandleStream) bool {
	return strategy.conditions.MeetsConditions(stream)
}
