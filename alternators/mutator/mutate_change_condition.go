package mutator

import (
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
	conditions := strategy.Conditions()
	mutatedCond := conditions.GetByIndex(randIndex).Mutate(strategy.WindowSize())
	conditions.SetInIndex(mutatedCond, randIndex)

	return st.CreateStrategy(strategy.WindowSize(), strategy.TakeProfit(), strategy.StopLoss(), conditions)
}
