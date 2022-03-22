package random

import (
	condition "branch_learning/condition"
	condition_list "branch_learning/condition_list"
	"math/rand"
)

var (
	randConditionCreators = []func(int) condition.ICondition{
		CreateRandomCandleComparisonCondition,
		CreateRandomPivotPointCondition,
		CreateRandomIndicatorCondition,
		CreateRandomIndicatorComparisonCondition,
	}
	possibleIndicatorsForIndicatorCondition        = []string{"macd", "rsi", "ao", "apo", "bias", "bop", "entropy", "kurtosis", "zscore", "adx", "amat_lr", "amat_sr", "aroon_osc", "chop", "decreasing", "increasing"}
	possibleIndicatorsForIndicatorCompareCondition = []string{"macd", "rsi", "volume", "ao", "apo", "bias", "bop", "entropy", "kurtosis", "zscore", "adx", "amat_lr", "amat_sr", "aroon_osc", "chop", "decreasing", "increasing"}
	greaterThanOptions                             = []bool{true, false}
)

func CreateRandomCandleComparisonCondition(streamsize int) condition.ICondition {
	index1 := rand.Intn(streamsize)
	index2 := rand.Intn(streamsize)
	part1 := condition_list.CandleParts[rand.Intn(len(condition_list.CandleParts))]
	part2 := condition_list.CandleParts[rand.Intn(len(condition_list.CandleParts))]
	percentage := (rand.Float64() - 0.5) * 200
	return condition_list.CandleComparisonCondition{CandleIndex1: index1, CandlePart1: part1,
		CandleIndex2: index2, CandlePart2: part2, Percentage: percentage}
}

func CreateRandomPivotPointCondition(streamsize int) condition.ICondition {
	randIndex := rand.Intn(streamsize)
	isGreaterThan := rand.Intn(2)
	percentage := getRandomPercentage()
	candlePart := condition_list.CandleParts[rand.Intn(len(condition_list.CandleParts))]
	pivotPart := condition_list.PivotParts[rand.Intn(len(condition_list.PivotParts))]
	if isGreaterThan == 1 {
		return condition_list.PivotPointCondition{CandleIndex: randIndex, CandlePart: candlePart, PivotPart: pivotPart, Percentage: percentage, GreaterThan: true}
	} else {
		return condition_list.PivotPointCondition{CandleIndex: randIndex, CandlePart: candlePart, PivotPart: pivotPart, Percentage: percentage, GreaterThan: false}
	}
}

func CreateRandomIndicatorCondition(streamsize int) condition.ICondition {
	randIndicator := possibleIndicatorsForIndicatorCondition[rand.Intn(len(possibleIndicatorsForIndicatorCondition))]
	randIndex := rand.Intn(streamsize)
	isGreaterThan := getRandomGreaterThan()
	percentage := getRandomPercentage()
	randIndicatorValue := getRandomIndicatorValueByIndicator(randIndicator)

	return condition_list.IndicatorCondition{
		Indicator:      randIndicator,
		CandleIndex:    randIndex,
		IndicatorValue: randIndicatorValue,
		Percentage:     percentage,
		GreaterThan:    isGreaterThan,
	}
}

func CreateRandomIndicatorComparisonCondition(streamsize int) condition.ICondition {
	randIndicator := possibleIndicatorsForIndicatorCompareCondition[rand.Intn(len(possibleIndicatorsForIndicatorCompareCondition))]
	randIndex1 := rand.Intn(streamsize)
	randIndex2 := rand.Intn(streamsize)
	percentage := getRandomPercentage()

	return condition_list.IndicatorCompareCondition{Indicator: randIndicator, CandleIndex1: randIndex1, CandleIndex2: randIndex2, Percentage: percentage}

}

func getRandomIndicatorValueByIndicator(indicator string) float64 {
	switch indicator {
	case "macd":
		return (rand.Float64() - 0.5) * 10
	case "rsi":
		return rand.Float64() * 100
	case "ao":
		return (rand.Float64() - 0.5) * 4000
	case "apo":
		return (rand.Float64() - 0.5) * 4000
	case "bias":
		return (rand.Float64() - 0.5) * 0.5
	case "bop":
		return (rand.Float64() - 0.5) * 2
	case "entropy":
		return rand.Float64() + 3
	case "kurtosis":
		return rand.Float64()*22 - 2
	case "zscore":
		return (rand.Float64() - 0.5) * 10
	case "adx":
		return rand.Float64() * 100
	case "amat_lr":
		return rand.Float64()
	case "amat_sr":
		return rand.Float64()
	case "aroon_osc":
		return (rand.Float64() - 0.5) * 200
	case "chop":
		return rand.Float64() * 100
	case "decreasing":
		return rand.Float64()
	case "increasing":
		return rand.Float64()
	default:
		return 0
	}
}

func getRandomGreaterThan() bool {
	return greaterThanOptions[rand.Intn(2)]
}

func getRandomPercentage() float64 {
	return (rand.Float64() - 0.5) * 20000
}
