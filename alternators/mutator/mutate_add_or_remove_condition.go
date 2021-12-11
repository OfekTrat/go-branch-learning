package mutator

import (
	condition "branch_learning/condition"
	st "branch_learning/strategy"
	"math/rand"
	"time"
)

var (
	randConditionCreators = []func(int) condition.ICondition{
		condition.CreateRandomGreenCondition,
		condition.CreateRandomRedCondition,
	}
)

func MutateAddCondition(strategy *st.Strategy) *st.Strategy {
	rand.Seed(time.Now().Unix())
	randCreator := rand.Intn(len(randConditionCreators))
	c := randConditionCreators[randCreator](strategy.WindowSize())
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
