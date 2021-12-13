package evolutioner

import (
	candle_stream "branch_learning/candle_stream"
	"branch_learning/utils/random"
)

func Evolve(data *candle_stream.CandleStream, config *EvolutionConfig) {
	generation := random.CreateRandomGeneration(config.GenerationSize, &config.RandomStrategyConfig) // Creates Random Generation

	for i := 0; i < config.NumEvolutions; i++ {
		scores := backtestGeneration(data, generation)
		chances := calcChances(scores)
		generation = createNextGeneration(chances, generation, config)
	}
}
