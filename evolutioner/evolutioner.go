package evolutioner

import (
	bt "branch_learning/backtester"
	candle_stream "branch_learning/candle_stream"
	st "branch_learning/strategy"
	"branch_learning/utils/random"
	"fmt"
)

func Evolve(data *candle_stream.CandleStream, config *EvolutionConfig, printResults bool) {
	var scs []float64
	var chs chances
	var backtesters []*bt.BackTester
	generation := random.CreateRandomGeneration(config.GenerationSize, &config.RandomConfig)

	for i := 0; i < config.NumEvolutions; i++ {
		backtesters = createBacktesters(generation)
		scs = backtestGeneration(data, backtesters)
		chs = calcChances(scs)
		if printResults {
			printBestStrategy(backtesters, chs, scs)
		}
		generation = createNextGeneration(chs, generation, config)
	}
	printBestStrategy(backtesters, chs, scs)
}

func printBestStrategy(backtesters []*bt.BackTester, chs chances, scs []float64) {
	bestStrategyIndex := chs[len(chs)-1].strategyIndex
	bestScore := scs[bestStrategyIndex]
	backtesterOfBestStrategy := backtesters[bestStrategyIndex]
	bestStrategy := backtesterOfBestStrategy.Strategy()
	fmt.Println()
	fmt.Println("##################")
	fmt.Printf("Score: %f\n", bestScore)
	fmt.Printf("Strategy - TakeProfit: %v, StopLoss: %v, WindowSize: %v\n", bestStrategy.TakeProfit(),
		bestStrategy.StopLoss(), bestStrategy.WindowSize())
	fmt.Printf("Best Conditions: %v\n", bestStrategy.Conditions())
	fmt.Printf("Wins: %v, Losses: %v", backtesters[bestStrategyIndex].Stats().Wins(), backtesters[bestStrategyIndex].Stats().Losses())
	fmt.Println("##################")
	fmt.Println()
}

func createBacktesters(generation []*st.Strategy) []*bt.BackTester {
	backtesters := make([]*bt.BackTester, len(generation))

	for i := 0; i < len(generation); i++ {
		backtesters[i] = bt.CreateBackTester(generation[i])
	}
	return backtesters
}
