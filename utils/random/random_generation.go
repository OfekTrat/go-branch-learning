package random

import st "branch_learning/strategy"

func CreateRandomGeneration(generationNumber int, generationSize int, config *RandomStrategyConfig) []*st.Strategy {
	generation := make([]*st.Strategy, generationSize)

	for i := 0; i < generationSize; i++ {
		generation[i] = CreateRandomStrategy(i, generationNumber, config)
	}
	return generation
}
