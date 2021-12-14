package random

import (
	"branch_learning/condition"
	"math/rand"
)

var (
	randConditionCreators = []func(int) condition.ICondition{
		CreateRandomGreenCondition,
		CreateRandomRedCondition,
	}
)

func GetRandomCondition(windowSize int) condition.ICondition {
	index := rand.Intn(windowSize)
	return randConditionCreators[index](windowSize)
}
