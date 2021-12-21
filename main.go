package main

import (
	candle_stream "branch_learning/candle_stream"
	"branch_learning/evolutioner"
	log_init "branch_learning/log_initializer"
	"branch_learning/output"
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
		GenerationSize:       100,
		NumEvolutions:        10,
		OldPercentage:        0.1,
		MutatePercentage:     0.3, // 30%
		ReproducedPercentage: 0.3, // 30%
		RandomPercentage:     0.3,

		RandomConfig:         randomConfig,
		ExitMutateMultiplier: 5,
		WindowSizeMultiplier: 5,
	}

	outputConfig := output.OutputConfig{
		LogFile:           "",
		PrintFrequency:    "foreach",
		PrintBestStrategy: false,
	}
	log_init.LogInitialize(&outputConfig)
	evolutioner.Evolve(cs, &configuration, &outputConfig)
}
