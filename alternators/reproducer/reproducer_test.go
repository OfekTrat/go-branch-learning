package reproducer

import (
	"branch_learning/condition"
	condition_list "branch_learning/condition_list"
	"branch_learning/strategy"
	"testing"
)

func TestReproducer(t *testing.T) {
	cond1 := condition_list.CandleTypeCondition{CandleIndex: 1, IsGreen: false}
	cond2 := condition_list.CandleTypeCondition{CandleIndex: 2, IsGreen: false}
	cond3 := condition_list.CandleTypeCondition{CandleIndex: 3, IsGreen: false}
	cond4 := condition_list.CandleTypeCondition{CandleIndex: 4, IsGreen: false}
	cond5 := condition_list.CandleTypeCondition{CandleIndex: 5, IsGreen: false}

	s1 := strategy.CreateStrategy(10, 1, 1, condition.CreateConditions([]condition.ICondition{cond1, cond2, cond3}))
	s2 := strategy.CreateStrategy(10, 1, 1, condition.CreateConditions([]condition.ICondition{cond4, cond5}))
	s3 := reproduceByNConditions(s1, s2, 1)

	t.Log(s1.Conditions().ToList())
	t.Log(s2.Conditions().ToList())
	t.Log(s3.Conditions().ToList())

	if s3.Conditions().Length() != s1.Conditions().Length()+1 {
		t.Error("AssertionError1")
	}
}
