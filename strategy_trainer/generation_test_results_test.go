package strategytrainer

import (
	st "branch_learning/strategy_tester"
	"math/rand"
	"testing"
)

func TestMergeSortTestersByScore(t *testing.T) {
	rand.Seed(1)
	testers := createTesters(20)
	scramble(testers, 20)

	orderedTesters := orderStrategyTestersByScore(testers)
	for i, tester := range orderedTesters {
		if i+1 != int(tester.Results().Score) || i+1 != tester.Strategy().Id() {
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

func TestGetBestStrategyTester(t *testing.T) {
	nStrategies := 20
	testers := createTesters(nStrategies)
	testResults := newGenerationTestResultsFromStrategyTesters(testers)
	chance := testResults.maxChance
	bestStrategyTester := testResults.tree.GetStrategyTesterByChance(chance)

	if bestStrategyTester.Strategy().Id() != 20 {
		t.Logf("Expected %d, Got %d", nStrategies, bestStrategyTester.Strategy().Id())
		t.Error("AssertionError - Build generation test results wrong.")
	}
}

func TestGetBestStrategies(t *testing.T) {
	nStrategies := 20
	testers := createTesters(nStrategies)
	testResults := newGenerationTestResultsFromStrategyTesters(testers)
	strategies := testResults.GetNBestStrategy(3)

	for _, strategy := range strategies {
		if strategy.Id() != nStrategies {
			t.Logf("Expected %d Got %d", nStrategies, strategy.Id())
			t.Error("AssertionError")
		}
		nStrategies -= 1
	}
}
