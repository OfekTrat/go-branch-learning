package strategy

import (
	"branch_learning/condition"
	"encoding/json"
	"log"
	"os"
)

func (strategy *Strategy) ToJsonString() string {
	mappedStrategy := make(map[string]interface{})
	mappedStrategy["generation"] = strategy.Generation()
	mappedStrategy["id"] = strategy.Id()
	mappedStrategy["stop_loss"] = strategy.StopLoss()
	mappedStrategy["take_profit"] = strategy.TakeProfit()
	mappedStrategy["window"] = strategy.WindowSize()
	mappedStrategy["conditions"] = mapConditions(strategy.Conditions())
	data, err := json.MarshalIndent(mappedStrategy, "", "    ")

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return string(data)
}

func mapConditions(conditions *condition.Conditions) []map[string]interface{} {
	var err error
	var conditionJsonNoType []byte
	mappedConditions := []map[string]interface{}{}

	for i := 0; i < conditions.Length(); i++ {
		cond := conditions.GetByIndex(i)
		mappedCond := make(map[string]interface{})
		conditionJsonNoType, err = json.Marshal(cond)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		err = json.Unmarshal(conditionJsonNoType, &mappedCond)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		mappedCond["type"] = cond.ConditionType()
		mappedConditions = append(mappedConditions, mappedCond)
	}
	return mappedConditions
}
