package output

import (
	"branch_learning/condition"
	st "branch_learning/strategy"
	"log"
	"strings"
)

func PrintStrategy(strategy *st.Strategy) {
	log.Printf("StopLoss: %v\n", strategy.StopLoss())
	log.Printf("TakeProfit: %v\n", strategy.TakeProfit())
	log.Printf("Window: %v\n", strategy.WindowSize())
	printStrategyConditions(strategy)
}

func printStrategyConditions(strategy *st.Strategy) {
	conditionStrings := []string{}
	for _, cond := range strategy.Conditions().ToList() {
		conditionStrings = append(conditionStrings, condition.ConditionToJson(cond))
	}
	log.Println(strings.Join(conditionStrings, "\n"))
}
