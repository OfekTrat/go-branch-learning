package condition

import (
	"branch_learning/candle"
	candlestream "branch_learning/candle_stream"
	"fmt"
	"testing"
)

func TestCreateConditions_MeetsConditions(t *testing.T) {
	candle1 := map[string]float32{"close": 2, "open": 1}
	candle2 := map[string]float32{"close": 1, "open": 2}
	c1 := DummyCondition{0, true}
	c2 := DummyCondition{1, false}
	stream := candlestream.CreateCandleStream([]candle.Candle{
		candle.CreateCandle(candle1),
		candle.CreateCandle(candle2),
	})

	condsList := []ICondition{c1, c2}
	conditions := CreateConditions(condsList)
	if !conditions.MeetsConditions(stream) {
		t.Error("Expected True Got False")
	}
}

func TestCreateConditions_DoesNotMeetConditions(t *testing.T) {
	candle1 := map[string]float32{"close": 1, "open": 2}
	candle2 := map[string]float32{"close": 2, "open": 1}
	c1 := DummyCondition{0, true}
	c2 := DummyCondition{1, false}
	stream := candlestream.CreateCandleStream([]candle.Candle{
		candle.CreateCandle(candle1),
		candle.CreateCandle(candle2),
	})

	condsList := []ICondition{c1, c2}
	conditions := CreateConditions(condsList)
	if conditions.MeetsConditions(stream) {
		t.Error("Expected False Got True")
	}
}

func TestAddCondition(t *testing.T) {
	cond1 := DummyCondition{1, false}
	cond2 := DummyCondition{2, true}
	cond3 := DummyCondition{3, false}
	conds := []ICondition{cond1, cond2}
	cs := CreateConditions(conds)
	cs.Add(cond3)
	fmt.Println(cs)
	if cs.Length() != 3 {
		t.Error("AssertionError")
	}

	cs.Add(cond2)
	fmt.Println(cs)
	if cs.Length() != 3 {
		t.Error("AssertionError")
	}
}

func TestLength(t *testing.T) {
	cond1 := DummyCondition{1, false}
	cond2 := DummyCondition{2, true}
	cond3 := DummyCondition{3, false}
	conds := []ICondition{cond1, cond2, cond3}
	cs := CreateConditions(conds)

	if cs.Length() != 3 {
		t.Error("AssertionError")
	}
}

func TestToList(t *testing.T) {
	cond1 := DummyCondition{1, false}
	cond2 := DummyCondition{2, true}
	cond3 := DummyCondition{3, false}
	conds := []ICondition{cond1, cond2, cond3}
	cs := CreateConditions(conds)

	if cs.Length() != 3 {
		t.Error("AssertionError")
	}
}
