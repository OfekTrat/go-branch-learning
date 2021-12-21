package args_parser

import (
	"branch_learning/evolutioner"
	"branch_learning/output"
)

type Configuration struct {
	DataFile        string
	CallType        string
	StrategyFile    string
	EvolutionConfig *evolutioner.EvolutionConfig
	OutputConfig    *output.OutputConfig
}
