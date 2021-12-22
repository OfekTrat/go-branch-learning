package args_parser

import (
	evolution "branch_learning/evolutioner"
	"branch_learning/output"
	"branch_learning/utils/random"
	"flag"
)

const (
	DefaultMultiplier = 5
)

func InitArgs() *Configuration {
	config := Configuration{}
	config.EvolutionConfig = initEvolutionConfig()
	config.EvolutionConfig.RandomConfig = *initRandConfig()
	config.OutputConfig = initOutputConfig()
	flag.StringVar(&config.DataPath, "data", "", "The file with the csv data")
	flag.StringVar(&config.StrategyFile, "strategy", "", "The file in which the strategy to be tested is found")
	flag.Parse()
	config.CallType = flag.Arg(0)
	return &config
}

func initEvolutionConfig() *evolution.EvolutionConfig {
	evoConfig := evolution.EvolutionConfig{}

	var oldPercentage float64
	var mutatePercentage float64
	var reproducedPercentage float64
	var randomPercentage float64

	flag.IntVar(&evoConfig.NumEvolutions, "epochs", 100, "Number of evolutions")
	flag.IntVar(&evoConfig.GenerationSize, "size", 100, "Number of strategies to backtest every iteration")
	flag.Float64Var(&oldPercentage, "old-percentage", 0.05, "The percentage of old strategies that will persist")
	flag.Float64Var(&mutatePercentage, "mutate-percentage", 0.2, "The percentage of strategy mutations that will be created")
	flag.Float64Var(&reproducedPercentage, "reproduced-percentage", 0.2, "The percentage of reproduced strategies that will created")
	flag.Float64Var(&randomPercentage, "new-percentage", 0.55, "The percentage of new strategies that will be generated")

	evoConfig.OldPercentage = float32(oldPercentage)
	evoConfig.MutatePercentage = float32(mutatePercentage)
	evoConfig.ReproducedPercentage = float32(reproducedPercentage)
	evoConfig.RandomPercentage = float32(randomPercentage)

	evoConfig.ExitMutateMultiplier = 5
	evoConfig.WindowSizeMultiplier = 5
	return &evoConfig
}

func initRandConfig() *random.RandomStrategyConfig {
	rand_config := random.RandomStrategyConfig{}
	var exitMin float64
	var exitMax float64

	flag.IntVar(&rand_config.WindowMin, "min-window", 10, "The minimum number of candles a strategy can be created with")
	flag.IntVar(&rand_config.WindowMax, "max-window", 30, "The maximum number of candles a strategy can be created with")
	flag.Float64Var(&exitMin, "min-exit", 0.5, "The minimum stop loss (in percentage) a strategy is created with")
	flag.Float64Var(&exitMax, "max-exit", 1.5, "The maximum take profit (in percentage) a strategy is created with")
	flag.IntVar(&rand_config.ConditionNumberMin, "min-conditions", 1, "The minimum number of conditions a strategy is created with")
	flag.IntVar(&rand_config.ConditionNumberMax, "max-conditions", 4, "The maximum number of conditions a strategy is created with")

	rand_config.ExitMin = float32(exitMin)
	rand_config.ExitMax = float32(exitMax)
	return &rand_config
}

func initOutputConfig() *output.OutputConfig {
	output_config := output.OutputConfig{}

	flag.StringVar(&output_config.LogFile, "log-file", "", "File to be printed all the results")
	flag.StringVar(&output_config.PrintFrequency, "print-frequency", "foreach", "Decides the frequency of printing the strategies: 'foreach' - every iteration the best strategy, 'once' - The best last strategy")
	flag.BoolVar(&output_config.PrintBestStrategy, "print-strategy", false, "Prints the conditions of the best strategy each iteration.")
	return &output_config
}
