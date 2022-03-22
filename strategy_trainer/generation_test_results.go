package strategytrainer

import (
	st "branch_learning/strategy"
	tester "branch_learning/strategy_tester"
)

type generationTestResults struct {
	maxChance float64
	tree      *chanceTree
}

func newGenerationTestResultsFromStrategyTesters(testers []*tester.StrategyTester) *generationTestResults {
	orderedStrategyTesters := orderStrategyTestersByScore(testers)
	chances := calcChancesFromOrderedStrategyTesters(testers)
	tree := createChanceTree(orderedStrategyTesters, chances)
	maxChance := chances[len(chances)-1]

	logger.Info.Println("Calculated Test Results and probabilities for choosing strategy")

	return &generationTestResults{
		maxChance: maxChance,
		tree:      tree,
	}
}

// Merge Sort
func orderStrategyTestersByScore(testers []*tester.StrategyTester) []*tester.StrategyTester {
	if len(testers) == 1 || len(testers) == 0 {
		return testers
	}
	testersLength := len(testers)
	middlePoint := testersLength / 2
	orderedStrategyTesters := make([]*tester.StrategyTester, testersLength)
	leftOrdered := orderStrategyTestersByScore(testers[:middlePoint])
	rightOrdered := orderStrategyTestersByScore(testers[middlePoint:])

	leftLength := len(leftOrdered)
	rightLength := len(rightOrdered)
	rightPos := 0
	leftPos := 0
	mainPos := 0

	for rightPos < rightLength && leftPos < leftLength {
		if leftOrdered[leftPos].Results().Score < rightOrdered[rightPos].Results().Score {
			orderedStrategyTesters[mainPos] = leftOrdered[leftPos]
			leftPos++
		} else {
			orderedStrategyTesters[mainPos] = rightOrdered[rightPos]
			rightPos++
		}
		mainPos++
	}

	if leftPos == leftLength {
		for rightPos < rightLength {
			orderedStrategyTesters[mainPos] = rightOrdered[rightPos]
			rightPos++
			mainPos++
		}
	} else {
		for leftPos < leftLength {
			orderedStrategyTesters[mainPos] = leftOrdered[leftPos]
			leftPos++
			mainPos++
		}
	}
	return orderedStrategyTesters
}

func (gtr *generationTestResults) GetStrategyByChance(chance float64) *st.Strategy {
	return gtr.tree.GetStrategyTesterByChance(chance).Strategy()
}

func (gtr *generationTestResults) GetMaxChance() float64 {
	return gtr.maxChance
}
