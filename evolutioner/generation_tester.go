package evolutioner

import (
	b "branch_learning/broker"
	c "branch_learning/candle"
	cs "branch_learning/candle_stream"
	st "branch_learning/strategy"
	"sync"
)

func CheckStrategiesOnData(strategies []*st.Strategy, stream *cs.CandleStream) *b.Broker {

}

func checkStrategiesOnSlice(strategies []*st.Strategy, slicedStream *cs.CandleStream, broker *b.Broker) {
	var wg sync.WaitGroup
	nWorkers := 4
	sliceSize := len(strategies) / nWorkers

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			checkStrategiesWorker(strategies[i*sliceSize:(i+1)*sliceSize], slicedStream, broker)
		}(i)
	}
	wg.Wait()
}

func checkStrategiesWorker(strategies []*st.Strategy, stream *cs.CandleStream, broker *b.Broker) {
	for _, strategy := range strategies {
		checkStrategyOnSlice(strategy, stream, broker)
	}
}

func checkStrategyOnSlice(strategy *st.Strategy, stream *cs.CandleStream, broker *b.Broker) {
	if strategy.MeetsConditions(stream) {
		order := createOrder(strategy, stream.Get(stream.Length()-1))
		broker.AddOrder(order)
	}
}

func createOrder(strategy *st.Strategy, candle c.Candle) b.Order {
	mts := int(candle.Get("mts"))
	close := candle.Get("close")
	stopLoss := (100 - strategy.StopLoss()) * close
	takeProfit := (100 + strategy.TakeProfit()) * close
	return b.MakeOrder(mts, strategy.Id(), close, takeProfit, stopLoss)
}

// TODO: Implement the following things
// 1. Function that gets an order from single strategy on single slice - need to check for mutex stuff of broker.
// 2. Function that runs the function in section 1 in a multiprocessing mechanism and updates broker.
// Fix strategy ID mechanism
