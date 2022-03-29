package strategytrainer

import (
	candlestream "branch_learning/candle_stream"
	c "branch_learning/configuration"
	l "branch_learning/logger"
)

var logger *l.Logger = l.CreateLogger()
var configuration *c.Configuration = c.GetConfiguration()

type StrategyTrainer struct {
	generationSize       int
	epochs               int
	oldPercentage        float32
	mutatePercentage     float32
	reproducedPercentage float32
	randomPercentage     float32

	windowMin          int
	windowMax          int
	stopLossMin        float64
	takeProfitMax      float64
	conditionNumberMin int
	conditionNumberMax int
}

func CreateStrategyTrainer() *StrategyTrainer {
	return &StrategyTrainer{
		generationSize:       configuration.GenerationSize(),
		epochs:               configuration.Epochs(),
		oldPercentage:        configuration.OldPercentage(),
		mutatePercentage:     configuration.MutatePercentage(),
		reproducedPercentage: configuration.ReproducedPercentage(),
		randomPercentage:     configuration.RandomPercentage(),
		windowMin:            configuration.WindowMin(),
		windowMax:            configuration.WindowMax(),
		stopLossMin:          configuration.StopLossMin(),
		takeProfitMax:        configuration.TakeProfitMax(),
		conditionNumberMin:   configuration.ConditionNumberMin(),
		conditionNumberMax:   configuration.ConditionNumberMax(),
	}
}

func (trainer *StrategyTrainer) Train(streams []*candlestream.CandleStream) {
	logger.Info.Println("Train - Started training")

	generation := createRandomGeneration(0, trainer.generationSize)

	for epoch := 0; epoch <= trainer.epochs-1; epoch++ {
		logger.Info.Printf("----------- START Generation %d -----------\n", epoch)

		testResults := generation.test(streams)
		logBestScore(testResults)
		if epoch != trainer.epochs-1 {
			generation = createNextGenerationFromTestResults(epoch+1, testResults)
		}

		logger.Info.Printf("----------- END Generation %d -----------\n", epoch)
	}
}

func logBestScore(results *generationTestResults) {
	maxChance := results.GetMaxChance()
	bestResults := results.tree.GetStrategyTesterByChance(maxChance + 1).Results()
	logger.Info.Printf("#### BEST SCORE: %f ####", bestResults.Score)
}
