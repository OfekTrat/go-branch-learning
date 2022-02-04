package parser

import (
	"branch_learning/condition"
	"branch_learning/condition_list"
	"fmt"
	"os"
)

func assertCond(condMap map[string]interface{}) condition.ICondition {
	var cond condition.ICondition
	switch condMap["type"] {
	case "CandleComparison":
		cond = condition_list.CandleComparisonCondition{
			CandleIndex1: int(condMap["CandleIndex1"].(float64)),
			CandlePart1:  condMap["CandlePart1"].(string),
			CandleIndex2: int(condMap["CandleIndex2"].(float64)),
			CandlePart2:  condMap["CandlePart2"].(string),
			Percentage:   float32(condMap["Percentage"].(float64)),
		}
	case "PivotPointCondition":
		cond = condition_list.PivotPointCondition{
			CandleIndex: int(condMap["CandleIndex"].(float64)),
			CandlePart:  condMap["CandlePart"].(string),
			Percentage:  float32(condMap["Percentage"].(float64)),
			GreaterThan: condMap["GreaterThan"].(bool),
			PivotPart:   condMap["PivotPart"].(string),
		}
	case "IndicatorCondition":
		cond = condition_list.IndicatorCondition{
			Indicator:      condMap["Indicator"].(string),
			CandleIndex:    int(condMap["CandleIndex"].(float64)),
			IndicatorValue: float32(condMap["IndicatorValue"].(float64)),
			Percentage:     float32(condMap["Percentage"].(float64)),
			GreaterThan:    condMap["GreaterThan"].(bool),
		}
	case "IndicatorCompareCondition":
		cond = condition_list.IndicatorCompareCondition{
			Indicator:    condMap["Indicator"].(string),
			CandleIndex1: int(condMap["CandleIndex1"].(float64)),
			CandleIndex2: int(condMap["CandleIndex2"].(float64)),
			Percentage:   float32(condMap["Percentage"].(float64)),
		}
	default:
		fmt.Printf("Please implement a parser that will include %s\n", condMap["type"])
		os.Exit(1)
	}
	return cond
}
