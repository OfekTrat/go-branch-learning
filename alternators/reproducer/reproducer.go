package reproducer

import (
	st "branch_learning/strategy"
	"math/rand"
)

func Reproduce(id, generation int, s1 *st.Strategy, s2 *st.Strategy) *st.Strategy {
	nConds := 0
	if s2.Conditions().Length() > 0 {
		nConds = rand.Intn(s2.Conditions().Length())
	}
	return reproduceByNConditions(id, generation, s1, s2, nConds)
}

func reproduceByNConditions(id, generation int, s1, s2 *st.Strategy, nConds int) *st.Strategy {
	conditions1 := s1.Conditions()
	conditions2 := s2.Conditions()

	for i := 0; i < nConds; i++ {
		if conditions2.GetByIndex(i).IsValidStreamSize(s1.WindowSize()) {
			conditions1.Add(conditions2.GetByIndex(i))
		}
	}
	return st.CreateStrategy(id, generation, s1.WindowSize(), s1.TakeProfit(), s1.StopLoss(), conditions1)
}
