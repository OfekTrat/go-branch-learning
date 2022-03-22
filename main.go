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
	"branch_learning/configuration"
	l "branch_learning/logger"
	s "branch_learning/strategy"
	st "branch_learning/strategy_tester"
	t "branch_learning/strategy_trainer"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

var isTrain bool
var trainConfiguration *configuration.TrainConfiguration
var testConfiguration *configuration.TestConfiguration
var shouldLogOrders bool

func init() {
	var configFile string
	flag.StringVar(&configFile, "file", "", "Yaml Configuration file")
	flag.Parse()

	command := flag.Arg(0)

	switch command {
	case "train":
		isTrain = true
		trainConfiguration = configuration.ParseTrainConfiguration(configFile)
		shouldLogOrders = trainConfiguration.ShouldLogOrders
	case "test":
		isTrain = false
		testConfiguration = configuration.ParseTestConfiguration(configFile)
		shouldLogOrders = testConfiguration.ShouldLogOrders
	default:
		fmt.Println("Wrong command type")
		os.Exit(1)
	}

	if shouldLogOrders {
		l.EnableOrdersLogs()
	}
}

func main() {
	logger := l.CreateLogger()

	if isTrain {
		logger.Info.Printf(
			"Starting To Train\nEpochs: %d\nGeneration Size: %d\nData: %s\n\n",
			trainConfiguration.EvolutionConf.Epochs,
			trainConfiguration.EvolutionConf.GenerationSize,
			trainConfiguration.DataPath,
		)
		trainer := t.CreateStrategyTrainer(trainConfiguration)
		data := candlestream.GetStreamsFromPath(trainConfiguration.DataPath)
		trainer.Train(data)
	} else {
		strategy := s.CreateStrategyFromFile(testConfiguration.Strategy)
		data := candlestream.GetStreamsFromPath(testConfiguration.DataPath)
		tester := st.NewStrategyTester(strategy)
		tester.Test(data)
		logger.Info.Printf(
			"Testing Strategy\nStrategy %s\nData: %s\n",
			testConfiguration.Strategy,
			testConfiguration.DataPath,
		)
	}

	filename := createTimeFilename()
	l.ZipLogs(filename)
}

func createTimeFilename() string {
	now := time.Now()
	filename := strconv.FormatInt(int64(now.Year()), 10) +
		strconv.FormatInt(int64(now.Month()), 10) +
		strconv.FormatInt(int64(now.Day()), 10) +
		"-" +
		strconv.FormatInt(int64(now.Hour()), 10) +
		strconv.FormatInt(int64(now.Minute()), 10)

	return filename + ".zip"
}
