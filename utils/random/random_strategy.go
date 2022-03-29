package random

import (
	"branch_learning/condition"
	c "branch_learning/configuration"
	st "branch_learning/strategy"
	"math/rand"
)

var configuration *c.Configuration = c.GetConfiguration()

func CreateRandomStrategy(id, generation int) *st.Strategy {
	windowSize := getRandomInt(configuration.WindowMin(), configuration.WindowMax())
	takeProfit := getRandomFloat64(1.0, configuration.TakeProfitMax())
	stopLoss := getRandomFloat64(configuration.StopLossMin(), 1.0)
	nConditions := getRandomInt(configuration.ConditionNumberMin(), configuration.ConditionNumberMax())

	randConditions := getRandomConditions(nConditions, windowSize)
	return st.CreateStrategy(id, generation, windowSize, takeProfit, stopLoss, randConditions)
}

func getRandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func getRandomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func getRandomConditions(nConditions, windowSize int) *condition.Conditions {
	conditions := make([]condition.ICondition, nConditions)

	for i := 0; i < nConditions; i++ {
		conditions[i] = GetRandomCondition(windowSize)
	}
	return condition.CreateConditions(conditions)
}

// TODO: Make random function for each value and not for each type (for stop loss not for float32)
