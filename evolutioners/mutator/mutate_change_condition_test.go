package mutator

import (
	"branch_learning/condition"
	st "branch_learning/strategy"
	"testing"
)

func TestChangeCondition(t *testing.T) {
	cond1 := condition.RedCondition{CandleIndex: 0}
	cond2 := condition.GreenCondition{CandleIndex: 1}
	s := st.CreateStrategy(1000, 1, 1, []condition.ICondition{cond1, cond2})
	s2 := mutateChangeConditionByIndex(s, 0)
	if s.Conditions()[0] == s2.Conditions()[0] {
		t.Error("AssertionError")
	}
}
