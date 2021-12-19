package mutator

import (
	"branch_learning/condition"
	st "branch_learning/strategy"
	"math/rand"
)

func MutateChangeCondition(strategy *st.Strategy) *st.Strategy {
	if strategy.Conditions().Length() > 0 {
		randomConditionIndex := rand.Intn(strategy.Conditions().Length())
		return mutateChangeConditionByIndex(strategy, randomConditionIndex)
	}
	return strategy
}

func mutateChangeConditionByIndex(strategy *st.Strategy, randIndex int) *st.Strategy {
	conditions := strategy.Conditions().ToList()
	conditions[randIndex] = conditions[randIndex].Mutate(strategy.WindowSize())
	return st.CreateStrategy(strategy.WindowSize(), strategy.TakeProfit(), strategy.StopLoss(), condition.CreateConditions(conditions))
}
