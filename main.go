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
	t "branch_learning/strategy_trainer"
	"flag"
	"fmt"
	"os"
)

var isTrain bool
var trainConfiguration *configuration.TrainConfiguration
var testConfiguration *configuration.TestConfiguration
var logConfiguration *configuration.LogConfiguration

func init() {
	var configFile string
	flag.StringVar(&configFile, "file", "", "Yaml Configuration file")
	flag.Parse()

	command := flag.Arg(0)

	switch command {
	case "train":
		isTrain = true
		trainConfiguration = configuration.ParseTrainConfiguration(configFile)
		logConfiguration = &trainConfiguration.LogConf
	case "test":
		isTrain = false
		testConfiguration = configuration.ParseTestConfiguration(configFile)
		logConfiguration = &testConfiguration.LogConf
	default:
		fmt.Println("Wrong command type")
		os.Exit(1)
	}
}

func main() {
	l.InitLoggers(*logConfiguration)
	logger := l.CreateLogger()

	if isTrain {
		logger.Info.Printf(
			"Starting To Train\nEpochs: %d\nGeneration Size: %d\nData: %s\n\nLogs Information\nLogs File: %s\nStrategies File: %s\nResults File: %s\n\n",
			trainConfiguration.EvolutionConf.Epochs,
			trainConfiguration.EvolutionConf.GenerationSize,
			trainConfiguration.DataPath,
			trainConfiguration.LogConf.LogsFile,
			trainConfiguration.LogConf.StrategiesFile,
			trainConfiguration.LogConf.ResultsFile,
		)
		trainer := t.CreateStrategyTrainer(trainConfiguration)
		data := candlestream.GetStreamsFromPath(trainConfiguration.DataPath)
		trainer.Train(data)
	} else {
		logger.Info.Printf(
			"Testing Strategy\nStrategy %s\nData: %s\n",
			testConfiguration.Strategy,
			testConfiguration.DataPath,
		)
	}

}
