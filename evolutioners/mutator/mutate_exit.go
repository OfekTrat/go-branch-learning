package mutator

import (
	st "branch_learning/strategy"
	"math/rand"
)

const (
	ExitReducer = 5
)

func mutateStopLoss(strategy *st.Strategy) *st.Strategy {
	multiplier := 0.5 + (rand.Float32() / ExitReducer) + 0.5*ExitReducer
	stopLoss := strategy.StopLoss() * multiplier
	return st.CreateStrategy(strategy.WindowSize(), strategy.TakeProfit(), stopLoss, strategy.Conditions())
}

func mutateTakeProfit(strategy *st.Strategy) *st.Strategy {
	multiplier := 0.5 + (rand.Float32() / ExitReducer) + 0.5*ExitReducer
	takeProfit := strategy.TakeProfit() * multiplier
	return st.CreateStrategy(strategy.WindowSize(), takeProfit, strategy.StopLoss(), strategy.Conditions())
}
