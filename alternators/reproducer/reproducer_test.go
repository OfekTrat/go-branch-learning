package reproducer

import (
	"branch_learning/condition"
	"branch_learning/strategy"
	"testing"
)

func TestReproducer(t *testing.T) {
	cond1 := condition.CandleTypeCondition{CandleIndex: 1, IsGreen: false}
	cond2 := condition.CandleTypeCondition{CandleIndex: 2, IsGreen: false}
	cond3 := condition.CandleTypeCondition{CandleIndex: 3, IsGreen: false}
	cond4 := condition.CandleTypeCondition{CandleIndex: 4, IsGreen: false}
	cond5 := condition.CandleTypeCondition{CandleIndex: 5, IsGreen: false}

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
