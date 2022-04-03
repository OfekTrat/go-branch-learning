package broker

type exitNode struct {
	value  float64
	orders []Order
	lower  *exitNode
	upper  *exitNode
}

func createExitNode(value float64, ord Order) *exitNode {
	return &exitNode{value, []Order{ord}, nil, nil}
}

func (en *exitNode) Add(value float64, ord Order) {
	pos := en

	for pos != nil {
		if pos.value == value {
			pos.orders = append(pos.orders, ord)
			return
		} else {
			if pos.value <= value {
				if pos.upper == nil {
					pos.upper = createExitNode(value, ord)
					return
				} else {
					pos = pos.upper
				}
			} else {
				if pos.lower == nil {
					pos.lower = createExitNode(value, ord)
					return
				} else {
					pos = pos.lower
				}
			}
		}
	}
}

// TODO: Implement Tree Pruning. So that when GetStopLossExits, or GetTakeProfitExits are called, the tree is pruned to the only left nodes.
func (en *exitNode) GetStopLossExits(value float64) (*exitNode, []Order) {
	orders := []Order{}

	if en == nil {
		return nil, orders
	} else if en.value >= value {
		var lowerNode *exitNode = nil
		orders = append(orders, en.orders...)

		if en.upper != nil {
			orders = append(orders, en.upper.getAndRemoveAllOrders()...)
		}
		if en.lower != nil {
			var closedOrders []Order
			lowerNode, closedOrders = en.lower.GetStopLossExits(value)
			orders = append(orders, closedOrders...)
		}
		en.orders = []Order{} //Clean Up

		return lowerNode, orders
	} else {
		newUpper, orders := en.upper.GetStopLossExits(value)
		en.upper = newUpper
		return en, orders
	}
}

func (en *exitNode) GetTakeProfitExits(value float64) (*exitNode, []Order) {
	orders := []Order{}

	if en == nil {
		return nil, orders
	} else if en.value < value {
		var upperNode *exitNode = nil
		orders = append(orders, en.orders...)

		if en.lower != nil {
			orders = append(orders, en.lower.getAndRemoveAllOrders()...) // Also cleans all orders
		}
		if en.upper != nil {
			var closedOrders []Order
			upperNode, closedOrders = en.upper.GetTakeProfitExits(value)
			orders = append(orders, closedOrders...)
		}
		en.orders = []Order{} // Clean Up
		return upperNode, orders
	} else {
		newLower, orders := en.lower.GetTakeProfitExits(value)
		en.lower = newLower
		return en, orders
	}
}

func (en *exitNode) getAndRemoveAllOrders() []Order {
	orders := make([]Order, len(en.orders))
	copy(orders, en.orders)
	en.orders = []Order{}

	if en.lower != nil {
		orders = append(orders, en.lower.getAndRemoveAllOrders()...)
	}
	if en.upper != nil {
		orders = append(orders, en.upper.getAndRemoveAllOrders()...)
	}
	return orders
}
