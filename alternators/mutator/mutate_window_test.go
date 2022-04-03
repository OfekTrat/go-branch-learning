package mutator

import (
	"branch_learning/condition"
	"branch_learning/strategy"
	"testing"
)

func TestMutateWindow(t *testing.T) {
	s := strategy.CreateStrategy(0, 0, 100, 1, 1, condition.CreateConditions([]condition.ICondition{}))
	s2 := MutateWindowSize(1, 1, s)

	if s2.WindowSize() == s.WindowSize() {
		t.Error("Assertion Error")
	}
}
