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

func MutateStrategy(id, generation int, strategy *st.Strategy) *st.Strategy {
	randMutationType := rand.Intn(5)
	return mutateStrategy(id, generation, strategy, randMutationType)
}

func mutateStrategy(id, generation int, strategy *st.Strategy, mutationType int) *st.Strategy {
	switch mutationType {
	case MutateWindowChance:
		return MutateWindowSize(id, generation, strategy)
	case MutateTakeProfitChance:
		return MutateTakeProfit(id, generation, strategy)
	case MutateStopLossChance:
		return MutateStopLoss(id, generation, strategy)
	case MutateAddConditionChance:
		return MutateAddCondition(id, generation, strategy)
	case MutateChangeConditionChance:
		return MutateChangeCondition(id, generation, strategy)
	case MutateRemoveConditionChance:
		return MutateRemoveCondition(id, generation, strategy)
	default:
		return nil
	}
}
