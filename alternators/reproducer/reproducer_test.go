package reproducer

import (
	"branch_learning/condition"
	condition_list "branch_learning/condition_list"
	"branch_learning/strategy"
	"fmt"
	"testing"
)

func TestReproducer(t *testing.T) {
	cond1 := condition_list.CandleComparisonCondition{CandleIndex1: 1, CandlePart1: "open", CandleIndex2: 1, CandlePart2: "close", Percentage: 0}
	cond2 := condition_list.CandleComparisonCondition{CandleIndex1: 2, CandlePart1: "open", CandleIndex2: 2, CandlePart2: "close", Percentage: 0}
	cond3 := condition_list.CandleComparisonCondition{CandleIndex1: 3, CandlePart1: "open", CandleIndex2: 3, CandlePart2: "close", Percentage: 0}
	cond4 := condition_list.CandleComparisonCondition{CandleIndex1: 4, CandlePart1: "open", CandleIndex2: 4, CandlePart2: "close", Percentage: 0}
	cond5 := condition_list.CandleComparisonCondition{CandleIndex1: 5, CandlePart1: "open", CandleIndex2: 5, CandlePart2: "close", Percentage: 0}

	s1 := strategy.CreateStrategy(0, 0, 10, 1, 1, condition.CreateConditions([]condition.ICondition{cond1, cond2, cond3}))
	s2 := strategy.CreateStrategy(0, 0, 10, 1, 1, condition.CreateConditions([]condition.ICondition{cond4, cond5}))
	s3 := reproduceByNConditions(1, 1, s1, s2, 1)

	if s3.Conditions().Length() != s1.Conditions().Length()+1 {
		t.Error("AssertionError")
	}

	if s3.Generation() != 1 || s3.Id() != 1 {
		fmt.Println(s3.Generation(), s3.Id())
		t.Error("AssertionError")
	}
}
