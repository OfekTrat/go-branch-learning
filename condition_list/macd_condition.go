package condition_list

import (
	candle_stream "branch_learning/candle_stream"
	"branch_learning/condition"
	"math/rand"
	"strconv"
	"strings"
)

type MACDCondition struct {
	CandleIndex int
	MacdValue   float32
	GreaterThan bool
}

func (c MACDCondition) ConditionType() string {
	return "MACD"
}

func (c MACDCondition) MeetsCondition(cs *candle_stream.CandleStream) bool {
	candle := cs.Get(c.CandleIndex)
	macd := candle.Get("macd")
	if c.GreaterThan {
		return macd > c.MacdValue
	} else {
		return macd <= c.MacdValue
	}
}

func (c MACDCondition) IsValidStreamSize(streamsize int) bool {
	return c.CandleIndex < streamsize
}

func (c MACDCondition) Mutate(streamsize int) condition.ICondition {
	randIndex := rand.Intn(streamsize)
	isGreaterThan := rand.Intn(2)
	randMacdValue := c.MacdValue + ((rand.Float32() - 0.5) * 20)

	if isGreaterThan == 1 {
		return MACDCondition{CandleIndex: randIndex, GreaterThan: true, MacdValue: randMacdValue}
	} else {
		return MACDCondition{CandleIndex: randIndex, GreaterThan: true, MacdValue: randMacdValue}
	}
}

func (c MACDCondition) Hash() string {
	return strings.Join(
		[]string{
			c.ConditionType(),
			strconv.FormatInt(int64(c.CandleIndex), 10),
			strconv.FormatFloat(float64(c.MacdValue), 'e', -1, 32),
			strconv.FormatBool(c.GreaterThan),
		}, "|")
}
