package strategytrainer

import (
	cs "branch_learning/condition"
	s "branch_learning/strategy"
	st "branch_learning/strategy_tester"
	"testing"
)

func TestCreateChanceTree(t *testing.T) {
	testersNumber := 20
	testers := createTesters(testersNumber)
	chanceTree := createChanceTreeFromTesters(testers)
	chances := make([]float64, testersNumber)
	
	createListFromTree(chanceTree, chances, 0)
	t.Log(chances)
	
	sum := 0
	for i, v := range chances {
		sum += (i+1)
		if float64(sum) != v {
			t.Error("AssertionError")
		}
	}
}

func createTesters(testersNumber int) []*st.StrategyTester {
	conditions := cs.EmptyConditions()
	testers := make([]*st.StrategyTester, testersNumber)

	for i := 1; i < testersNumber+1; i++ {
		tmpStrategy := s.CreateStrategy(0, 0, 10, 1.1, 1.1, conditions)
		tester := st.NewStrategyTester(tmpStrategy)
		tester.Results().Score = float64(i)
		testers[i-1] = tester
	}
	return testers
}

func createListFromTree(tree *chanceTree, chances []float64, currentIndex int) int {
	if tree == nil {
		return currentIndex
	}
	currentIndex = createListFromTree(tree.lowerChance, chances, currentIndex)
	chances[currentIndex] = tree.chance
	currentIndex = createListFromTree(tree.higherChance, chances, currentIndex + 1)
	return currentIndex
}
