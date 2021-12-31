package condition_list

import (
	candle_stream "branch_learning/candle_stream"
	icondition "branch_learning/condition"
	"math/rand"
	"strconv"
	"strings"
)

type PivotPointCondition struct {
	CandleIndex int
	CandlePart  string
	Percentage  float32
	GreaterThan bool
}

func (c PivotPointCondition) ConditionType() string {
	return "PivotPointCondition"
}

func (c PivotPointCondition) MeetsCondition(cs *candle_stream.CandleStream) bool {
	candle := cs.Get(c.CandleIndex)
	pivotpoint := candle.Get("pivot")

	if c.GreaterThan {
		return (candle.Get(c.CandlePart)-pivotpoint)/pivotpoint > c.Percentage
	} else {
		return (candle.Get(c.CandlePart)-pivotpoint)/pivotpoint < c.Percentage
	}
}

func (c PivotPointCondition) IsValidStreamSize(streamsize int) bool {
	return c.CandleIndex < streamsize
}

func (c PivotPointCondition) Mutate(streamsize int) icondition.ICondition {
	randIndex := rand.Intn(streamsize)
	randPart := candleParts[rand.Intn(len(candleParts))]
	isGreaterThan := rand.Intn(2)
	randPercentage := c.Percentage + (rand.Float32()-0.5)*20
	if isGreaterThan == 1 {
		return PivotPointCondition{CandleIndex: randIndex, CandlePart: randPart, Percentage: randPercentage, GreaterThan: true}
	} else {
		return PivotPointCondition{CandleIndex: randIndex, CandlePart: randPart, Percentage: randPercentage, GreaterThan: false}
	}
}

func (c PivotPointCondition) Hash() string {
	return strings.Join(
		[]string{
			c.ConditionType(),
			strconv.FormatInt(int64(c.CandleIndex), 10),
			c.CandlePart,
			strconv.FormatFloat(float64(c.Percentage), 'e', -1, 32),
			strconv.FormatBool(c.GreaterThan),
		}, "|")
}
