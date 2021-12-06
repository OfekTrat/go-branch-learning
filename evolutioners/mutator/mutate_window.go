package mutator

import (
	st "branch_learning/strategy"
	"math/rand"
)

const (
	WindowSizeReducer = 5 // either increases or decreases by 10% (1 + random/reducer - (1/reducer/2))
)

func mutateWindowSize(strategy *st.Strategy) *st.Strategy {
	multiplier := 0.5 + (rand.Float32() / WindowSizeReducer) + (0.5 * WindowSizeReducer)
	windowSize := int(multiplier * float32(strategy.WindowSize()))
	conditions := strategy.Conditions()

	for i := len(conditions) - 1; i >= 0; i-- {
		c := conditions[i]
		if !c.IsValidStreamSize(windowSize) {
			conditions = append(conditions[:i], conditions[i+1:]...)
		}
	}
	return st.CreateStrategy(windowSize, strategy.TakeProfit(), strategy.StopLoss(), conditions)
}
