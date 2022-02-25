package tester

import (
	bt "branch_learning/backtester"
	candle_stream "branch_learning/candle_stream"
	"branch_learning/output"
	"branch_learning/parser"
	"fmt"
	"os"
)

func TestStrategy(configuration *parser.Configuration) {
	validateStrategyFile(configuration.StrategyFile)
	fileData, err := os.ReadFile(configuration.StrategyFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	strategy := parser.ParseStrategy(fileData)
	backtester := bt.CreateBackTester(strategy)
	streams := candle_stream.GetStreamsFromPath(configuration.DataPath)

	for _, s := range streams {
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
