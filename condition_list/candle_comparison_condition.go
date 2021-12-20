package condition_list

import (
	candle_stream "branch_learning/candle_stream"
	icondition "branch_learning/condition"
	"math/rand"
	"strconv"
	"strings"
)

var (
	candleParts = []string{"open", "high", "close", "low"}
)

type CandleComparisonCondition struct {
	CandleIndex1 int
	CandlePart1  string
	CandleIndex2 int
	CandlePart2  string
	Percentage   float32
}

func (c CandleComparisonCondition) ConditionType() string {
	return "CandleComparison"
}

func (c CandleComparisonCondition) MeetsCondition(cs *candle_stream.CandleStream) bool {
	candle1 := cs.Get(c.CandleIndex1)
	candle2 := cs.Get(c.CandleIndex2)

	candle1Value := candle1.Get(c.CandlePart1)
	candle2Value := candle2.Get(c.CandlePart2)

	return (candle1Value-candle2Value)/candle1Value > c.Percentage
}

func (c CandleComparisonCondition) IsValidStreamSize(streamsize int) bool {
	return c.CandleIndex1 < streamsize && c.CandleIndex2 < streamsize
}

func (c CandleComparisonCondition) Mutate(streamsize int) icondition.ICondition {
	randCandle := rand.Intn(2)
	randCandlePart := rand.Intn(len(candleParts))
	randIndex := rand.Intn(streamsize)
	randPercentage := (rand.Float32() - 0.5) * 20

	if randCandle == 0 {
		return CandleComparisonCondition{
			CandleIndex1: randIndex,
			CandlePart1:  candleParts[randCandlePart],
			CandleIndex2: c.CandleIndex2,
			CandlePart2:  c.CandlePart2,
			Percentage:   c.Percentage + randPercentage,
		}
	} else {
		return CandleComparisonCondition{
			CandleIndex1: c.CandleIndex1,
			CandlePart1:  c.CandlePart1,
			CandleIndex2: randIndex,
			CandlePart2:  candleParts[randCandlePart],
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
			strconv.FormatFloat(float64(c.Percentage), 'e', -1, 32),
		}, "|")
}
