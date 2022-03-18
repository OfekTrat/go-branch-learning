package condition_list

import (
	candle_stream "branch_learning/candle_stream"
	"branch_learning/condition"
	"math/rand"
	"strconv"
	"strings"
)

type IndicatorCompareCondition struct {
	Indicator    string
	CandleIndex1 int
	CandleIndex2 int
	Percentage   float64
}

func (c IndicatorCompareCondition) ConditionType() string {
	return "IndicatorCompareCondition"
}

func (c IndicatorCompareCondition) MeetsCondition(cs *candle_stream.CandleStream) bool {
	candle1 := cs.Get(c.CandleIndex1)
	candle2 := cs.Get(c.CandleIndex2)
	candleIndicatorValue1 := candle1.Get(c.Indicator)
	candleIndicatorValue2 := candle2.Get(c.Indicator)

	return ((candleIndicatorValue1-candleIndicatorValue2)/candleIndicatorValue2)*100 > c.Percentage
}

func (c IndicatorCompareCondition) IsValidStreamSize(streamsize int) bool {
	return c.CandleIndex1 < streamsize && c.CandleIndex2 < streamsize
}

func (c IndicatorCompareCondition) Mutate(streamsize int) condition.ICondition {
	// Everything but the indicator type can be changed
	randCandleIndex := rand.Intn(streamsize)
	randCandle := rand.Intn(2)
	randPercentage := c.Percentage + (rand.Float64()-0.5)*20

	if randCandle == 0 {
		return IndicatorCompareCondition{
			Indicator:    c.Indicator,
			CandleIndex1: randCandleIndex,
			CandleIndex2: c.CandleIndex2,
			Percentage:   randPercentage,
		}
	} else {
		return IndicatorCompareCondition{
			Indicator:    c.Indicator,
			CandleIndex1: c.CandleIndex1,
			CandleIndex2: randCandleIndex,
			Percentage:   randPercentage,
		}
	}

}

func (c IndicatorCompareCondition) Hash() string {
	return strings.Join(
		[]string{
			c.ConditionType(),
			c.Indicator,
			strconv.FormatInt(int64(c.CandleIndex1), 10),
			strconv.FormatInt(int64(c.CandleIndex2), 10),
			strconv.FormatFloat(float64(c.Percentage), 'e', -1, 32),
		}, "|")
}

func (c IndicatorCompareCondition) IsOverriddenBy(o condition.ICondition) bool {
	other := o.(IndicatorCompareCondition)
	return c.Percentage > other.Percentage
}

func (c IndicatorCompareCondition) Equals(o condition.ICondition) bool {
	other, ok := o.(IndicatorCompareCondition)
	if !ok {
		return false
	}
	return other.Indicator == c.Indicator && other.CandleIndex1 == c.CandleIndex1 && other.CandleIndex2 == c.CandleIndex2 &&
		other.Percentage == c.Percentage
}
