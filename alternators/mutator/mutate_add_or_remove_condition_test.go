package mutator

import (
	"branch_learning/condition"
	"branch_learning/strategy"
	"testing"
)

func TestMutateAddCondition(t *testing.T) {
	redCondition := condition.CandleTypeCondition{CandleIndex: 2, IsGreen: false}
	s := strategy.CreateStrategy(10, 1, 1, condition.CreateConditions([]condition.ICondition{redCondition}))
	newS := MutateAddCondition(s)

	if newS.Conditions().Length() != s.Conditions().Length()+1 {
		t.Error("AssertionError")
	}
}

func TestMutateRemoveCondition(t *testing.T) {
	gCond := condition.CandleTypeCondition{CandleIndex: 4, IsGreen: true}
	s := strategy.CreateStrategy(10, 1, 1, condition.CreateConditions([]condition.ICondition{gCond}))
	newS := MutateRemoveCondition(s)

	if newS.Conditions().Length() != s.Conditions().Length()-1 {
		t.Error("AssertionError")
	}
}

func TestMutateRemove_ZeroConditions(t *testing.T) {
	s := strategy.CreateStrategy(10, 1, 1, condition.CreateConditions([]condition.ICondition{}))
	s2 := MutateRemoveCondition(s)
	if s2.Conditions().Length() != 0 {
		t.Error("AssertionError")
	}
}
