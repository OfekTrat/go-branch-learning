package strategy

import (
	cl "branch_learning/condition_list"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func CreateStrategyFromFile(strategyPath string) *Strategy {
	strategyBytes, err := ioutil.ReadFile(strategyPath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return createStrategyFromJsonBytes(strategyBytes)
}

func createStrategyFromJsonBytes(strategyBytes []byte) *Strategy {
	mappedStrategy := make(map[string]interface{})
	err := json.Unmarshal(strategyBytes, &mappedStrategy)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	stopLoss := mappedStrategy["stop_loss"].(float64)
	takeProfit := mappedStrategy["take_profit"].(float64)
	window := int(mappedStrategy["window"].(float64))
	conditions := cl.ParseConditions(mappedStrategy["conditions"].([]interface{}))
	strategy := CreateStrategy(0, 0, window, takeProfit, stopLoss, conditions) // The parser is used only in strategy tester which means its default id, generation are 0
	return strategy
}
