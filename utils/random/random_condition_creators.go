package random

import (
	condition "branch_learning/condition"
	condition_list "branch_learning/condition_list"
	"math/rand"
)

var (
	randConditionCreators = []func(int) condition.ICondition{
		CreateRandomCandleTypeCondition,
		CreateRandomCandleComparisonCondition,
	}
)

func CreateRandomCandleTypeCondition(streamsize int) condition.ICondition {
	n := rand.Intn(streamsize)
	isGreen := rand.Intn(2)

	if isGreen == 1 {
		return condition_list.CandleTypeCondition{CandleIndex: n, IsGreen: true}
	}
	return condition_list.CandleTypeCondition{CandleIndex: n, IsGreen: false}
}

func CreateRandomCandleComparisonCondition(streamsize int) condition.ICondition {
	candleParts := []string{"open", "high", "close", "low"}
	index1 := rand.Intn(streamsize)
	index2 := rand.Intn(streamsize)
	part1 := candleParts[rand.Intn(len(candleParts))]
	part2 := candleParts[rand.Intn(len(candleParts))]
	percentage := (rand.Float32() - 0.5) * 200
	return condition_list.CandleComparisonCondition{CandleIndex1: index1, CandlePart1: part1,
		CandleIndex2: index2, CandlePart2: part2, Percentage: percentage}
}
