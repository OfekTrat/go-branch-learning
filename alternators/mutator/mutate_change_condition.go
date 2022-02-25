package mutator

import (
	st "branch_learning/strategy"
	"math/rand"
)

func MutateChangeCondition(id, generation int, strategy *st.Strategy) *st.Strategy {
	if strategy.Conditions().Length() > 0 {
		randomConditionIndex := rand.Intn(strategy.Conditions().Length())
		return mutateChangeConditionByIndex(id, generation, strategy, randomConditionIndex)
	}
	return strategy
}

func mutateChangeConditionByIndex(id, generation int, strategy *st.Strategy, randIndex int) *st.Strategy {
	conditions := strategy.Conditions()
	mutatedCond := conditions.GetByIndex(randIndex).Mutate(strategy.WindowSize())
	conditions.SetInIndex(mutatedCond, randIndex)

	return st.CreateStrategy(id, generation, strategy.WindowSize(), strategy.TakeProfit(), strategy.StopLoss(), conditions)
}
