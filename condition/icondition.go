package condition

import (
	candleStream "branch_learning/candle_stream"
	"encoding/json"
)

type ICondition interface {
	ConditionType() string
	MeetsCondition(*candleStream.CandleStream) bool
	IsValidStreamSize(int) bool
	Mutate(int) ICondition
	Hash() string
}

func ConditionToJson(c ICondition) string {
	conditionAttrs := make(map[string]interface{})
	conditionJsonNoType, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(conditionJsonNoType, &conditionAttrs)
	conditionAttrs["type"] = c.ConditionType()

	conditionJson, err := json.MarshalIndent(conditionAttrs, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(conditionJson)
}

func ConditionFromJson(conditionJson string) ICondition {
	var condition ICondition
	json.Unmarshal([]byte(conditionJson), condition)
	return condition
}

// TODO: Create a hash function for conditions
// TODO: Create a struct of conditions for pruning the conditions when they are created
