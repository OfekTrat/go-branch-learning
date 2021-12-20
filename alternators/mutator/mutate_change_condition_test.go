package mutator

import (
	"branch_learning/condition"
	condition_list "branch_learning/condition_list"
	st "branch_learning/strategy"
	"testing"
)

func TestChangeCondition(t *testing.T) {
	cond1 := condition_list.CandleTypeCondition{CandleIndex: 0, IsGreen: false}
	cond2 := condition_list.CandleTypeCondition{CandleIndex: 1, IsGreen: true}
	s := st.CreateStrategy(1000, 1, 1, condition.CreateConditions([]condition.ICondition{cond1, cond2}))
	s2 := mutateChangeConditionByIndex(s, 0)
	if s.Conditions().ToList()[0].Hash() == s2.Conditions().ToList()[0].Hash() {
		t.Error("AssertionError")
	}
}
