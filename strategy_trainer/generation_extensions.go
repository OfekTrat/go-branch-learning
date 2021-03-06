package strategytrainer

import (
	mutator "branch_learning/alternators/mutator"
	reproducer "branch_learning/alternators/reproducer"
	st "branch_learning/strategy"
	random_util "branch_learning/utils/random"
	"math/rand"
)

func createRandomGeneration(generationId, generationSize int) *generation {
	strategies := make([]*st.Strategy, generationSize)

	for i := 0; i < generationSize; i++ {
		strategy := random_util.CreateRandomStrategy(i, generationId)
		strategies[i] = strategy

		logger.Strategies.Println(strategy.ToJsonString())
	}
	return newGeneration(generationId, strategies)
}

func createNextGenerationFromTestResults(generationId int, generationTestResults *generationTestResults) *generation {
	logger.Info.Printf("Start creating generation %d using test results\n", generationId)
	maxChance := generationTestResults.maxChance + 1
	newStrategies := make([]*st.Strategy, configuration.GenerationSize())
	pos := 0

	// Size of the pieces that make up new generation.
	oldSize := int(configuration.OldPercentage() * float32(configuration.GenerationSize()))
	mutateSize := int(configuration.MutatePercentage() * float32(configuration.GenerationSize()))
	reproducedSize := int(configuration.ReproducedPercentage() * float32(configuration.GenerationSize()))
	randomSize := configuration.GenerationSize() - oldSize - mutateSize - reproducedSize

	validateSizes(configuration.GenerationSize(), oldSize, mutateSize, reproducedSize, randomSize)

	oldStrategies := generationTestResults.GetNBestStrategy(oldSize)
	for i := 0; i < oldSize; i++ {
		strategy := oldStrategies[i]
		newStrategy := st.CreateStrategyFromOtherStrategy(pos, generationId, strategy)
		newStrategies[pos] = newStrategy
		pos++

		logger.Strategies.Println(newStrategy.ToJsonString())
	}
	for i := 0; i < mutateSize; i++ {
		chance := generateChance(maxChance)
		strategy := generationTestResults.GetStrategyByChance(chance)
		newStrategy := mutator.MutateStrategy(pos, generationId, strategy)
		newStrategies[pos] = newStrategy
		pos++

		logger.Strategies.Println(newStrategy.ToJsonString())
	}
	for i := 0; i < reproducedSize; i++ {
		chance1 := generateChance(maxChance)
		chance2 := generateChance(maxChance)

		strategy1 := generationTestResults.GetStrategyByChance(chance1)
		strategy2 := generationTestResults.GetStrategyByChance(chance2)
		newStrategy := reproducer.Reproduce(pos, generationId, strategy1, strategy2)
		newStrategies[pos] = newStrategy
		pos++

		logger.Strategies.Println(newStrategy.ToJsonString())
	}
	for i := 0; i < randomSize; i++ {
		strategy := random_util.CreateRandomStrategy(pos, generationId)
		newStrategies[pos] = strategy
		pos++

		logger.Strategies.Println(strategy.ToJsonString())
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
