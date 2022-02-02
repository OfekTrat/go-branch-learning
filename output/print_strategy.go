package output

import (
	st "branch_learning/strategy"
	"encoding/json"
	"log"
	"os"
)

func PrintStrategy(strategy *st.Strategy) {
	mappedStrategy := make(map[string]interface{})
	mappedStrategy["stop_loss"] = strategy.StopLoss()
	mappedStrategy["take_profit"] = strategy.TakeProfit()
	mappedStrategy["window"] = strategy.WindowSize()
	// mappedStrategy["conditions"] = getConditionStrList(strategy)
	mappedStrategy["conditions"] = getConditionStrList(strategy)
	data, err := json.MarshalIndent(mappedStrategy, "", "    ")

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	log.Println(string(data))
}

func getConditionStrList(strategy *st.Strategy) []map[string]interface{} {
	var err error
	var conditionJsonNoType []byte
	conditions := []map[string]interface{}{}

	for i := 0; i < strategy.Conditions().Length(); i++ {
		cond := strategy.Conditions().GetByIndex(i)
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
		conditions = append(conditions, mappedCond)
	}
	return conditions
}
