package backtester

import (
	"branch_learning/exit"
	"fmt"
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

func (om *OrderManager) CheckExits(high, low float32) (int, int) { // (wins, losses)
	wins := 0
	losses := 0
	indicesToRemove := []int{}

	for i, e := range om.exits {
		isTake := e.IsTake(high)
		isLoss := e.IsStop(low)
		if isLoss {
			fmt.Println("Loss", e, high, low)
			indicesToRemove = append(indicesToRemove, i)
			losses++
		} else if isTake {
			fmt.Println("Win", e, high, low)
			indicesToRemove = append(indicesToRemove, i)
			wins++
		}
	}
	om.cleanExitsByIndices(indicesToRemove)
	return wins, losses
}

func (om *OrderManager) cleanExitsByIndices(indicesToRemove []int) {
	for i := len(indicesToRemove) - 1; i >= 0; i-- {
		indexToRemove := indicesToRemove[i]
		om.exits = append(om.exits[:indexToRemove], om.exits[indexToRemove+1:]...)
	}
}
