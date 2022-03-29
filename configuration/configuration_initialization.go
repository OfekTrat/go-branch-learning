package configuration

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v3"
)

var configuration *Configuration

const (
	GENERATION_SIZE       = 100
	EPOCHS                = 100
	OLD_PERCENTAGE        = 0.05
	MUTATE_PERCENTAGE     = 0.35
	REPRODUCED_PERCENTAGE = 0.35
	RANDOM_PERCENTAGE     = 0.25
	WINDOW_MIN            = 10
	WINDOW_MAX            = 30
	STOPLOSS_MIN          = 1.5
	TAKE_PROFIT_MAX       = 1.5
	CONDITION_NUMBER_MIN  = 3
	CONDITION_NUMBER_MAX  = 10
)

func init() {
	var configFile string
	var data string
	var shouldLogOrders bool
	var strategy string
	var generationSize int
	var epochs int
	var oldPercentage float64
	var mutatePercentage float64
	var reproducedPercentage float64
	var randomPercentage float64
	var windowMin int
	var windowMax int
	var stopLossMin float64
	var takeProfitMax float64
	var conditionNumberMin int
	var conditionNumberMax int

	flag.StringVar(&configFile, "f", "", "Yaml configuration file")
	flag.StringVar(&data, "d", "", "Data Path")
	flag.BoolVar(&shouldLogOrders, "debug", false, "Debug mode - logs orders")
	flag.StringVar(&strategy, "strategy", "", "Strategy to test path")
	flag.IntVar(&generationSize, "generation-size", GENERATION_SIZE, "Generation size")
	flag.IntVar(&epochs, "epochs", EPOCHS, "Epochs (Iterations)")
	flag.Float64Var(&oldPercentage, "old-percentage", float64(OLD_PERCENTAGE), "Old percentage in created generation")
	flag.Float64Var(&mutatePercentage, "mutate-percentage", float64(MUTATE_PERCENTAGE), "Mutate percentage in created generation")
	flag.Float64Var(&reproducedPercentage, "reproduced-percentage", float64(OLD_PERCENTAGE), "Reproduced percentage in created generation")
	flag.Float64Var(&randomPercentage, "random-percentage", float64(RANDOM_PERCENTAGE), "Random percentage in created generation")
	flag.IntVar(&windowMin, "window-min", WINDOW_MIN, "Minimum window size random strategy starts with")
	flag.IntVar(&windowMax, "window-max", WINDOW_MAX, "Maximum window size random strategy starts with")
	flag.Float64Var(&stopLossMin, "stoploss", STOPLOSS_MIN, "The maximum stoploss a random strategy will start with")
	flag.Float64Var(&takeProfitMax, "takeprofit", TAKE_PROFIT_MAX, "The maximum takeprofit a random strategy will start with")
	flag.IntVar(&conditionNumberMin, "condition-number-min", CONDITION_NUMBER_MIN, "Minimum condition number")
	flag.IntVar(&conditionNumberMax, "condition-number-max", CONDITION_NUMBER_MAX, "Maximum condition number")
	flag.Parse()

	configuration = &Configuration{
		data:                 data,
		shouldLogOrders:      shouldLogOrders,
		strategy:             strategy,
		generationSize:       generationSize,
		epochs:               epochs,
		oldPercentage:        float32(oldPercentage),
		mutatePercentage:     float32(mutatePercentage),
		reproducedPercentage: float32(reproducedPercentage),
		randomPercentage:     float32(randomPercentage),
		windowMin:            windowMin,
		windowMax:            windowMax,
		stopLossMin:          stopLossMin,
		takeProfitMax:        takeProfitMax,
		conditionNumberMin:   conditionNumberMin,
		conditionNumberMax:   conditionNumberMax,
	}

	if configFile != "" {
		parseYamlConfiguration(configFile)
	}
}

func GetConfiguration() *Configuration {
	return configuration
}

func parseYamlConfiguration(filename string) {
	confFile, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	parsedConfiguration := make(map[string]interface{})
	yaml.Unmarshal(confFile, parsedConfiguration)

	configuration.data = getFirstOrDefaultString(parsedConfiguration["data"], "")
	configuration.shouldLogOrders = getFirstOrDefaultBool(parsedConfiguration["shouldLogOrders"], false)
	configuration.strategy = getFirstOrDefaultString(parsedConfiguration["strategy"], "")

	configuration.generationSize = getFirstOrDefaultInt(parsedConfiguration["generation_size"], GENERATION_SIZE)
	configuration.epochs = getFirstOrDefaultInt(parsedConfiguration["epochs"], EPOCHS)
	configuration.oldPercentage = getFirstOrDefaultFloat32(parsedConfiguration["old_percentage"], OLD_PERCENTAGE)
	configuration.mutatePercentage = getFirstOrDefaultFloat32(parsedConfiguration["mutate_percentage"], MUTATE_PERCENTAGE)
	configuration.reproducedPercentage = getFirstOrDefaultFloat32(parsedConfiguration["reproduced_percentage"], REPRODUCED_PERCENTAGE)
	configuration.randomPercentage = getFirstOrDefaultFloat32(parsedConfiguration["random_percentage"], RANDOM_PERCENTAGE)
	configuration.windowMin = getFirstOrDefaultInt(parsedConfiguration["window_min"], WINDOW_MIN)
	configuration.windowMax = getFirstOrDefaultInt(parsedConfiguration["window_max"], WINDOW_MAX)
	configuration.stopLossMin = getFirstOrDefaultFloat64(parsedConfiguration["stoploss_min"], STOPLOSS_MIN)
	configuration.takeProfitMax = getFirstOrDefaultFloat64(parsedConfiguration["takeprofit_max"], TAKE_PROFIT_MAX)
	configuration.conditionNumberMin = getFirstOrDefaultInt(parsedConfiguration["condition_number_min"], CONDITION_NUMBER_MIN)
	configuration.conditionNumberMax = getFirstOrDefaultInt(parsedConfiguration["condition_number_max"], CONDITION_NUMBER_MAX)
}

func getFirstOrDefaultString(value interface{}, defaultValue string) string {
	if value != nil {
		return value.(string)
	}
	return defaultValue
}

func getFirstOrDefaultBool(value interface{}, defaultValue bool) bool {
	if value != nil {
		return value.(bool)
	}
	return defaultValue
}

func getFirstOrDefaultInt(value interface{}, defaultValue int) int {
	if value != nil {
		return value.(int)
	}
	return defaultValue
}

func getFirstOrDefaultFloat32(value interface{}, defaultValue float32) float32 {
	if value != nil {
		return float32(value.(float64))
	}
	return defaultValue
}

func getFirstOrDefaultFloat64(value interface{}, defaultValue float64) float64 {
	if value != nil {
		return value.(float64)
	}
	return defaultValue
}
