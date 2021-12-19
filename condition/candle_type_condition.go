package condition

import (
	candle_stream "branch_learning/candle_stream"
	"math/rand"
	"strconv"
	"strings"
)

type CandleTypeCondition struct {
	CandleIndex int  `json:"candle_index"`
	IsGreen     bool `json:"is_green"`
}

func (c CandleTypeCondition) ConditionType() string {
	return "CandleType"
}

func (c CandleTypeCondition) MeetsCondition(cs *candle_stream.CandleStream) bool {
	candle := cs.Get(c.CandleIndex)
	if c.IsGreen {
		return candle.Get("close") > candle.Get("open")
	}
	return candle.Get("open") > candle.Get("close")
}

func (c CandleTypeCondition) IsValidStreamSize(streamsize int) bool {
	return c.CandleIndex < streamsize
}

func (c CandleTypeCondition) Mutate(streamsize int) ICondition {
	randIndex := rand.Intn(streamsize)
	isGreen := rand.Intn(2)

	if isGreen == 1 {
		return CandleTypeCondition{CandleIndex: randIndex, IsGreen: true}
	} else {
		return CandleTypeCondition{CandleIndex: randIndex, IsGreen: false}
	}
}

func (c CandleTypeCondition) Hash() string {
	return strings.Join([]string{c.ConditionType(), strconv.FormatInt(int64(c.CandleIndex), 10), strconv.FormatBool(c.IsGreen)}, "|")
}
