package args_parser

import (
	"branch_learning/evolutioner"
	"branch_learning/output"
)

type Configuration struct {
	DataPath        string
	CallType        string
	StrategyFile    string
	EvolutionConfig *evolutioner.EvolutionConfig
	OutputConfig    *output.OutputConfig
}
