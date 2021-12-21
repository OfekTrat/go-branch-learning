package condition_list

import (
	candle_stream "branch_learning/candle_stream"
	icondition "branch_learning/condition"
	"math/rand"
	"strconv"
	"strings"
)

type RSICondition struct {
	CandleIndex int
	RsiValue    float32
	GreaterThan bool
}

func (c RSICondition) ConditionType() string {
	return "RSI"
}

func (c RSICondition) MeetsCondition(cs *candle_stream.CandleStream) bool {
	candle := cs.Get(c.CandleIndex)
	rsi := candle.Get("rsi")

	if c.GreaterThan {
		return rsi > c.RsiValue
	} else {
		return rsi <= c.RsiValue
	}
}

func (c RSICondition) IsValidStreamSize(streamsize int) bool {
	return c.CandleIndex < streamsize
}

func (c RSICondition) Mutate(streamsize int) icondition.ICondition {
	randIndex := rand.Intn(streamsize)
	isGreaterThan := rand.Intn(2)
	randRSIValue := c.RsiValue + ((rand.Float32() - 0.5) * 20)
	if isGreaterThan == 1 {
		return RSICondition{CandleIndex: randIndex, RsiValue: randRSIValue, GreaterThan: true}
	} else {
		return RSICondition{CandleIndex: randIndex, RsiValue: randRSIValue, GreaterThan: false}
	}
}

func (c RSICondition) Hash() string {
	return strings.Join(
		[]string{
			c.ConditionType(),
			strconv.FormatInt(int64(c.CandleIndex), 10),
			strconv.FormatFloat(float64(c.RsiValue), 'e', -1, 32),
			strconv.FormatBool(c.GreaterThan),
		}, "|")
}
