package mutator

import (
	"branch_learning/condition"
	"branch_learning/strategy"
	"testing"
)

func TestMutateWindow(t *testing.T) {
	s := strategy.CreateStrategy(100, 1, 1, []condition.ICondition{})
	s2 := MutateWindowSize(s)

	if s2.WindowSize() == s.WindowSize() {
		t.Error("Assertion Error")
	}
}
