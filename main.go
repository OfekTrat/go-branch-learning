package main

// import (
// 	"branch_learning/evolutioner"
// 	"branch_learning/initialization"
// 	"branch_learning/output"
// 	"branch_learning/parser"
// 	tester "branch_learning/strategy_tester"
// 	"math/rand"
// 	"time"
// )

import (
	candlestream "branch_learning/candle_stream"
	c "branch_learning/configuration"
	l "branch_learning/logger"
	s "branch_learning/strategy"
	st "branch_learning/strategy_tester"
	t "branch_learning/strategy_trainer"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var isTrain bool
var configuration *c.Configuration = c.GetConfiguration()

func init() {
	command := flag.Arg(0)
	fmt.Println(configuration)
	switch command {
	case "train":
		isTrain = true
	case "test":
		isTrain = false
	default:
		fmt.Println("Wrong command type")
		os.Exit(1)
	}
	rand.Seed(time.Now().UnixNano())
}

func main() {
	logger := l.CreateLogger()
	dataPath := configuration.Data()

	if isTrain {
		logger.Info.Printf(
			"Starting To Train\nEpochs: %d\nGeneration Size: %d\nData: %s\n\n",
			configuration.Epochs(),
			configuration.GenerationSize(),
			dataPath,
		)
		trainer := t.CreateStrategyTrainer()
		data := candlestream.GetStreamsFromPath(dataPath)
		trainer.Train(data)
	} else {
		strategyPath := configuration.Strategy()
		strategy := s.CreateStrategyFromFile(strategyPath)
		data := candlestream.GetStreamsFromPath(dataPath)
		tester := st.NewStrategyTester(strategy)
		tester.Test(data)
		logger.Info.Printf(
			"Testing Strategy\nStrategy %s\nData: %s\n",
			configuration.Strategy(),
			dataPath,
		)
	}
	l.ZipLogs(configuration.Output())
}
