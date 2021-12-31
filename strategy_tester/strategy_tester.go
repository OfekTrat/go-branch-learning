package tester

import (
	args_parser "branch_learning/args_parser"
	bt "branch_learning/backtester"
	candle_stream "branch_learning/candle_stream"
	"branch_learning/output"
	"branch_learning/parser"
	"fmt"
	"os"
)

func TestStrategy(configuration *args_parser.Configuration) {
	validateStrategyFile(configuration.StrategyFile)
	fileData, err := os.ReadFile(configuration.StrategyFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	strategy := parser.ParseStrategy(string(fileData))
	backtester := bt.CreateBackTester(strategy)
	stream := candle_stream.GetStreamsFromPath(configuration.DataPath)

	for _, s := range stream {
		backtester.Test(s)
	}
	output.PrintScore(backtester)
}

func validateStrategyFile(filename string) {
	if filename == "" {
		fmt.Println("Please input a strategy file")
		os.Exit(1)
	}
}
