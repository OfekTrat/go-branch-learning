package mutator

import (
	st "branch_learning/strategy"
	"math/rand"
	"time"
)

func MutateChangeCondition(strategy *st.Strategy) *st.Strategy {
	rand.Seed(time.Now().Unix())
	randomConditionIndex := rand.Intn(len(strategy.Conditions()))
	return mutateChangeConditionByIndex(strategy, randomConditionIndex)
}

func mutateChangeConditionByIndex(strategy *st.Strategy, randIndex int) *st.Strategy {
	rand.Seed(time.Now().Unix())
	conditions := strategy.Conditions()
	conditions[randIndex] = conditions[randIndex].Mutate(strategy.WindowSize())
	return st.CreateStrategy(strategy.WindowSize(), strategy.TakeProfit(), strategy.StopLoss(), conditions)
}
