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
	PivotPart   string
	Percentage  float64
	GreaterThan bool
}

func (c PivotPointCondition) ConditionType() string {
	return "PivotPointCondition"
}

func (c PivotPointCondition) MeetsCondition(cs *candle_stream.CandleStream) bool {
	candle := cs.Get(c.CandleIndex)
	pp_res_sup := candle.Get(c.PivotPart)

	if c.GreaterThan {
		return ((candle.Get(c.CandlePart)-pp_res_sup)/pp_res_sup)*100 > c.Percentage
	} else {
		return ((candle.Get(c.CandlePart)-pp_res_sup)/pp_res_sup)*100 < c.Percentage
	}
}

func (c PivotPointCondition) IsValidStreamSize(streamsize int) bool {
	return c.CandleIndex < streamsize
}

func (c PivotPointCondition) Mutate(streamsize int) icondition.ICondition {
	randIndex := rand.Intn(streamsize)
	randCandlePart := CandleParts[rand.Intn(len(CandleParts))]
	randPivotPart := PivotParts[rand.Intn(len(PivotParts))]
	isGreaterThan := rand.Intn(2)
	randPercentage := c.Percentage + (rand.Float64()-0.5)*20
	if isGreaterThan == 1 {
		return PivotPointCondition{CandleIndex: randIndex, CandlePart: randCandlePart, PivotPart: randPivotPart, Percentage: randPercentage, GreaterThan: true}
	} else {
		return PivotPointCondition{CandleIndex: randIndex, CandlePart: randCandlePart, PivotPart: randPivotPart, Percentage: randPercentage, GreaterThan: false}
	}
}

func (c PivotPointCondition) Hash() string {
	return strings.Join(
		[]string{
			c.ConditionType(),
			c.PivotPart,
			strconv.FormatInt(int64(c.CandleIndex), 10),
			c.CandlePart,
			strconv.FormatBool(c.GreaterThan),
		}, "|")
}

func (c PivotPointCondition) IsOverriddenBy(o icondition.ICondition) bool {
	other := o.(PivotPointCondition)
	if c.GreaterThan {
		return c.Percentage > other.Percentage
	} else {
		return c.Percentage < other.Percentage
	}
}

func (c PivotPointCondition) Equals(o icondition.ICondition) bool {
	other, ok := o.(PivotPointCondition)
	if !ok {
		return false
	}
	return other.PivotPart == c.PivotPart && other.CandleIndex == c.CandleIndex && other.CandlePart == c.CandlePart &&
		other.Percentage == c.Percentage && other.GreaterThan == c.GreaterThan
}
