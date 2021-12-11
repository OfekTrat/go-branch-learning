package reproducer

import (
	st "branch_learning/strategy"
	"math/rand"
	"time"
)

func Reproduce(s1 *st.Strategy, s2 *st.Strategy) *st.Strategy {
	rand.Seed(time.Now().Unix())
	conditionsLength := len(s2.Conditions())
	nConditions := rand.Intn(conditionsLength / 2)
	indexes := []int{}

	for i := 0; i < nConditions; i++ {
		indexes = append(indexes, rand.Intn(conditionsLength))
	}
	return reproduceByIndexList(s1, s2, indexes)
}

func reproduceByIndexList(s1, s2 *st.Strategy, indexes []int) *st.Strategy {
	conditions1 := s1.Conditions()
	conditions2 := s2.Conditions()

	for _, index := range indexes {
		conditions1 = append(conditions1, conditions2[index])
	}
	return st.CreateStrategy(s1.WindowSize(), s1.TakeProfit(), s1.StopLoss(), conditions1)
}
