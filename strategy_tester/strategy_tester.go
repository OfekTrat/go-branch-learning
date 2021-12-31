package tester

import (
	args_parser "branch_learning/args_parser"
	bt "branch_learning/backtester"
	candle_stream "branch_learning/candle_stream"
	"branch_learning/output"
	"fmt"
	"os"
)

func TestStrategy(configuration *args_parser.Configuration) {
	fmt.Println(configuration.StrategyFile)
	if configuration.StrategyFile == "" {
		fmt.Println("Strategy File not specified")
		os.Exit(1)
	}
	strategy := loadStrategyFromFile(configuration.StrategyFile)
	backtester := bt.CreateBackTester(strategy)
	stream := candle_stream.GetStreamsFromPath(configuration.DataPath)

	for _, s := range stream {
		backtester.Test(s)
	}
	output.PrintScore(backtester)
}
