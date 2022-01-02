package output

import (
	bt "branch_learning/backtester"
	"log"
)

func PrintScore(backtester *bt.BackTester) {
	wins := backtester.Stats().Wins()
	losses := backtester.Stats().Losses()
	score := backtester.Score()
	takeProfit := backtester.Strategy().TakeProfit()
	stopLoss := backtester.Strategy().StopLoss()

	log.Printf("Score: %v\n", score)
	log.Printf("Wins: %v, Losses: %v\n", wins, losses)
	log.Printf("TakeProfit|StopLoss: %v|%v\n", takeProfit, stopLoss)
	log.Printf("Win Rate: %v\n", float32(wins)/float32((wins+losses))*100)
	log.Printf("Reward/Risk: %v\n", takeProfit/stopLoss)
}
