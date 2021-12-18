package mutator

import (
	"branch_learning/condition"
	"branch_learning/strategy"
	"testing"
)

func TestMutateAddCondition(t *testing.T) {
	redCondition := condition.CandleTypeCondition{CandleIndex: 2, IsGreen: false}
	s := strategy.CreateStrategy(10, 1, 1, []condition.ICondition{redCondition})
	newS := MutateAddCondition(s)

	if len(newS.Conditions()) != len(s.Conditions())+1 {
		t.Error("AssertionError")
	}
}

func TestMutateRemoveCondition(t *testing.T) {
	gCond := condition.CandleTypeCondition{CandleIndex: 4, IsGreen: true}
	s := strategy.CreateStrategy(10, 1, 1, []condition.ICondition{gCond})
	newS := MutateRemoveCondition(s)

	if len(newS.Conditions()) != len(s.Conditions())-1 {
		t.Error("AssertionError")
	}
}

func TestMutateRemove_ZeroConditions(t *testing.T) {
	s := strategy.CreateStrategy(10, 1, 1, []condition.ICondition{})
	s2 := MutateRemoveCondition(s)
	if len(s2.Conditions()) != 0 {
		t.Error("AssertionError")
	}
}
