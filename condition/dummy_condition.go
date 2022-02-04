package condition

import (
	candle_stream "branch_learning/candle_stream"
	"math/rand"
	"strconv"
	"strings"
)

type DummyCondition struct {
	CandleIndex int    `json:"candle_index"`
	IsGreen     bool   `json:"is_green"`
	Text        string `json:"text"`
}

func (c DummyCondition) ConditionType() string {
	return "CandleType"
}

func (c DummyCondition) MeetsCondition(cs *candle_stream.CandleStream) bool {
	candle := cs.Get(c.CandleIndex)
	if c.IsGreen {
		return candle.Get("close") > candle.Get("open")
	}
	return candle.Get("open") > candle.Get("close")
}

func (c DummyCondition) IsValidStreamSize(streamsize int) bool {
	return c.CandleIndex < streamsize
}

func (c DummyCondition) Mutate(streamsize int) ICondition {
	randIndex := rand.Intn(streamsize)
	isGreen := rand.Intn(2)

	if isGreen == 1 {
		return DummyCondition{CandleIndex: randIndex, IsGreen: true}
	} else {
		return DummyCondition{CandleIndex: randIndex, IsGreen: false}
	}
}

func (c DummyCondition) Hash() string {
	return strings.Join([]string{c.ConditionType(), strconv.FormatInt(int64(c.CandleIndex), 10), strconv.FormatBool(c.IsGreen)}, "|")
}

func (c DummyCondition) IsOverriddenBy(o ICondition) bool {
	return true
}

func (c DummyCondition) Equals(o ICondition) bool {
	other, ok := o.(DummyCondition)
	if !ok {
		return false
	}
	return c.IsGreen == other.IsGreen && c.CandleIndex == other.CandleIndex && c.Text == other.Text
}
