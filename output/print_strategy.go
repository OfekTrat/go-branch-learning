package output

import (
	"branch_learning/condition"
	st "branch_learning/strategy"
	"log"
	"strings"
)

func PrintStrategyConditions(strategy *st.Strategy) {
	conditionStrings := []string{}
	for _, cond := range strategy.Conditions().ToList() {
		conditionStrings = append(conditionStrings, condition.ConditionToJson(cond))
	}
	log.Println(strings.Join(conditionStrings, "\n"))
}
