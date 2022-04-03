package strategytester

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
	if sumOrders >= configuration.SumOrdersThreshold() {
		return 1
	}
	return float64(sumOrders) / float64(configuration.SumOrdersThreshold())
}

func calcConditionCountWeight(numberOfConditions int) float64 {
	/*
		This function gives wait to the number of conditions of a strategy.
		The reason for doing that is to keep the strategies simple and not get too much copmplex.
	*/
	if numberOfConditions <= configuration.ConditionCountThreshold() {
		return float64(1)
	} else {
		return float64(1) + (float64(configuration.ConditionCountSlope()) * float64(configuration.ConditionCountThreshold()))
	}
}
