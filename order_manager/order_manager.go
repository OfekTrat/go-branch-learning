package ordermanager

import (
	"branch_learning/exit"
)

type OrderManager struct {
	exits []exit.Exit
}

func (om *OrderManager) Exits() []exit.Exit {
	return om.exits
}

func (om *OrderManager) AddExit(e exit.Exit) {
	om.exits = append(om.exits, e)
}

func (om *OrderManager) CheckExits(price float32) (int, int) { // (wins, losses)
	wins := 0
	losses := 0
	indexesToRemove := []int{}

	for i, e := range om.exits {
		isTake := e.Take(price)
		isLoss := e.Stop(price)

		if isTake && !isLoss {
			wins++

			continue
		}
		if isLoss && !isTake {
			indexesToRemove = append(indexesToRemove, i)
			losses++
			continue
		}
	}
	for i := len(indexesToRemove) - 1; i >= 0; i-- {
		indexToRemove := indexesToRemove[i]
		om.exits = append(om.exits[:indexToRemove], om.exits[indexToRemove+1:]...)
	}
	return wins, losses
}
