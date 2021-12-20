package random

import (
	condition "branch_learning/condition"
	condition_list "branch_learning/condition_list"
	"math/rand"
)

func CreateRandomCandleTypeCondition(streamsize int) condition.ICondition {
	n := rand.Intn(streamsize)
	isGreen := rand.Intn(2)

	if isGreen == 1 {
		return condition_list.CandleTypeCondition{CandleIndex: n, IsGreen: true}
	}
	return condition_list.CandleTypeCondition{CandleIndex: n, IsGreen: false}
}
