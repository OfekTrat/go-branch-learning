package random

import (
	condition "branch_learning/condition"
	"math/rand"
)

func CreateRandomCandleTypeCondition(streamsize int) condition.ICondition {
	n := rand.Intn(streamsize)
	isGreen := rand.Intn(2)

	if isGreen == 1 {
		return condition.CandleTypeCondition{CandleIndex: n, IsGreen: true}
	}
	return condition.CandleTypeCondition{CandleIndex: n, IsGreen: false}
}
