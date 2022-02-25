package parser

import (
	"branch_learning/condition"
	st "branch_learning/strategy"
	"encoding/json"
	"log"
	"os"
)

func ParseStrategy(strategyStr []byte) *st.Strategy {
	mappedStrategy := make(map[string]interface{})
	err := json.Unmarshal(strategyStr, &mappedStrategy)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	stopLoss := float32(mappedStrategy["stop_loss"].(float64))
	takeProfit := float32(mappedStrategy["take_profit"].(float64))
	window := int(mappedStrategy["window"].(float64))
	conditions := ParseConditions(mappedStrategy["conditions"].([]interface{}))
	strategy := st.CreateStrategy(0, 0, window, takeProfit, stopLoss, conditions) // The parser is used only in strategy tester which means its default id, generation are 0
	return strategy
}

func ParseConditions(conditions []interface{}) *condition.Conditions {
	conditionList := []condition.ICondition{}
	for _, cond := range conditions {
		mapCond := cond.(map[string]interface{})
		assertedCond := assertCond(mapCond)
		conditionList = append(conditionList, assertedCond)
	}
	return condition.CreateConditions(conditionList)
}
