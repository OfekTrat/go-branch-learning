package condition_list

import (
	candle_stream "branch_learning/candle_stream"
	"branch_learning/condition"
	"math/rand"
	"strconv"
	"strings"
)

type IndicatorCondition struct {
	Indicator      string
	CandleIndex    int
	IndicatorValue float32
	Percentage     float32
	GreaterThan    bool
}

func (c IndicatorCondition) ConditionType() string {
	return "IndicatorCondition"
}

func (c IndicatorCondition) MeetsCondition(cs *candle_stream.CandleStream) bool {
	candle := cs.Get(c.CandleIndex)
	candleIndicatorValue := candle.Get(c.Indicator)

	if c.GreaterThan {
		return (candleIndicatorValue-c.IndicatorValue)/c.IndicatorValue > c.Percentage
	} else {
		return (candleIndicatorValue-c.IndicatorValue)/c.IndicatorValue < c.Percentage
	}
}

func (c IndicatorCondition) IsValidStreamSize(streamsize int) bool {
	return c.CandleIndex < streamsize
}

func (c IndicatorCondition) Mutate(streamsize int) condition.ICondition {
	// Everything but the indicator type can be changed
	randCandleIndex := rand.Intn(streamsize)
	randPercentage := c.Percentage + (rand.Float32()-0.5)*20
	randIndicatorValue := c.IndicatorValue + ((rand.Float32() - 0.5) * 20)
	isGreaterThan := rand.Intn(2)

	greaterThan := false
	if isGreaterThan == 1 {
		greaterThan = true
	}
	return IndicatorCondition{Indicator: c.Indicator, CandleIndex: randCandleIndex, IndicatorValue: randIndicatorValue, Percentage: randPercentage, GreaterThan: greaterThan}
}

func (c IndicatorCondition) Hash() string {
	return strings.Join(
		[]string{
			c.ConditionType(),
			c.Indicator,
			strconv.FormatInt(int64(c.CandleIndex), 10),
			strconv.FormatFloat(float64(c.IndicatorValue), 'e', -1, 32),
			strconv.FormatFloat(float64(c.Percentage), 'e', -1, 32),
			strconv.FormatBool(c.GreaterThan),
		}, "|")
}
