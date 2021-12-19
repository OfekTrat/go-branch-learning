package random

import (
	"branch_learning/condition"
	st "branch_learning/strategy"
	"math/rand"
)

type RandomStrategyConfig struct {
	WindowMin          int
	WindowMax          int
	ExitMin            float32
	ExitMax            float32
	ConditionNumberMin int
	ConditionNumberMax int
}

func CreateRandomStrategy(config *RandomStrategyConfig) *st.Strategy {
	windowSize := getRandomInt(config.WindowMin, config.WindowMax)
	takeProfit := getRandomFloat32(1.0, config.ExitMax)
	stopLoss := getRandomFloat32(config.ExitMin, 1.0)
	nConditions := getRandomInt(config.ConditionNumberMin, config.ConditionNumberMax)

	randConditions := getRandomConditions(nConditions, windowSize)
	return st.CreateStrategy(windowSize, takeProfit, stopLoss, randConditions)
}

func getRandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func getRandomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func getRandomConditions(nConditions, windowSize int) *condition.Conditions {
	conditions := make([]condition.ICondition, nConditions)

	for i := 0; i < nConditions; i++ {
		conditions[i] = GetRandomCondition(windowSize)
	}
	return condition.CreateConditions(conditions)
}
