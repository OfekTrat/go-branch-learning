package condition

import (
	"testing"
)

const (
	ToJsonTest = "{\n\t\"candle_index\": 2,\n\t\"type\": \"RedCondition\"\n}"
)

func TestRedConditionToJson(t *testing.T) {
	c := RedCondition{CandleIndex: 2}
	conditionJson := ConditionToJson(c)

	if conditionJson != ToJsonTest {
		t.Log("Assertion Error:")
		t.Logf("Got: %s", conditionJson)
		t.Logf("Expected %s", ToJsonTest)
		t.Error("AssertionError")
	}
}
