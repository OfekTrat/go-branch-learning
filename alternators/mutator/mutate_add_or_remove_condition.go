package mutator

import (
	st "branch_learning/strategy"
	utils "branch_learning/utils/random"
	"math/rand"
	"time"
)

func MutateAddCondition(strategy *st.Strategy) *st.Strategy {
	c := utils.GetRandomCondition(strategy.WindowSize())
	conditions := append(strategy.Conditions(), c)
	return st.CreateStrategy(strategy.WindowSize(), strategy.TakeProfit(), strategy.StopLoss(), conditions)
}

func MutateRemoveCondition(strategy *st.Strategy) *st.Strategy {
	rand.Seed(time.Now().Unix())
	conditions := strategy.Conditions()
	if len(conditions) > 0 {
		randCondition := rand.Intn(len(conditions))
		conditions = append(conditions[:randCondition], conditions[randCondition+1:]...)
		return st.CreateStrategy(strategy.WindowSize(), strategy.TakeProfit(), strategy.StopLoss(), conditions)
	}
	return strategy
}
