package condition

import (
	"branch_learning/candle"
	candlestream "branch_learning/candle_stream"
	"testing"
)

func TestCreateConditions_MeetsConditions(t *testing.T) {
	candle1 := map[string]float32{"close": 2, "open": 1}
	candle2 := map[string]float32{"close": 1, "open": 2}
	c1 := DummyCondition{0, true, "c1"}
	c2 := DummyCondition{1, false, "c2"}
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
	c1 := DummyCondition{0, true, "c1"}
	c2 := DummyCondition{1, false, "c2"}
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
	cond1 := DummyCondition{1, false, "cond1"}
	cond2 := DummyCondition{2, true, "cond2"}
	cond3 := DummyCondition{3, false, "cond3"}
	cond4 := DummyCondition{3, false, "cond4"}
	conds := []ICondition{cond1, cond2}
	cs := CreateConditions(conds)
	cs.Add(cond3)
	if cs.Length() != 3 && !cs.GetByIndex(2).Equals(cond3) {
		t.Error("AssertionError")
	}

	cs.Add(cond2)
	if cs.Length() != 3 && !cs.GetByIndex(1).Equals(cond2) {
		t.Error("AssertionError")
	}
	cs.Add(cond4)
	if cs.Length() != 3 && !cs.GetByIndex(2).Equals(cond4) {
		t.Error("AssertionError")
	}
}

func TestConditions_AddToIndexSameHash(t *testing.T) {
	cond1 := DummyCondition{1, false, "cond1"}
	cond2 := DummyCondition{2, false, "cond2"}
	cond3 := DummyCondition{1, false, "cond3"}

	conditions := CreateConditions([]ICondition{cond1, cond2})
	conditions.SetInIndex(cond3, 1)

	if conditions.Length() != 2 {
		t.Log(conditions)
		t.Error("AssertionError")
	}

	if conditions.GetByHash(cond1.Hash()).Equals(cond1) {
		t.Log(conditions)
		t.Error("AssertionError")
	}
}

func TestLength(t *testing.T) {
	cond1 := DummyCondition{1, false, "cond1"}
	cond2 := DummyCondition{2, true, "cond2"}
	cond3 := DummyCondition{3, false, "cond3"}
	conds := []ICondition{cond1, cond2, cond3}
	cs := CreateConditions(conds)

	if cs.Length() != 3 {
		t.Error("AssertionError")
	}
}

func TestConditions_SetInIndex(t *testing.T) {
	cond1 := DummyCondition{1, false, "cond1"}
	cond2 := DummyCondition{1, false, "cond2"}
	conditions := CreateConditions([]ICondition{cond1})
	conditions.SetInIndex(cond2, 0)

	if conditions.Length() != 1 {
		t.Error("AssertionError")
	}
	if !conditions.GetByIndex(0).Equals(cond2) {
		t.Error("AssertionError")
	}
}

func TestConditions_RemoveByIndex(t *testing.T) {
	cond1 := DummyCondition{1, false, "cond1"}
	cond2 := DummyCondition{1, true, "cond2"}
	conditions := CreateConditions([]ICondition{cond1, cond2})
	conditions.RemoveByIndex(1)

	if conditions.Length() != 1 {
		t.Log(conditions)
		t.Error("AssertionError")
	}
	if conditions.GetByIndex(0).Equals(cond2) && !conditions.GetByIndex(0).Equals(cond1) {
		t.Log(conditions)
		t.Error("AssertionError")
	}
}
