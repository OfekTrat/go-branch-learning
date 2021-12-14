package evolutioner

import (
	"branch_learning/utils/random"
)

type EvolutionConfig struct {
	EvolutionLogFile     string
	GenerationSize       int
	NumEvolutions        int
	OldPercentage        float32
	MutatePercentage     float32
	ReproducedPercentage float32
	RandomPercentage     float32

	RandomConfig random.RandomStrategyConfig

	ExitMutateMultiplier int
	WindowSizeMultiplier int
}
