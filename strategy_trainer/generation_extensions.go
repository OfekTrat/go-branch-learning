package strategytrainer

import (
	mutator "branch_learning/alternators/mutator"
	reproducer "branch_learning/alternators/reproducer"
	"branch_learning/configuration"
	st "branch_learning/strategy"
	random_util "branch_learning/utils/random"
	"math/rand"
)

func createRandomGeneration(generationId, generationSize int, randConf *configuration.RandomConfiguration) *generation {
	strategies := make([]*st.Strategy, generationSize)

	for i := 0; i < generationSize; i++ {
		strategies[i] = random_util.CreateRandomStrategy(i, generationId, randConf)
	}
	return newGeneration(generationId, strategies)
}

func createNextGenerationFromTestResults(generationId int, generationTestResults *generationTestResults, evoConf *configuration.EvolutionConfiguration, randConf *configuration.RandomConfiguration) *generation {
	logger.Info.Printf("Start creating generation %d using test results\n", generationId)
	maxChance := generationTestResults.maxChance + 1
	newStrategies := make([]*st.Strategy, evoConf.GenerationSize)
	pos := 0

	// Size of the pieces that make up new generation.
	oldSize := int(evoConf.OldPercentage * float32(evoConf.GenerationSize))
	mutateSize := int(evoConf.MutatePercentage * float32(evoConf.GenerationSize))
	reproducedSize := int(evoConf.ReproducedPercentage * float32(evoConf.GenerationSize))
	randomSize := evoConf.GenerationSize - oldSize - mutateSize - reproducedSize

	validateSizes(evoConf.GenerationSize, oldSize, mutateSize, reproducedSize, randomSize)

	for i := 0; i < oldSize; i++ {
		chance := generateChance(maxChance)
		strategy := generationTestResults.GetStrategyByChance(chance)
		newStrategies[pos] = st.CreateStrategyFromOtherStrategy(pos, generationId, strategy)
		pos++
	}
	for i := 0; i < mutateSize; i++ {
		chance := generateChance(maxChance)
		strategy := generationTestResults.GetStrategyByChance(chance)
		newStrategies[pos] = mutator.MutateStrategy(pos, generationId, strategy)
		pos++
	}
	for i := 0; i < reproducedSize; i++ {
		chance1 := generateChance(maxChance)
		chance2 := generateChance(maxChance)

		strategy1 := generationTestResults.GetStrategyByChance(chance1)
		strategy2 := generationTestResults.GetStrategyByChance(chance2)

		newStrategies[pos] = reproducer.Reproduce(pos, generationId, strategy1, strategy2)
		pos++
	}
	for i := 0; i < randomSize; i++ {
		newStrategies[pos] = random_util.CreateRandomStrategy(pos, generationId, randConf)
		pos++
	}
	logger.Info.Printf("Created new generation #%d", generationId)
	return newGeneration(generationId, newStrategies)
}

func validateSizes(wantedSize, oldSize, mutateSize, reproducedSize, randomSize int) {
	sizesSum := oldSize + mutateSize + reproducedSize + randomSize
	if sizesSum != wantedSize {
		logger.Error.Println("Calculated Sizes were wrong")
		logger.Error.Printf("Wanted Size: %d, Sum Sizes: %d\n", wantedSize, sizesSum)
	}
}

func generateChance(maxChance float64) float64 {
	return rand.Float64() * maxChance
}
