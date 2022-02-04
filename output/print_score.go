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

	log.Printf("{\n")
	log.Printf("    score: %v,\n", score)
	log.Printf("    Wins: %v\n", wins)
	log.Printf("    losses: %v", losses)
	log.Printf("    stop_loss: %v\n", stopLoss)
	log.Printf("    take_profit: %v", takeProfit)
	log.Printf("    win_rate: %v\n", float32(wins)/float32((wins+losses))*100)
	log.Printf("    reward_risk: %v\n", takeProfit/stopLoss)
	log.Printf("}\n")
}
