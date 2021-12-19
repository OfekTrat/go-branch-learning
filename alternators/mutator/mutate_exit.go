package mutator

import (
	st "branch_learning/strategy"
	"math/rand"
	"time"
)

const (
	ExitReducer = 5
)

func MutateStopLoss(strategy *st.Strategy) *st.Strategy {
	rand.Seed(time.Now().Unix())
	multiplier := 1.0 + (rand.Float32()/5.0 - 1.0/2.0/ExitReducer)
	stopLoss := strategy.StopLoss() * multiplier

	if stopLoss > strategy.TakeProfit() {
		stopLoss = strategy.TakeProfit() * 0.99
	}

	return st.CreateStrategy(strategy.WindowSize(), strategy.TakeProfit(), stopLoss, strategy.Conditions())
}

func MutateTakeProfit(strategy *st.Strategy) *st.Strategy {
	multiplier := 1.0 + (rand.Float32()/5.0 - 1.0/2.0/ExitReducer)
	takeProfit := strategy.TakeProfit() * multiplier

	if takeProfit < strategy.StopLoss() {
		takeProfit = strategy.StopLoss() * 1.01
	}

	return st.CreateStrategy(strategy.WindowSize(), takeProfit, strategy.StopLoss(), strategy.Conditions())
}
