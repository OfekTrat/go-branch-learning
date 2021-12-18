package evolutioner

import (
	"branch_learning/alternators/mutator"
	"branch_learning/alternators/reproducer"
	st "branch_learning/strategy"
	"branch_learning/utils/random"
	"math/rand"
)

func createNextGeneration(chs chances, lastGeneration []*st.Strategy, config *EvolutionConfig) []*st.Strategy {
	var i int = 0
	var strategyIndex int
	var strategyIndex2 int
	var randNumber float64
	var randNumber2 float64
	maxScore := chs[len(chs)-1].score

	nOld := int(config.OldPercentage * float32(config.GenerationSize))
	nMutated := int(config.MutatePercentage * float32(config.GenerationSize))
	nReproduced := int(config.ReproducedPercentage * float32(config.GenerationSize))
	nNew := config.GenerationSize - nOld - nMutated - nReproduced
	nextGeneration := make([]*st.Strategy, config.GenerationSize)

	for k := 0; k < nOld; k++ {
		strategyIndex = chs[len(chs)-k-1].strategyIndex
		nextGeneration[i] = lastGeneration[strategyIndex]
		i++
	}

	for k := 0; k < nMutated; k++ {
		randNumber = rand.Float64() * maxScore
		strategyIndex := chs.getIndexByChance(randNumber)
		nextGeneration[i] = mutator.MutateStrategy(lastGeneration[strategyIndex])
		i++
	}

	for k := 0; k < nReproduced; k++ {
		randNumber = rand.Float64() * maxScore
		randNumber2 = rand.Float64() * maxScore
		strategyIndex = chs.getIndexByChance(randNumber)
		strategyIndex2 = chs.getIndexByChance(randNumber2)

		nextGeneration[i] = reproducer.Reproduce(lastGeneration[strategyIndex], lastGeneration[strategyIndex2])
		i++
	}

	for k := 0; k < nNew; k++ {
		nextGeneration[i] = random.CreateRandomStrategy(&config.RandomConfig)
		i++
	}
	return nextGeneration
}
