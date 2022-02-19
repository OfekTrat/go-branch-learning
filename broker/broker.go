package broker

type Broker struct {
	exitStopLossTree   *exitNode
	exitTakeProfitTree *exitNode
	results            *accountStats
	orders             map[int]bool // map[TimeOfOrder]IsClosed
}

func CreateBroker() *Broker {
	broker := &Broker{}
	broker.exitStopLossTree = nil
	broker.exitTakeProfitTree = nil
	broker.results = AccountStats()
	broker.orders = make(map[int]bool)

	return broker
}

func (broker *Broker) ScanResults() *accountStats {
	return broker.results
}

func (broker *Broker) AddOrder(ord Order) {
	broker.orders[ord.Time()] = false

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
		if !broker.isOrderClosed(ord) {
			broker.closeOrder(ord, false)
		}
	}

}

func (broker *Broker) updateTakeProfitExits(highPrice float32) {
	newHead, orders := broker.exitTakeProfitTree.GetTakeProfitExits(highPrice)
	broker.exitTakeProfitTree = newHead

	for _, ord := range orders {
		if !broker.isOrderClosed(ord) {
			broker.closeOrder(ord, true)
		}
	}
}

func (broker *Broker) isOrderClosed(ord Order) bool {
	return broker.orders[ord.Time()]
}

func (broker *Broker) closeOrder(ord Order, isWin bool) {
	broker.orders[ord.Time()] = true

	if isWin {
		broker.results.AddWin()
	} else {
		broker.results.AddLoss()
	}
}
