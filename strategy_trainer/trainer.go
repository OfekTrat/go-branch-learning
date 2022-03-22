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

	for epoch := 0; epoch <= trainer.EvolutionConfiguration.Epochs-1; epoch++ {
		logger.Info.Printf("----------- START Generation %d -----------\n", epoch)

		testResults := generation.test(streams)
		logBestScore(testResults)
		if epoch != trainer.EvolutionConfiguration.Epochs-1 {
			generation = createNextGenerationFromTestResults(epoch+1, testResults, trainer.EvolutionConfiguration, trainer.RandomConfiguration)
		}

		logger.Info.Printf("----------- END Generation %d -----------\n", epoch)
	}
}

func logBestScore(results *generationTestResults) {
	maxChance := results.GetMaxChance()
	bestResults := results.tree.GetStrategyTesterByChance(maxChance + 1).Results()
	logger.Info.Printf("#### BEST SCORE: %f ####", bestResults.Score)
}
