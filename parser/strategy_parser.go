package parser

import (
	"branch_learning/condition"
	st "branch_learning/strategy"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	parametersPattern = `(?s)StopLoss: (?P<StopLoss>[\d\.]*).*TakeProfit: (?P<TakeProfit>[\d\.]*).*Window: (?P<Window>\d*)`
	conditionPattern  = `(?s)({.*?})+`
)

func ParseStrategy(strategyStr string) *st.Strategy {
	params := getParams(parametersPattern, strategyStr)
	validateParams(params)
	stoploss, slErr := strconv.ParseFloat(params["StopLoss"], 32)
	takeprofit, tpErr := strconv.ParseFloat(params["TakeProfit"], 32)
	window, wErr := strconv.ParseInt(params["Window"], 10, 32)

	if slErr != nil || tpErr != nil || wErr != nil {
		fmt.Println("Wrong values")
		os.Exit(1)
	}
	conditions := parseConditions(strategyStr)
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
		cond := ParseConditionFromJson(match)
		conditions.Add(cond)
	}
	return &conditions
}
