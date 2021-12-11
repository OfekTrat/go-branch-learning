package mutator

import (
	"branch_learning/condition"
	"branch_learning/strategy"
	"testing"
)

func TestMutateStopLoss(t *testing.T) {
	s := strategy.CreateStrategy(10, 1, 1, []condition.ICondition{})
	s2 := MutateStopLoss(s)

	if s2.StopLoss() == s.StopLoss() {
		t.Error("AssertionError")
	}
}

func TestMutateTakeProfit(t *testing.T) {
	s := strategy.CreateStrategy(10, 1, 1, []condition.ICondition{})
	s2 := MutateTakeProfit(s)
	if s2.TakeProfit() == s.TakeProfit() {
		t.Error("AssertionError")
	}
}
