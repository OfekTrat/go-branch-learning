package condition_list

import (
	candle_stream "branch_learning/candle_stream"
	"branch_learning/condition"
	"math/rand"
	"strconv"
	"strings"
)

type RSICompareCondition struct {
	CandleIndex1 int
	CandleIndex2 int
	Percentage   float32
}

func (c RSICompareCondition) ConditionType() string {
	return "RSICompare"
}

func (c RSICompareCondition) MeetsCondition(cs *candle_stream.CandleStream) bool {
	candle1 := cs.Get(c.CandleIndex1)
	candle2 := cs.Get(c.CandleIndex2)

	rsi1 := candle1.Get("rsi")
	rsi2 := candle2.Get("rsi")

	return (rsi1-rsi2)/rsi1 > c.Percentage
}

func (c RSICompareCondition) IsValidStreamSize(streamsize int) bool {
	return c.CandleIndex1 < streamsize && c.CandleIndex2 < streamsize
}

func (c RSICompareCondition) Mutate(streamsize int) condition.ICondition {
	randCandle := rand.Intn(2)
	randIndex := rand.Intn(streamsize)
	randPercentage := c.Percentage + (rand.Float32()-0.5)*20
	if randCandle == 0 {
		return RSICompareCondition{CandleIndex1: randIndex, CandleIndex2: c.CandleIndex2, Percentage: randPercentage}
	} else {
		return RSICompareCondition{CandleIndex1: c.CandleIndex1, CandleIndex2: randIndex, Percentage: randPercentage}
	}
}

func (c RSICompareCondition) Hash() string {
	return strings.Join(
		[]string{
			c.ConditionType(),
			strconv.FormatInt(int64(c.CandleIndex1), 10),
			strconv.FormatInt(int64(c.CandleIndex2), 10),
			strconv.FormatFloat(float64(c.Percentage), 'e', -1, 32),
		},
		"|")
}
