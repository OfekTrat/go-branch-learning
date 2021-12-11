package reproducer

import (
	"branch_learning/condition"
	"branch_learning/strategy"
	"testing"
)

func TestReproducer(t *testing.T) {
	cond1 := condition.RedCondition{CandleIndex: 1}
	cond2 := condition.RedCondition{CandleIndex: 2}
	cond3 := condition.RedCondition{CandleIndex: 3}
	cond4 := condition.RedCondition{CandleIndex: 4}
	cond5 := condition.RedCondition{CandleIndex: 5}

	s1 := strategy.CreateStrategy(10, 1, 1, []condition.ICondition{cond1, cond2, cond3})
	s2 := strategy.CreateStrategy(10, 1, 1, []condition.ICondition{cond4, cond5})
	s3 := reproduceByIndexList(s1, s2, []int{1})

	if len(s3.Conditions()) != len(s1.Conditions())+1 {
		t.Error("AssertionError")
	}
	if s3.Conditions()[3] != s2.Conditions()[1] {
		t.Error("AssertionError")
	}
}
