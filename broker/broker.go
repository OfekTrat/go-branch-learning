package broker

import (
	l "branch_learning/logger"
	s "branch_learning/strategy"
)

type Broker struct {
	exitStopLossTree   *exitNode
	exitTakeProfitTree *exitNode
	results            *AccountStats
	orders             map[int]bool // map[TimeOfOrder]IsClosed
}

var logger *l.Logger = l.CreateLogger()

const TIME_FORMAT = "2006-01-02 15:04:05"

func CreateBroker() *Broker {
	broker := &Broker{}
	broker.exitStopLossTree = nil
	broker.exitTakeProfitTree = nil
	broker.results = CreateEmptyAccountStats()
	broker.orders = make(map[int]bool)

	return broker
}

func (broker *Broker) ScanResults() *AccountStats {
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

func (broker *Broker) ScanOrders(lowPrice, highPrice float64) ([]Order, []Order) {
	ordersLost := broker.updateStopLossExits(lowPrice)
	ordersWon := broker.updateTakeProfitExits(highPrice)
	return ordersLost, ordersWon
}

func (broker *Broker) updateStopLossExits(lowPrice float64) []Order {
	newStopLoss, orders := broker.exitStopLossTree.GetStopLossExits(lowPrice)
	broker.exitStopLossTree = newStopLoss
	return orders
}

func (broker *Broker) updateTakeProfitExits(highPrice float64) []Order {
	newHead, orders := broker.exitTakeProfitTree.GetTakeProfitExits(highPrice)
	broker.exitTakeProfitTree = newHead
	return orders
}

func (broker *Broker) CloseWinOrders(closeTime int, strategy *s.Strategy, orders []Order) {
	broker.CloseOrders(closeTime, strategy, orders, true)
}

func (broker *Broker) CloseLossOrders(closeTime int, strategy *s.Strategy, orders []Order) {
	broker.CloseOrders(closeTime, strategy, orders, false)
}

func (broker *Broker) CloseOrders(closeTime int, strategy *s.Strategy, orders []Order, isWin bool) {
	for _, ord := range orders {
		if !broker.isOrderClosed(ord) {
			broker.closeOrder(ord, isWin)

			logger.LogOrder("%s,%d,%d,%d,%d,%f,%d\n", ord.ticker, ord.Time(), strategy.Generation(), strategy.Id(), 1, ord.stopLoss, closeTime)
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
