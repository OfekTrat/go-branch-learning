package parser

import (
	"branch_learning/condition"
	"branch_learning/condition_list"
	"encoding/json"
	"fmt"
	"os"
)

func ParseConditionFromJson(condStr string) condition.ICondition {
	condMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(condStr), &condMap)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return assertCond(condMap)
}

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
	case "MACDCompare":
		cond = condition_list.MACDCompareCondition{
			CandleIndex1: condMap["CandleIndex1"].(int),
			CandleIndex2: condMap["CandleIndex2"].(int),
			Percentage:   condMap["Percentage"].(float32),
		}
	case "MACD":
		cond = condition_list.MACDCondition{
			CandleIndex: int(condMap["CandleIndex"].(float64)),
			MacdValue:   float32(condMap["MacdValue"].(float64)),
			GreaterThan: condMap["GreaterThan"].(bool),
		}
	case "RSICompare":
		cond = condition_list.RSICompareCondition{
			CandleIndex1: condMap["CandleIndex1"].(int),
			CandleIndex2: condMap["CandleIndex2"].(int),
			Percentage:   condMap["Percentage"].(float32),
		}
	case "RSI":
		cond = condition_list.RSICondition{
			CandleIndex: condMap["CandleIndex1"].(int),
			RsiValue:    condMap["RsiValue"].(float32),
			GreaterThan: condMap["GreaterThan"].(bool),
		}
	case "VolumeCompare":
		cond = condition_list.VolumeCompareCondition{
			CandleIndex1: condMap["CandleIndex1"].(int),
			CandleIndex2: condMap["CandleIndex2"].(int),
			Percentage:   condMap["Percentage"].(float32),
		}
	default:
		fmt.Printf("Please implement a parser that will include %s", condMap["type"])
		os.Exit(1)
	}
	return cond
}
