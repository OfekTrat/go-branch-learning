package strategytrainer

import (
	candlestream "branch_learning/candle_stream"
	"branch_learning/configuration"
	l "branch_learning/logger"
)

var logger *l.Logger = l.CreateLogger()

type StrategyTrainer struct {
	*configuration.EvolutionConfiguration
	*configuration.RandomConfiguration
}

func CreateStrategyTrainer(trainConf *configuration.TrainConfiguration) *StrategyTrainer {
	return &StrategyTrainer{
		EvolutionConfiguration: &trainConf.EvolutionConf,
		RandomConfiguration:    &trainConf.RandomConf,
	}
}

func (trainer *StrategyTrainer) Train(streams []*candlestream.CandleStream) {
	logger.Info.Println("Train - Started training")

	generation := createRandomGeneration(0, trainer.EvolutionConfiguration.GenerationSize, trainer.RandomConfiguration)

	for epoch := 1; epoch <= trainer.EvolutionConfiguration.Epochs; epoch++ {
		logger.Info.Printf("----------- START Generation %d -----------\n", epoch)

		testResults := generation.test(streams)
		generation = createNextGenerationFromTestResults(epoch, testResults, trainer.EvolutionConfiguration, trainer.RandomConfiguration)

		logger.Info.Printf("----------- END Generation %d -----------\n", epoch)
	}
}
