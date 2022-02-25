package main

import (
	"branch_learning/evolutioner"
	"branch_learning/output"
	"branch_learning/parser"
	tester "branch_learning/strategy_tester"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	configuration := parser.InitArgs()

	switch configuration.CallType {
	case "test":
		tester.TestStrategy(configuration)
	default:
		output.LogInitialize(configuration.OutputConfig)
		evolutioner.Evolve(configuration.DataPath, configuration.EvolutionConfig, configuration.OutputConfig)
	}
}

// There is a crazy overfit, check how to demolish it.
// I have a problem in test, testing on trained data does not give the same result as the training.
// Make sure the train works well, because it does not match well the number of wanted orders
