package mutator

import (
	st "branch_learning/strategy"
	"math/rand"
)

const (
	MutateWindowChance          = 0
	MutateTakeProfitChance      = 1
	MutateStopLossChance        = 2
	MutateAddConditionChance    = 3
	MutateChangeConditionChance = 4
	MutateRemoveConditionChance = 5
)

func MutateStrategy(strategy *st.Strategy) *st.Strategy {
	randMutationType := rand.Intn(5)
	return mutateStrategy(strategy, randMutationType)
}

func mutateStrategy(strategy *st.Strategy, mutationType int) *st.Strategy {
	switch mutationType {
	case MutateWindowChance:
		return MutateWindowSize(strategy)
	case MutateTakeProfitChance:
		return MutateTakeProfit(strategy)
	case MutateStopLossChance:
		return MutateStopLoss(strategy)
	case MutateAddConditionChance:
		return MutateAddCondition(strategy)
	case MutateChangeConditionChance:
		return MutateChangeCondition(strategy)
	case MutateRemoveConditionChance:
		return MutateRemoveCondition(strategy)
	default:
		return nil
	}
}
