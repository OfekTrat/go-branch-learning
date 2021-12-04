package condition

import (
	condition_list "branch_learning/condition/condition_list"
	"testing"
)

const (
	ToJsonTest = "{\n\t\"candle_index\": 2,\n\t\"type\": \"RedCondition\"\n}"
)

func TestRedConditionToJson(t *testing.T) {
	condition := condition_list.RedCondition{CandleIndex: 2}
	conditionJson := ConditionToJson(&condition)

	if conditionJson != ToJsonTest {
		t.Log("Assertion Error:")
		t.Logf("Got: %s", conditionJson)
		t.Logf("Expected %s", ToJsonTest)
		t.Error("AssertionError")
	}
}
