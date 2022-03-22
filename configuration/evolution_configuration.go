package configuration

type EvolutionConfiguration struct {
	GenerationSize       int     `yaml:"generation_size"`
	Epochs               int     `yaml:"epochs"`
	OldPercentage        float32 `yaml:"old_percentage"`
	MutatePercentage     float32 `yaml:"mutated_percentage"`
	ReproducedPercentage float32 `yaml:"reproduced_percentage"`
	RandomPercentage     float32 `yaml:"random_percentage"`
}
