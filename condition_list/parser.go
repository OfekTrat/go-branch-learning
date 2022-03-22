package condition_list

import (
	"branch_learning/condition"
	"fmt"
	"os"
)

func ParseConditions(conditions []interface{}) *condition.Conditions {
	conditionList := []condition.ICondition{}
	for _, cond := range conditions {
		mapCond := cond.(map[string]interface{})
		assertedCond := assertCond(mapCond)
		conditionList = append(conditionList, assertedCond)
	}
	return condition.CreateConditions(conditionList)
}

func assertCond(condMap map[string]interface{}) condition.ICondition {
	var cond condition.ICondition
	switch condMap["type"] {
	case "CandleComparison":
		cond = CandleComparisonCondition{
			CandleIndex1: int(condMap["CandleIndex1"].(float64)),
			CandlePart1:  condMap["CandlePart1"].(string),
			CandleIndex2: int(condMap["CandleIndex2"].(float64)),
			CandlePart2:  condMap["CandlePart2"].(string),
			Percentage:   condMap["Percentage"].(float64),
		}
	case "PivotPointCondition":
		cond = PivotPointCondition{
			CandleIndex: int(condMap["CandleIndex"].(float64)),
			CandlePart:  condMap["CandlePart"].(string),
			Percentage:  condMap["Percentage"].(float64),
			GreaterThan: condMap["GreaterThan"].(bool),
			PivotPart:   condMap["PivotPart"].(string),
		}
	case "IndicatorCondition":
		cond = IndicatorCondition{
			Indicator:      condMap["Indicator"].(string),
			CandleIndex:    int(condMap["CandleIndex"].(float64)),
			IndicatorValue: condMap["IndicatorValue"].(float64),
			Percentage:     condMap["Percentage"].(float64),
			GreaterThan:    condMap["GreaterThan"].(bool),
		}
	case "IndicatorCompareCondition":
		cond = IndicatorCompareCondition{
			Indicator:    condMap["Indicator"].(string),
			CandleIndex1: int(condMap["CandleIndex1"].(float64)),
			CandleIndex2: int(condMap["CandleIndex2"].(float64)),
			Percentage:   condMap["Percentage"].(float64),
		}
	default:
		fmt.Printf("Please implement a parser that will include %s\n", condMap["type"])
		os.Exit(1)
	}
	return cond
}
