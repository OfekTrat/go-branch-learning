package random

import (
	"branch_learning/condition"
	"math/rand"
)

var (
	randConditionCreators = []func(int) condition.ICondition{
		CreateRandomCandleTypeCondition,
	}
)

func GetRandomCondition(windowSize int) condition.ICondition {
	index := rand.Intn(len(randConditionCreators))
	return randConditionCreators[index](windowSize)
}
