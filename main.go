package main

import (
	"branch_learning/args_parser"
	"branch_learning/evolutioner"
	log_init "branch_learning/log_initializer"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	configuration := args_parser.InitArgs()

	switch configuration.CallType {
	case "test": // TODO: Create a mechanism for testing strategies (load strategy, and test it)
		break
	default:
		log_init.LogInitialize(configuration.OutputConfig)
		evolutioner.Evolve(configuration.DataPath, configuration.EvolutionConfig, configuration.OutputConfig)
	}
}

// TODO: Accept folder of data and train on multiple data files.
// TODO: Speed up training (use goroutines - check if multiple cores are used)
// TODO: Create a test strategy option. This is for strategies that have been trained but not tested on the test data that i'll have.
