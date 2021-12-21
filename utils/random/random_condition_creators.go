package random

import (
	condition "branch_learning/condition"
	condition_list "branch_learning/condition_list"
	"math/rand"
)

var (
	candleParts           = []string{"open", "high", "close", "low"}
	randConditionCreators = []func(int) condition.ICondition{
		CreateRandomCandleComparisonCondition,
		CreateRandomRSICondition,
		CreateRandomRSICompareCondition,
		CreateRandomMACDCondition,
		CreateRandomMACDCompareCondition,
	}
)

func CreateRandomCandleComparisonCondition(streamsize int) condition.ICondition {
	index1 := rand.Intn(streamsize)
	index2 := rand.Intn(streamsize)
	part1 := candleParts[rand.Intn(len(candleParts))]
	part2 := candleParts[rand.Intn(len(candleParts))]
	percentage := (rand.Float32() - 0.5) * 200
	return condition_list.CandleComparisonCondition{CandleIndex1: index1, CandlePart1: part1,
		CandleIndex2: index2, CandlePart2: part2, Percentage: percentage}
}

func CreateRandomRSICondition(streamsize int) condition.ICondition {
	randIndex := rand.Intn(streamsize)
	isGreaterThan := rand.Intn(2)
	randRSIValue := rand.Float32() * 100
	if isGreaterThan == 1 {
		return condition_list.RSICondition{CandleIndex: randIndex, GreaterThan: true, RsiValue: randRSIValue}
	} else {
		return condition_list.RSICondition{CandleIndex: randIndex, GreaterThan: false, RsiValue: randRSIValue}
	}
}

func CreateRandomRSICompareCondition(streamsize int) condition.ICondition {
	randIndex1 := rand.Intn(streamsize)
	randIndex2 := rand.Intn(streamsize)
	percentage := (rand.Float32() - 0.5) * 200
	return condition_list.RSICompareCondition{CandleIndex1: randIndex1, CandleIndex2: randIndex2, Percentage: percentage}
}

func CreateRandomMACDCondition(streamsize int) condition.ICondition {
	randIndex := rand.Intn(streamsize)
	randMACDValue := (rand.Float32() - 0.5) * 10
	greaterThan := rand.Intn(2)
	if greaterThan == 1 {
		return condition_list.MACDCondition{CandleIndex: randIndex, MacdValue: randMACDValue, GreaterThan: true}
	} else {
		return condition_list.MACDCondition{CandleIndex: randIndex, MacdValue: randMACDValue, GreaterThan: false}
	}
}

func CreateRandomMACDCompareCondition(streamsize int) condition.ICondition {
	randIndex1 := rand.Intn(streamsize)
	randIndex2 := rand.Intn(streamsize)
	percentage := (rand.Float32() - 0.5) * 20
	return condition_list.MACDCompareCondition{CandleIndex1: randIndex1, CandleIndex2: randIndex2, Percentage: percentage}
}
