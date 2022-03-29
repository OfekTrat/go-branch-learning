package strategytester

const (
	SUM_ORDERS_THRESHOLD      = 100
	CONDITION_COUNT_THRESHOLD = 100
	CONDITION_SLOPE           = -0.02
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
	// I want to have more than a 100 orders.
	// It does not matter how many orders, but the win rate.
	if sumOrders >= SUM_ORDERS_THRESHOLD {
		return 1
	}
	return float64(sumOrders) * 1 / SUM_ORDERS_THRESHOLD
}

func calcConditionCountWeight(numberOfConditions int) float64 {
	/*
		This function gives wait to the number of conditions of a strategy.
		The reason for doing that is to keep the strategies simple and not get too much copmplex.
	*/
	if float64(numberOfConditions) <= CONDITION_COUNT_THRESHOLD {
		return float64(1)
	} else {
		return CONDITION_SLOPE*float64(numberOfConditions) + (1 - (CONDITION_SLOPE * SUM_ORDERS_THRESHOLD))
	}
}
