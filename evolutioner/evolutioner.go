package evolutioner

import (
	bt "branch_learning/backtester"
	candle_stream "branch_learning/candle_stream"
	"branch_learning/output"
	st "branch_learning/strategy"
	"branch_learning/utils/random"
	"fmt"
)

func Evolve(data_path string, config *EvolutionConfig, output_config *output.OutputConfig) {
	var scs []float64
	var chs chances
	var backtesters []*bt.BackTester
	data := candle_stream.GetStreamsFromPath(data_path)
	generation := random.CreateRandomGeneration(config.GenerationSize, &config.RandomConfig)

	for i := 0; i < config.NumEvolutions; i++ {
		backtesters = createBacktesters(generation)
		scs = backtestGeneration(data, backtesters)
		chs = calcChances(scs)
		printBestStrategy(backtesters, chs, scs, i, output_config)
		generation = createNextGeneration(chs, generation, config)
	}
	if output_config.PrintFrequency == "once" {
		output_config.PrintFrequency = "foreach"
		printBestStrategy(backtesters, chs, scs, config.NumEvolutions, output_config)
		output_config.PrintFrequency = "once"
	}
}

func printBestStrategy(backtesters []*bt.BackTester, chs chances, scs []float64, iteration int, output_config *output.OutputConfig) {
	fmt.Printf("\n####### %v\n", iteration)
	bestStrategyIndex := chs[len(chs)-1].strategyIndex
	backtesterOfBestStrategy := backtesters[bestStrategyIndex]
	bestStrategy := backtesterOfBestStrategy.Strategy()

	if output_config.PrintFrequency == "foreach" {
		output.PrintScore(backtesterOfBestStrategy)
	}
	if output_config.PrintBestStrategy && output_config.PrintFrequency == "foreach" {
		output.PrintStrategyConditions(bestStrategy)
	}
}

func createBacktesters(generation []*st.Strategy) []*bt.BackTester {
	backtesters := make([]*bt.BackTester, len(generation))

	for i := 0; i < len(generation); i++ {
		backtesters[i] = bt.CreateBackTester(generation[i])
	}
	return backtesters
}
