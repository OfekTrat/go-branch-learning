package reproducer

import (
	"branch_learning/condition"
	st "branch_learning/strategy"
	"math/rand"
)

func Reproduce(s1 *st.Strategy, s2 *st.Strategy) *st.Strategy {
	nConds := 0
	if s2.Conditions().Length() > 0 {
		nConds = rand.Intn(s2.Conditions().Length())
	}
	return reproduceByNConditions(s1, s2, nConds)
}

func reproduceByNConditions(s1, s2 *st.Strategy, nConds int) *st.Strategy {
	conditions1 := s1.Conditions().ToList()
	conditions2 := s2.Conditions().ToList()

	for i := 0; i < nConds; i++ {
		if conditions2[i].IsValidStreamSize(s1.WindowSize()) {
			conditions1 = append(conditions1, conditions2[i])
		}
	}
	return st.CreateStrategy(s1.WindowSize(), s1.TakeProfit(), s1.StopLoss(), condition.CreateConditions(conditions1))
}
