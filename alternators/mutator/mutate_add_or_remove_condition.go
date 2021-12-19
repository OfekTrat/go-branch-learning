package mutator

import (
	"branch_learning/condition"
	st "branch_learning/strategy"
	utils "branch_learning/utils/random"
	"math/rand"
)

func MutateAddCondition(strategy *st.Strategy) *st.Strategy {
	c := utils.GetRandomCondition(strategy.WindowSize())
	conditionList := strategy.Conditions().ToList()
	conditions := append(conditionList, c)
	return st.CreateStrategy(strategy.WindowSize(), strategy.TakeProfit(), strategy.StopLoss(), condition.CreateConditions(conditions))
}

func MutateRemoveCondition(strategy *st.Strategy) *st.Strategy {
	conditions := strategy.Conditions().ToList()
	if len(conditions) > 1 {
		randCondition := rand.Intn(len(conditions))
		conditions = append(conditions[:randCondition], conditions[randCondition+1:]...)
		return st.CreateStrategy(strategy.WindowSize(), strategy.TakeProfit(), strategy.StopLoss(), condition.CreateConditions(conditions))
	} else {
		return st.CreateStrategy(strategy.WindowSize(), strategy.TakeProfit(), strategy.StopLoss(), condition.CreateConditions([]condition.ICondition{}))
	}
}
