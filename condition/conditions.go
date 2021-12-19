package condition

import (
	candle_stream "branch_learning/candle_stream"
)

type Conditions map[string]ICondition

func CreateConditions(cs []ICondition) *Conditions {
	conditions := Conditions{}

	for _, cond := range cs {
		if conditions[cond.Hash()] == nil {
			conditions[cond.Hash()] = cond
		}
	}
	return &conditions
}

func (cs *Conditions) Add(c ICondition) {
	if (*cs)[c.Hash()] == nil {
		(*cs)[c.Hash()] = c
	}
}

func (cs *Conditions) AddMultiple(conds []ICondition) {
	for _, cond := range conds {
		cs.Add(cond)
	}
}

func (cs *Conditions) Clone() *Conditions {
	newConditions := Conditions{}

	for _, cond := range *cs {
		(&newConditions).Add(cond)
	}
	return &newConditions
}

func (cs *Conditions) MeetsConditions(stream *candle_stream.CandleStream) bool {
	for _, cond := range *cs {
		if !cond.MeetsCondition(stream) {
			return false
		}
	}
	return true
}

func (cs *Conditions) ToList() []ICondition {
	conds := make([]ICondition, len(*cs))
	for _, cond := range *cs {
		conds = append(conds, cond)
	}
	return conds
}

func (cs *Conditions) Length() int {
	return len((*cs))
}
