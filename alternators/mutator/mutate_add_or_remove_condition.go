package mutator

import (
	st "branch_learning/strategy"
	utils "branch_learning/utils/random"
	"math/rand"
)

func MutateAddCondition(id, generation int, strategy *st.Strategy) *st.Strategy {
	c := utils.GetRandomCondition(strategy.WindowSize())
	conditions := strategy.Conditions()
	conditions.Add(c)
	return st.CreateStrategy(id, generation, strategy.WindowSize(), strategy.TakeProfit(), strategy.StopLoss(), conditions)
}

func MutateRemoveCondition(id, generation int, strategy *st.Strategy) *st.Strategy {
	conditions := strategy.Conditions()
	conditionsLength := conditions.Length()
	if conditionsLength == 0 {
		return st.CreateStrategy(id, generation, strategy.WindowSize(), strategy.TakeProfit(), strategy.StopLoss(), conditions)
	}
	randomConditionToRemove := rand.Intn(conditionsLength)
	conditions.RemoveByIndex(randomConditionToRemove)
	return st.CreateStrategy(id, generation, strategy.WindowSize(), strategy.TakeProfit(), strategy.StopLoss(), conditions)
}
