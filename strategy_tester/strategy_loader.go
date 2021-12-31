package tester

import (
	"branch_learning/condition"
	"branch_learning/condition_list"
	st "branch_learning/strategy"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	parametersPattern = `(?s)StopLoss: (?P<StopLoss>[\d\.]*).*TakeProfit: (?P<TakeProfit>[\d\.]*).*Window: (?P<Window>\d*)`
	conditionPattern  = `(?s)({.*?})+`
)

func loadStrategyFromFile(strategyFile string) *st.Strategy {
	fileData, err := os.ReadFile(strategyFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return parseStrategy(string(fileData))
}

func parseStrategy(fileData string) *st.Strategy {
	params := getParams(parametersPattern, fileData)
	validateParams(params)
	stoploss, slErr := strconv.ParseFloat(params["StopLoss"], 32)
	takeprofit, tpErr := strconv.ParseFloat(params["TakeProfit"], 32)
	window, wErr := strconv.ParseInt(params["Window"], 10, 32)

	if slErr != nil || tpErr != nil || wErr != nil {
		fmt.Println("Wrong values")
		os.Exit(1)
	}
	conditions := parseConditions(fileData)
	return st.CreateStrategy(int(window), float32(takeprofit), float32(stoploss), conditions)
}

func getParams(pattern, data string) (paramsMap map[string]string) {
	var compRegEx = regexp.MustCompile(pattern)
	match := compRegEx.FindStringSubmatch(data)

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}

func validateParams(params map[string]string) {
	if params["StopLoss"] == "" || params["TakeProfit"] == "" || params["Window"] == "" {
		fmt.Println("Missing parameters - check file")
		os.Exit(1)
	}
}

func parseConditions(data string) *condition.Conditions {
	compRegex := regexp.MustCompile(conditionPattern)
	matches := compRegex.FindAllString(data, -1)
	conditions := condition.Conditions{}
	for _, match := range matches {
		cond := parseStringCondition(match)
		conditions.Add(cond)
	}
	return &conditions
}

func parseStringCondition(condStr string) condition.ICondition {
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
	}
	return cond
}
