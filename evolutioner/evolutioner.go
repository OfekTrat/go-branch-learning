package evolutioner

import (
	candle_stream "branch_learning/candle_stream"
	st "branch_learning/strategy"
	"branch_learning/utils/random"
	"fmt"
)

func Evolve(data *candle_stream.CandleStream, config *EvolutionConfig) {
	generation := random.CreateRandomGeneration(config.GenerationSize, &config.RandomConfig)

	for i := 0; i < config.NumEvolutions; i++ {
		scs := backtestGeneration(data, generation)
		chs := calcChances(scs)
		printBestStrategy(generation, chs, scs)
		generation = createNextGeneration(chs, generation, config)
	}
}

func printBestStrategy(generation []*st.Strategy, chs chances, scs []float64) {
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
