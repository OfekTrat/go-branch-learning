package broker

type Broker struct {
	exitStopLossTree   *exitNode
	exitTakeProfitTree *exitNode
	results            map[int]*accountStats
	orders             map[int]map[int]bool // Bool specifies if order has been closed
}

func CreateBroker() *Broker {
	return &Broker{nil, nil, make(map[int]*accountStats), make(map[int]map[int]bool)}
}

func (broker *Broker) AddOrder(ord Order) {
	broker.orders[ord.StrategyId()][ord.Time()] = false
	if broker.exitStopLossTree == nil || broker.exitTakeProfitTree == nil {
		broker.exitStopLossTree = createExitNode(ord.stopLoss, ord)
		broker.exitTakeProfitTree = createExitNode(ord.takeProfit, ord)
	} else {
		broker.exitStopLossTree.Add(ord.stopLoss, ord)
		broker.exitTakeProfitTree.Add(ord.takeProfit, ord)
	}
}

func (broker *Broker) ScanOrders(lowPrice, highPrice float32) {
	broker.updateStopLossExits(lowPrice)
	broker.updateTakeProfitExits(highPrice)
}

func (broker *Broker) updateStopLossExits(lowPrice float32) {
	newStopLoss, orders := broker.exitStopLossTree.GetStopLossExits(lowPrice)
	broker.exitStopLossTree = newStopLoss

	for _, ord := range orders {
		if !broker.isClosed(ord) {
			broker.closeOrder(ord, false)
		}
	}

}

func (broker *Broker) updateTakeProfitExits(highPrice float32) {
	newHead, orders := broker.exitTakeProfitTree.GetTakeProfitExits(highPrice)
	broker.exitTakeProfitTree = newHead

	for _, ord := range orders {
		if !broker.isClosed(ord) {
			broker.closeOrder(ord, true)
		}
	}
}

func (broker *Broker) ScanResults() map[int]*accountStats {
	return broker.results
}

func (broker *Broker) isClosed(ord Order) bool {
	return broker.orders[ord.StrategyId()][ord.Time()]
}

func (broker *Broker) closeOrder(ord Order, isWin bool) {
	broker.orders[ord.StrategyId()][ord.Time()] = true
	if isWin {
		broker.results[ord.StrategyId()].AddLoss()
	} else {
		broker.results[ord.StrategyId()].AddWin()
	}
}

// TODO: Integrate broker in evolving.
