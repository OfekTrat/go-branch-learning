package mutator

import (
	condition "branch_learning/condition"
	condition_list "branch_learning/condition_list"
	"branch_learning/strategy"
	"testing"
)

func TestMutateAddCondition(t *testing.T) {
	redCondition := condition_list.CandleComparisonCondition{CandleIndex1: 0, CandlePart1: "open", CandleIndex2: 0, CandlePart2: "close", Percentage: 0}
	s := strategy.CreateStrategy(0, 0, 10, 1, 1, condition.CreateConditions([]condition.ICondition{redCondition}))

	newS := MutateAddCondition(1, 1, s)

	if newS.Conditions().Length() != s.Conditions().Length()+1 {
		t.Error("AssertionError")
	}
}

func TestMutateRemoveCondition(t *testing.T) {
	cond := condition_list.CandleComparisonCondition{CandleIndex1: 0, CandlePart1: "open", CandleIndex2: 0, CandlePart2: "close", Percentage: 0}
	s := strategy.CreateStrategy(0, 0, 10, 1, 1, condition.CreateConditions([]condition.ICondition{cond}))
	newS := MutateRemoveCondition(1, 1, s)

	if newS.Conditions().Length() != s.Conditions().Length()-1 {
		t.Error("AssertionError")
	}
}

func TestMutateRemove_ZeroConditions(t *testing.T) {
	s := strategy.CreateStrategy(0, 0, 10, 1, 1, condition.CreateConditions([]condition.ICondition{}))
	s2 := MutateRemoveCondition(1, 1, s)
	if s2.Conditions().Length() != 0 {
		t.Error("AssertionError")
	}
}
