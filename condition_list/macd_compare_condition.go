package condition_list

import (
	candle_stream "branch_learning/candle_stream"
	"branch_learning/condition"
	"math/rand"
	"strconv"
	"strings"
)

type MACDCompareCondition struct {
	CandleIndex1 int
	CandleIndex2 int
	Percentage   float32
}

func (c MACDCompareCondition) ConditionType() string {
	return "MACDCompare"
}

func (c MACDCompareCondition) MeetsCondition(cs *candle_stream.CandleStream) bool {
	candle1 := cs.Get(c.CandleIndex1)
	candle2 := cs.Get(c.CandleIndex2)

	rsi1 := candle1.Get("macd")
	rsi2 := candle2.Get("macd")

	return (rsi1-rsi2)/rsi1 > c.Percentage
}

func (c MACDCompareCondition) IsValidStreamSize(streamsize int) bool {
	return c.CandleIndex1 < streamsize && c.CandleIndex2 < streamsize
}

func (c MACDCompareCondition) Mutate(streamsize int) condition.ICondition {
	randCandle := rand.Intn(2)
	randIndex := rand.Intn(streamsize)
	randPercentage := c.Percentage + (rand.Float32()-0.5)*20
	if randCandle == 0 {
		return MACDCompareCondition{CandleIndex1: randIndex, CandleIndex2: c.CandleIndex2, Percentage: randPercentage}
	} else {
		return MACDCompareCondition{CandleIndex1: c.CandleIndex1, CandleIndex2: randIndex, Percentage: randPercentage}
	}
}

func (c MACDCompareCondition) Hash() string {
	return strings.Join(
		[]string{
			c.ConditionType(),
			strconv.FormatInt(int64(c.CandleIndex1), 10),
			strconv.FormatInt(int64(c.CandleIndex2), 10),
			strconv.FormatFloat(float64(c.Percentage), 'e', -1, 32),
		},
		"|")
}
