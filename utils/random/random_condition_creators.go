package random

import (
	condition "branch_learning/condition"
	"math/rand"
)

func CreateRandomGreenCondition(streamsize int) condition.ICondition {
	return condition.GreenCondition{CandleIndex: rand.Intn(streamsize)}
}

func CreateRandomRedCondition(streamsize int) condition.ICondition {
	return condition.RedCondition{CandleIndex: rand.Intn(streamsize)}
}
