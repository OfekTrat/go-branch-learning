package mutator

import (
	"branch_learning/condition"
	condition_list "branch_learning/condition_list"
	st "branch_learning/strategy"
	"testing"
)

func TestChangeCondition(t *testing.T) {
	cond1 := condition_list.CandleComparisonCondition{CandleIndex1: 0, CandlePart1: "open", CandleIndex2: 0, CandlePart2: "close", Percentage: 0}
	cond2 := condition_list.CandleComparisonCondition{CandleIndex1: 0, CandlePart1: "open", CandleIndex2: 0, CandlePart2: "close", Percentage: 0}
	s := st.CreateStrategy(0, 0, 1000, 1, 1, condition.CreateConditions([]condition.ICondition{cond1, cond2}))
	s2 := mutateChangeConditionByIndex(1, 1, s, 0)

	if s.Conditions().GetByIndex(0).Equals(s2.Conditions().GetByIndex(0)) {
		t.Error("AssertionError")
	}
}
