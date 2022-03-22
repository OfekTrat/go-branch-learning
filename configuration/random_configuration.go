package configuration

type RandomConfiguration struct {
	WindowMin          int     `yaml:"min_window"`
	WindowMax          int     `yaml:"max_window"`
	StopLossMin        float64 `yaml:"min_stop_loss"`
	TakeProfitMax      float64 `yaml:"max_take_profit"`
	ConditionNumberMin int     `yaml:"min_condition_number"`
	ConditionNumberMax int     `yaml:"max_condition_number"`
	// ExitMutateMultiplier int     // Currently Not Used
	// WindowSizeMultiplier int     // Currently Not Used
}
