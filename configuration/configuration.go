package configuration

import (
	"fmt"
	"os"
)

type Configuration struct {
	// General
	data            string `yaml:"data"`
	shouldLogOrders bool   `yaml:"should_log_orders"`

	// Test Related
	strategy string `yaml:"strategy"`

	// Evolution
	generationSize       int     `yaml:"generation_size"`
	epochs               int     `yaml:"epochs"`
	oldPercentage        float32 `yaml:"old_percentage"`
	mutatePercentage     float32 `yaml:"mutate_percentage"`
	reproducedPercentage float32 `yaml:"reproduced_percentage"`
	randomPercentage     float32 `yaml:"random_percentage"`

	// Random
	windowMin          int     `yaml:"window_min"`
	windowMax          int     `yaml:"window_max"`
	stopLossMin        float64 `yaml:"stoploss_min"`
	takeProfitMax      float64 `yaml:"takeprofit_max"`
	conditionNumberMin int     `yaml:"condition_number_min"`
	conditionNumberMax int     `yaml:"condition_number_max"`

	// Score
	sumOrdersThreshold      int     `yaml:"sum_orders_threshold"`
	conditionCountThreshold int     `yaml:"condition_count_threshold"`
	conditionCountSlope     float64 `yaml:"condition_count_slope"`
}

func (c *Configuration) Data() string {
	if c.data != "" {
		return c.data
	}
	fmt.Println("Missing data folder")
	os.Exit(1)
	return ""
}

func (c *Configuration) ShouldLogOrders() bool {
	return c.shouldLogOrders
}

func (c *Configuration) Strategy() string {
	if c.strategy != "" {
		return c.strategy
	}
	fmt.Println("Missing strategy file")
	os.Exit(1)
	return ""
}

func (c *Configuration) GenerationSize() int {
	return c.generationSize
}

func (c *Configuration) Epochs() int {
	return c.epochs
}

func (c *Configuration) OldPercentage() float32 {
	return c.oldPercentage
}

func (c *Configuration) MutatePercentage() float32 {
	return c.mutatePercentage
}

func (c *Configuration) ReproducedPercentage() float32 {
	return c.reproducedPercentage
}

func (c *Configuration) RandomPercentage() float32 {
	return c.randomPercentage
}

func (c *Configuration) WindowMin() int {
	return c.windowMin
}

func (c *Configuration) WindowMax() int {
	return c.windowMax
}

func (c *Configuration) StopLossMin() float64 {
	return c.stopLossMin
}

func (c *Configuration) TakeProfitMax() float64 {
	return c.takeProfitMax
}

func (c *Configuration) ConditionNumberMin() int {
	return c.conditionNumberMin
}

func (c *Configuration) ConditionNumberMax() int {
	return c.conditionNumberMax
}

func (c *Configuration) SumOrdersThreshold() int {
	return c.sumOrdersThreshold
}

func (c *Configuration) ConditionCountThreshold() int {
	return c.conditionCountThreshold
}

func (c *Configuration) ConditionCountSlope() float64 {
	return c.conditionCountSlope
}
