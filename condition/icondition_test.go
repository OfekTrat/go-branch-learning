package condition

import (
	"testing"
)

const (
	ToJsonTest = "{\n\t\"candle_index\": 2,\n\t\"is_green\": false,\n\t\"text\": \"1234\",\n\t\"type\": \"CandleType\"\n}"
)

func TestRedConditionToJson(t *testing.T) {
	c := DummyCondition{CandleIndex: 2, IsGreen: false, Text: "1234"}
	conditionJson := ConditionToJson(c)

	if conditionJson != ToJsonTest {
		t.Log("Assertion Error:")
		t.Logf("Got: %s", conditionJson)
		t.Logf("Expected %s", ToJsonTest)
		t.Error("AssertionError")
	}
}
