package output

import (
	bt "branch_learning/backtester"
	"fmt"
)

func PrintScore(backtester *bt.BackTester) {
	wins := backtester.Stats().Wins()
	losses := backtester.Stats().Losses()
	score := backtester.Score()
	takeProfit := backtester.Strategy().TakeProfit()
	stopLoss := backtester.Strategy().StopLoss()

	fmt.Printf("Score: %v\n", score)
	fmt.Printf("Wins: %v, Losses: %v\n", wins, losses)
	fmt.Printf("TakeProfit %v, StopLoss: %v\n", takeProfit, stopLoss)
}
