package main

import (
	"branch_learning/evolutioner"
	log_init "branch_learning/log_initializer"
	"branch_learning/parser"
	tester "branch_learning/strategy_tester"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	configuration := parser.InitArgs()

	switch configuration.CallType {
	case "test": // TODO: Create a mechanism for testing strategies (load strategy, and test it)
		tester.TestStrategy(configuration)
	default:
		log_init.LogInitialize(configuration.OutputConfig)
		evolutioner.Evolve(configuration.DataPath, configuration.EvolutionConfig, configuration.OutputConfig)
	}
}
