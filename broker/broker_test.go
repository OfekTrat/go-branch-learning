package broker

import (
	c "branch_learning/condition"
	s "branch_learning/strategy"
	"testing"
)

func getSampleBroker() *Broker {
	ord1 := MakeOrder("test", 0, 1, 6, 1)
	ord2 := MakeOrder("test", 1, 2, 2, 2)
	ord3 := MakeOrder("test", 2, 3, 3, 3)
	ord4 := MakeOrder("test", 3, 4, 4, 4)
	ord5 := MakeOrder("test", 4, 5, 5, 5)
	ord6 := MakeOrder("test", 5, 6, 6, 6)
	ord7 := MakeOrder("test", 6, 7, 7, 7)

	broker := CreateBroker()
	broker.AddOrder(ord4)
	broker.AddOrder(ord2)
	broker.AddOrder(ord1)
	broker.AddOrder(ord3)
	broker.AddOrder(ord6)
	broker.AddOrder(ord5)
	broker.AddOrder(ord7)

	return broker
}

func TestCreateBrokerAndInitialization(t *testing.T) {
	ord1 := MakeOrder("test", 0, 1, 1.5, 0.5)
	ord2 := MakeOrder("test", 0, 2, 2.5, 1.5)
	ord3 := MakeOrder("test", 0, 3, 3.5, 2.5)
	ord4 := MakeOrder("test", 0, 4, 4.5, 3.5)
	broker := CreateBroker()
	broker.AddOrder(ord3)
	broker.AddOrder(ord4)
	broker.AddOrder(ord2)
	broker.AddOrder(ord1)

	if broker.exitStopLossTree.value != ord3.StopLoss() || len(broker.exitStopLossTree.orders) != 1 ||
		broker.exitStopLossTree.lower == nil || broker.exitStopLossTree.lower.value != ord2.StopLoss() ||
		broker.exitStopLossTree.lower.lower == nil || broker.exitStopLossTree.lower.lower.value != ord1.StopLoss() ||
		broker.exitStopLossTree.upper == nil || broker.exitStopLossTree.upper.value != ord4.StopLoss() {

		t.Error("Assertion Error: Build of stop loss tree failed.")
	}

	if broker.exitTakeProfitTree.value != ord3.TakeProfit() || len(broker.exitTakeProfitTree.orders) != 1 ||
		broker.exitTakeProfitTree.lower == nil || broker.exitTakeProfitTree.lower.value != ord2.TakeProfit() ||
		broker.exitTakeProfitTree.lower.lower == nil || broker.exitTakeProfitTree.lower.lower.value != ord1.TakeProfit() ||
		broker.exitTakeProfitTree.upper == nil || broker.exitTakeProfitTree.upper.value != ord4.TakeProfit() {

		t.Error("Assertion Error: Build of stop loss tree failed.")
	}
}

func TestScanningOrders(t *testing.T) {
	conditions := &c.Conditions{}
	strategy := s.CreateStrategy(0, 0, 10, 1.1, 1.1, conditions)
	broker := getSampleBroker()
	ordersLost, ordersWon := broker.ScanOrders(5, 5.5)

	broker.CloseLossOrders(0, strategy, ordersLost)
	broker.CloseWinOrders(0, strategy, ordersWon)

	results := broker.ScanResults()
	if results.Losses() != 3 || results.Wins() != 3 {
		t.Error("AssertionError: wrong number of wins or losses")
	}
}

func TestClosingOrders(t *testing.T) {
	ord := MakeOrder("test", 3, 4, 4, 4)
	broker := getSampleBroker()

	if broker.isOrderClosed(ord) {
		t.Error("AssertionError: The order should no be closed")
	}
	broker.closeOrder(ord, false)
	if !broker.isOrderClosed(ord) {
		t.Error("AssertionError: The order should be closed")
	}
	results := broker.ScanResults()

	if results.Wins() != 0 || results.Losses() != 1 {
		t.Error("AssertionError: did not add close retults")
	}
}
