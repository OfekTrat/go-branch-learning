package condition

import (
	candle_stream "branch_learning/candle_stream"
)

type Conditions map[string]ICondition

func CreateConditions(conditionList []ICondition) *Conditions {
	conditions := Conditions{}

	for _, cond := range conditionList {
		conditions.Add(cond)
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
	if cs.Length() == 0 {
		return false
	}
	for _, cond := range *cs {
		if !cond.MeetsCondition(stream) {
			return false
		}
	}
	return true
}

func (cs *Conditions) ToList() []ICondition {
	i := 0
	conds := make([]ICondition, cs.Length())
	for _, cond := range *cs {
		conds[i] = cond
		i++
	}
	return conds
}

func (cs *Conditions) Length() int {
	return len((*cs))
}
