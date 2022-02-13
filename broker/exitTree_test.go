package broker

import (
	"fmt"
	"testing"
)

func TestCreateTree_SingleNodeSingleOrder(t *testing.T) {
	ord := MakeOrder(0, 0, 10, 11, 9)
	node := createExitNode(ord.stopLoss, ord)

	if node.value != ord.stopLoss {
		t.Error("AssertionError: nodes value is not iniaizlied correctly")
	}
	if len(node.orders) != 1 {
		t.Error("AssertionError: Wrong orders length")
	}
	if node.lower != nil {
		t.Error("AssertionError: Wrong lower node")
	}
	if node.upper != nil {
		t.Error("AssertionError: Wrong lower node")
	}
}

func TestCreateTree_MultipleNodes(t *testing.T) {
	ord1 := MakeOrder(0, 0, 10, 11, 9)
	ord2 := MakeOrder(0, 1, 10, 11, 8)
	ord3 := MakeOrder(0, 2, 10, 11, 10)
	head := createExitNode(ord1.stopLoss, ord1)
	head.Add(ord2.stopLoss, ord2)
	head.Add(ord3.stopLoss, ord3)

	if head == nil || head.value != ord1.stopLoss {
		t.Error("AssertionError: Creation Failed")
	}
	if head.lower == nil || head.lower.value != ord2.stopLoss {
		t.Error("AssertionError: Addition of lower failed")
	}
	if head.upper == nil || head.upper.value != ord3.stopLoss {
		t.Error("AssertionError: Addition of upper failed")
	}
}

func TestCreateTree_MultipleNodesMultipleOrders(t *testing.T) {
	ord1 := MakeOrder(0, 0, 10, 11, 9)
	ord2 := MakeOrder(0, 1, 10, 11, 10)
	ord3 := MakeOrder(0, 2, 10, 11, 10)

	head := createExitNode(ord1.stopLoss, ord1)
	head.Add(ord2.stopLoss, ord2)
	head.Add(ord3.stopLoss, ord3)

	if head == nil || head.value != ord1.stopLoss {
		t.Error("AssertionError: creation of node has failed")
	}
	if head.lower != nil {
		t.Error("AssertionError: Lower shouldn't have been there")
	}
	if head.upper == nil || head.upper.value != ord2.stopLoss || len(head.upper.orders) != 2 {
		t.Error("AssertionError: Addition of upper values has failed")
	}
}

func TestExitTree_GetStopLossOrders(t *testing.T) {
	ord1 := MakeOrder(0, 0, 10, 11, 9)
	ord2 := MakeOrder(0, 1, 10, 11, 8)
	ord5 := MakeOrder(1, 1, 11, 1, 7.5)
	ord3 := MakeOrder(0, 2, 10, 11, 10)
	ord4 := MakeOrder(1, 3, 12, 11, 9.5)
	head := createExitNode(ord1.StopLoss(), ord1)
	head.Add(ord2.StopLoss(), ord2)
	head.Add(ord3.StopLoss(), ord3)
	head.Add(ord4.StopLoss(), ord4)
	head.Add(ord5.StopLoss(), ord5)

	newHead, orders := head.GetStopLossExits(8.5)

	if len(orders) != 3 {
		t.Error("AssertionError: Wrong orders lenght")
	}
	for _, ord := range orders {
		if !(ord.Equals(ord1) || ord.Equals(ord3) || ord.Equals(ord4)) {
			t.Error("AssertionError: not right orders were retreived")
		}
	}
	fmt.Println(newHead)
	if newHead == nil || len(newHead.orders) != 1 || newHead.lower == nil || len(newHead.lower.orders) != 1 ||
		!newHead.orders[0].Equals(ord2) || !newHead.lower.orders[0].Equals(ord5) {
		t.Error("AssertionError: did not return the right tree")
	}
}

func TestExitTree_GetTakeProfitOrders(t *testing.T) {
	ord1 := MakeOrder(0, 0, 10, 11, 9)
	ord2 := MakeOrder(0, 1, 10, 12, 8)
	ord3 := MakeOrder(0, 2, 10, 10, 10)
	ord4 := MakeOrder(1, 2, 10, 10.4, 9)
	head := createExitNode(ord1.TakeProfit(), ord1)
	head.Add(ord2.TakeProfit(), ord2)
	head.Add(ord3.TakeProfit(), ord3)
	head.Add(ord4.TakeProfit(), ord4)

	newHead, orders := head.GetTakeProfitExits(11)

	if len(orders) != 2 {
		t.Error("AssertionError: wrong order length")
	}
	for _, ord := range orders {
		if !(ord.Equals(ord3)) && !(ord.Equals(ord4)) {
			t.Error("AssertionError: wrong orders")
		}
	}

	if newHead == nil || len(newHead.orders) != 1 || !newHead.orders[0].Equals(ord1) ||
		newHead.upper == nil || len(newHead.orders) != 1 || !newHead.upper.orders[0].Equals(ord2) {
		t.Error("AssertionError: Wrong new tree")
	}
}
