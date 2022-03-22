package condition_list

import (
	candle_stream "branch_learning/candle_stream"
	icondition "branch_learning/condition"
	"math/rand"
	"strconv"
	"strings"
)

type CandleComparisonCondition struct {
	CandleIndex1 int
	CandlePart1  string
	CandleIndex2 int
	CandlePart2  string
	Percentage   float64
}

func (c CandleComparisonCondition) ConditionType() string {
	return "CandleComparison"
}

func (c CandleComparisonCondition) MeetsCondition(cs *candle_stream.CandleStream) bool {
	candle1 := cs.Get(c.CandleIndex1)
	candle2 := cs.Get(c.CandleIndex2)

	candle1Value := candle1.Get(c.CandlePart1)
	candle2Value := candle2.Get(c.CandlePart2)

	return ((candle1Value-candle2Value)/candle2Value)*100 > c.Percentage
}

func (c CandleComparisonCondition) IsValidStreamSize(streamsize int) bool {
	return c.CandleIndex1 < streamsize && c.CandleIndex2 < streamsize
}

func (c CandleComparisonCondition) Mutate(streamsize int) icondition.ICondition {
	randCandle := rand.Intn(2)
	randCandlePart := rand.Intn(len(CandleParts))
	randIndex := rand.Intn(streamsize)
	randPercentage := (rand.Float64() - 0.5) * 20

	if randCandle == 0 {
		return CandleComparisonCondition{
			CandleIndex1: randIndex,
			CandlePart1:  CandleParts[randCandlePart],
			CandleIndex2: c.CandleIndex2,
			CandlePart2:  c.CandlePart2,
			Percentage:   c.Percentage + randPercentage,
		}
	} else {
		return CandleComparisonCondition{
			CandleIndex1: c.CandleIndex1,
			CandlePart1:  c.CandlePart1,
			CandleIndex2: randIndex,
			CandlePart2:  CandleParts[randCandlePart],
			Percentage:   c.Percentage + randPercentage,
		}
	}
}

func (c CandleComparisonCondition) Hash() string {
	return strings.Join(
		[]string{
			c.ConditionType(),
			strconv.FormatInt(int64(c.CandleIndex1), 10),
			c.CandlePart1,
			strconv.FormatInt(int64(c.CandleIndex2), 10),
			c.CandlePart2,
		}, "|")
}

func (c CandleComparisonCondition) IsOverriddenBy(o icondition.ICondition) bool {
	other := o.(CandleComparisonCondition)
	return c.Percentage > other.Percentage
}

func (c CandleComparisonCondition) Equals(o icondition.ICondition) bool {
	other, ok := o.(CandleComparisonCondition)

	if !ok {
		return false
	}
	return other.Percentage == c.Percentage && other.CandleIndex1 == c.CandleIndex1 && other.CandleIndex2 == c.CandleIndex2 &&
		other.CandlePart1 == c.CandlePart1 && other.CandlePart2 == c.CandlePart2

}
