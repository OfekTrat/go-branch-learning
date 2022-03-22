package mutator

import (
	st "branch_learning/strategy"
	"math/rand"
	"time"
)

const (
	ExitReducer = 5
)

func MutateStopLoss(id, generation int, strategy *st.Strategy) *st.Strategy {
	rand.Seed(time.Now().Unix())
	multiplier := 1.0 + (rand.Float64()/5.0 - 1.0/2.0/ExitReducer)
	stopLoss := strategy.StopLoss() * multiplier
	return st.CreateStrategy(id, generation, strategy.WindowSize(), strategy.TakeProfit(), stopLoss, strategy.Conditions())
}

func MutateTakeProfit(id, generation int, strategy *st.Strategy) *st.Strategy {
	multiplier := 1.0 + (rand.Float64()/5.0 - 1.0/2.0/ExitReducer)
	takeProfit := strategy.TakeProfit() * multiplier
	return st.CreateStrategy(id, generation, strategy.WindowSize(), takeProfit, strategy.StopLoss(), strategy.Conditions())
}
