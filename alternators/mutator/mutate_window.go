package mutator

import (
	st "branch_learning/strategy"
	"math/rand"
	"time"
)

const (
	WindowSizeReducer = 5 // either increases or decreases by 10% (1 + random/reducer - (1/reducer/2))
)

func MutateWindowSize(strategy *st.Strategy) *st.Strategy {
	rand.Seed(time.Now().Unix())
	multiplier := 1.0 + (rand.Float32()/5.0 - 1.0/2.0/WindowSizeReducer)
	windowSize := int(multiplier * float32(strategy.WindowSize()))
	conditions := strategy.Conditions()

	for i := len(conditions) - 1; i >= 0; i-- {
		if !conditions[i].IsValidStreamSize(windowSize) {
			conditions = append(conditions[:i], conditions[i+1:]...)
		}
	}
	return st.CreateStrategy(windowSize, strategy.TakeProfit(), strategy.StopLoss(), conditions)
}
