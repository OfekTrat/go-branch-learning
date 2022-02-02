package mutator

import (
	st "branch_learning/strategy"
	utils "branch_learning/utils/random"
	"math/rand"
)

func MutateAddCondition(strategy *st.Strategy) *st.Strategy {
	c := utils.GetRandomCondition(strategy.WindowSize())
	conditions := strategy.Conditions()
	conditions.Add(c)
	return st.CreateStrategy(strategy.WindowSize(), strategy.TakeProfit(), strategy.StopLoss(), conditions)
}

func MutateRemoveCondition(strategy *st.Strategy) *st.Strategy {
	conditions := strategy.Conditions()
	conditionsLength := conditions.Length()
	if conditionsLength == 0 {
		return st.CreateStrategy(strategy.WindowSize(), strategy.TakeProfit(), strategy.StopLoss(), conditions)
	}
	randomConditionToRemove := rand.Intn(conditionsLength)
	conditions.RemoveByIndex(randomConditionToRemove)
	return st.CreateStrategy(strategy.WindowSize(), strategy.TakeProfit(), strategy.StopLoss(), conditions)
}
