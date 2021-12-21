package main

import (
	"branch_learning/args_parser"
	candle_stream "branch_learning/candle_stream"
	"branch_learning/evolutioner"
	log_init "branch_learning/log_initializer"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	configuration := args_parser.InitArgs()
	cs := candle_stream.LoadCandleStreamFromCsv(configuration.DataFile)

	switch configuration.CallType {
	case "test": // TODO: Create a mechanism for testing strategies (load strategy, and test it)
		break
	default:
		log_init.LogInitialize(configuration.OutputConfig)
		evolutioner.Evolve(cs, configuration.EvolutionConfig, configuration.OutputConfig)
	}
}
