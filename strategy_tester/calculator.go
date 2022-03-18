package strategytester

import (
	"math"
)

func Score(results *TestResults) float64 {
	orderCount := results.Wins + results.Losses
	if orderCount == 0 {
		return 0
	}
	orderCountWeight := calcOrderCountWeight(orderCount)
	winRate := float64(results.Wins) / float64(orderCount)
	lossRate := float64(results.Losses) / float64(orderCount)
	totalEstimatedEarningsForHundredOrders := (winRate * results.TakeProfit) - (lossRate * results.StopLoss)
	conditionCountWeight := calcConditionCountWeight(results.ConditionCount)

	return totalEstimatedEarningsForHundredOrders * orderCountWeight * conditionCountWeight
}

func calcOrderCountWeight(sumOrders int) float64 {
	/*
		kind of sigmoid function aims for 3% of number of orders
		   4
		1 + e^(-0.005*sumOrders)
		minus 2

		It should be changed to be relative to the given data size (if possible)
	*/

	return 4/(1+math.Pow(float64(math.E), -0.005*float64(sumOrders))) - 2
}

func calcConditionCountWeight(numberOfConditions int) float64 {
	/*
		This function gives wait to the number of conditions of a strategy.
		The reason for doing that is to keep the strategies simple and not get too much copmplex.
	*/
	threshold := 100.0
	slope := -0.02
	if float64(numberOfConditions) <= threshold {
		return float64(1)
	} else {
		return slope*float64(numberOfConditions) + (1 - (slope * threshold))
	}
}
