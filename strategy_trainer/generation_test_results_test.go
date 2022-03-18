package strategytrainer

import (
	st "branch_learning/strategy_tester"
	"math/rand"
	"testing"
)

func TestMergeSortTestersByScore(t *testing.T) {
	rand.Seed(1)
	testers := createTesters(20)
	scramble(testers, 10)

	orderedTesters := orderStrategyTestersByScore(testers)
	for i, tester := range orderedTesters {
		if i+1 != int(tester.Results().Score) {
			t.Error("AssertionError")
		}
	}
}

func scramble(lst []*st.StrategyTester, n int) {
	lstLength := len(lst)

	for i := 0; i < n; i++ {
		rand1 := rand.Intn(lstLength)
		rand2 := rand.Intn(lstLength)
		lst[rand1], lst[rand2] = lst[rand2], lst[rand1]
	}
}
