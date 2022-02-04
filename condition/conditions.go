package condition

import (
	candle_stream "branch_learning/candle_stream"
)

type Conditions struct {
	conditions map[string]ICondition
	keys       map[int]string
	length     int
}

func emptyConditions() *Conditions {
	conditions := new(Conditions)
	conditions.conditions = make(map[string]ICondition)
	conditions.keys = make(map[int]string)
	return conditions
}

func CreateConditions(conditionList []ICondition) *Conditions {
	conditions := emptyConditions()

	for _, cond := range conditionList {
		conditions.Add(cond)
	}
	return conditions
}

func (cs *Conditions) MeetsConditions(stream *candle_stream.CandleStream) bool {
	if cs.length == 0 {
		return false
	}

	for _, cond := range cs.conditions {
		if !cond.MeetsCondition(stream) {
			return false
		}
	}
	return true
}

func (cs *Conditions) Add(c ICondition) {
	if cs.conditions[c.Hash()] == nil {
		cs.keys[cs.length] = c.Hash()
		cs.conditions[c.Hash()] = c
		cs.length++
	} else {
		cond := cs.conditions[c.Hash()]
		if cond.IsOverriddenBy(c) {
			cs.conditions[c.Hash()] = c
		}
	}
}

func (cs *Conditions) SetInIndex(c ICondition, index int) {
	if index >= cs.length {
		panic("index is bigger than conditions length")
	}
	if cs.conditions[c.Hash()] == nil {
		currentHash := cs.keys[index]
		cs.keys[index] = c.Hash()
		delete(cs.conditions, currentHash)
		cs.conditions[c.Hash()] = c
		cs.length++
	} else {
		cs.conditions[c.Hash()] = c

		for index, hash := range cs.keys {
			if hash == c.Hash() {
				cs.keys[index] = c.Hash()
				break
			}
		}
	}

}

func (cs *Conditions) GetByHash(hash string) ICondition {
	return cs.conditions[hash]
}

func (cs *Conditions) GetByIndex(index int) ICondition {
	hash := cs.keys[index]
	return cs.conditions[hash]
}

func (cs *Conditions) RemoveByIndex(index int) {
	hashToRemove := cs.keys[index]

	for i := index; i < cs.length-1; i++ {
		cs.keys[index] = cs.keys[index+1]
	}
	delete(cs.keys, cs.length)
	delete(cs.conditions, hashToRemove)
	cs.length--
}

func (cs *Conditions) Clone() *Conditions {
	newConditions := emptyConditions()

	for _, cond := range cs.conditions {
		newConditions.Add(cond)
	}
	return newConditions
}

func (cs *Conditions) Length() int {
	return cs.length
}
