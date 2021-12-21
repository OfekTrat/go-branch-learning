package main

import (
	candle_stream "branch_learning/candle_stream"
	"branch_learning/evolutioner"
	"branch_learning/utils/random"
)

func main() {
	cs := candle_stream.LoadCandleStreamFromCsv("data/data.csv")
	randomConfig := random.RandomStrategyConfig{
		WindowMin:          10,
		WindowMax:          30,
		ExitMin:            0.5,
		ExitMax:            1.5,
		ConditionNumberMin: 1,
		ConditionNumberMax: 3,
	}

	configuration := evolutioner.EvolutionConfig{
		EvolutionLogFile:     "",
		GenerationSize:       100,
		NumEvolutions:        100,
		OldPercentage:        0.05,
		MutatePercentage:     0.3, // 30%
		ReproducedPercentage: 0.3, // 30%
		RandomPercentage:     0.35,

		RandomConfig:         randomConfig,
		ExitMutateMultiplier: 5,
		WindowSizeMultiplier: 5,
	}

	evolutioner.Evolve(cs, &configuration, true)
}
