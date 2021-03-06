package strategytrainer

import (
	tester "branch_learning/strategy_tester"
)

type chanceTree struct {
	chance         float64
	strategyTester *tester.StrategyTester
	lowerChance    *chanceTree
	higherChance   *chanceTree
}

func createChanceTree(orderedTestersByScore []*tester.StrategyTester, chances []float64) *chanceTree {
	testersLength := len(orderedTestersByScore)

	if testersLength == 0 {
		return nil
	}
	middleIndex := testersLength / 2
	chance := chances[middleIndex]
	higherChance := createChanceTree(orderedTestersByScore[middleIndex+1:], chances[middleIndex+1:])
	lowerChance := createChanceTree(orderedTestersByScore[:middleIndex], chances[:middleIndex])
	strategyTester := orderedTestersByScore[middleIndex]

	return &chanceTree{
		chance:         chance,
		strategyTester: strategyTester,
		lowerChance:    lowerChance,
		higherChance:   higherChance,
	}
}

func createChanceTreeFromTesters(orderedTestersByScore []*tester.StrategyTester) (float64, *chanceTree) {
	chances := calcChancesFromOrderedStrategyTesters(orderedTestersByScore)
	return chances[len(chances)-1], createChanceTree(orderedTestersByScore, chances)
}

func calcChancesFromOrderedStrategyTesters(orderedTesters []*tester.StrategyTester) []float64 {
	var sumChances float64 = 0
	chances := make([]float64, len(orderedTesters))

	for i, t := range orderedTesters {
		sumChances += t.Results().Score
		chances[i] = sumChances
	}
	return chances
}

func (ct *chanceTree) GetStrategyTesterByChance(chance float64) *tester.StrategyTester {
	if ct.chance == chance || (ct.higherChance == nil && ct.lowerChance == nil) {
		return ct.strategyTester
	} else if ct.chance < chance {
		if ct.higherChance == nil {
			return ct.strategyTester
		} else {
			return ct.higherChance.GetStrategyTesterByChance(chance)
		}
	} else {
		if ct.lowerChance == nil {
			return ct.strategyTester
		} else {
			return ct.lowerChance.GetStrategyTesterByChance(chance)
		}
	}
}
