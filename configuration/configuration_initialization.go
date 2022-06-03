package configuration

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	yaml "gopkg.in/yaml.v3"
)

var configuration *Configuration

const (
	GENERATION_SIZE           = 100
	EPOCHS                    = 100
	OLD_PERCENTAGE            = 0.05
	MUTATE_PERCENTAGE         = 0.35
	REPRODUCED_PERCENTAGE     = 0.35
	RANDOM_PERCENTAGE         = 0.25
	WINDOW_MIN                = 10
	WINDOW_MAX                = 30
	STOPLOSS_MIN              = 1.5
	TAKE_PROFIT_MAX           = 1.5
	CONDITION_NUMBER_MIN      = 3
	CONDITION_NUMBER_MAX      = 10
	SUM_ORDERS_THRESHOLD      = 100
	CONDITION_COUNT_THRESHOLD = 100
	CONDITION_COUNT_SLOPE     = -0.02
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
	var sumOrdersThreshold int
	var conditionCountThreshold int
	var conditionCountSlope float64
	var output string

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
	flag.IntVar(&sumOrdersThreshold, "sum-orders-threshold", SUM_ORDERS_THRESHOLD, "Threshold of sum orders score")
	flag.IntVar(&conditionCountThreshold, "condition-count-threshold", CONDITION_COUNT_THRESHOLD, "Threshold of condition count")
	flag.Float64Var(&conditionCountSlope, "condition-count-slope", CONDITION_COUNT_SLOPE, "Slope for decreasing condition count weight")
	flag.StringVar(&output, "output", generateDefaultOutputFile(), "The output file for the results")
	flag.Parse()

	configuration = &Configuration{
		data:                    data,
		shouldLogOrders:         shouldLogOrders,
		strategy:                strategy,
		generationSize:          generationSize,
		epochs:                  epochs,
		oldPercentage:           float32(oldPercentage),
		mutatePercentage:        float32(mutatePercentage),
		reproducedPercentage:    float32(reproducedPercentage),
		randomPercentage:        float32(randomPercentage),
		windowMin:               windowMin,
		windowMax:               windowMax,
		stopLossMin:             stopLossMin,
		takeProfitMax:           takeProfitMax,
		conditionNumberMin:      conditionNumberMin,
		conditionNumberMax:      conditionNumberMax,
		sumOrdersThreshold:      sumOrdersThreshold,
		conditionCountThreshold: conditionCountThreshold,
		conditionCountSlope:     conditionCountSlope,
		output:                  output,
	}

	if configFile != "" {
		parseYamlConfiguration(configFile)
	}
}

func GetConfiguration() *Configuration {
	return configuration
}

func generateDefaultOutputFile() string {
	now := time.Now()
	filename := strconv.FormatInt(int64(now.Year()), 10) +
		strconv.FormatInt(int64(now.Month()), 10) +
		strconv.FormatInt(int64(now.Day()), 10) +
		"-" +
		strconv.FormatInt(int64(now.Hour()), 10) +
		strconv.FormatInt(int64(now.Minute()), 10)

	return filename + ".zip"
}

func parseYamlConfiguration(filename string) {
	confFile, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	parsedConfiguration := make(map[string]interface{})
	yaml.Unmarshal(confFile, parsedConfiguration)

	configuration.data = getFirstOrDefaultString(parsedConfiguration["data"], configuration.data)
	configuration.shouldLogOrders = getFirstOrDefaultBool(parsedConfiguration["should_log_orders"], configuration.shouldLogOrders)
	configuration.strategy = getFirstOrDefaultString(parsedConfiguration["strategy"], configuration.strategy)

	configuration.generationSize = getFirstOrDefaultInt(parsedConfiguration["generation_size"], configuration.generationSize)
	configuration.epochs = getFirstOrDefaultInt(parsedConfiguration["epochs"], configuration.epochs)
	configuration.oldPercentage = getFirstOrDefaultFloat32(parsedConfiguration["old_percentage"], configuration.oldPercentage)
	configuration.mutatePercentage = getFirstOrDefaultFloat32(parsedConfiguration["mutate_percentage"], configuration.mutatePercentage)
	configuration.reproducedPercentage = getFirstOrDefaultFloat32(parsedConfiguration["reproduced_percentage"], configuration.reproducedPercentage)
	configuration.randomPercentage = getFirstOrDefaultFloat32(parsedConfiguration["random_percentage"], configuration.randomPercentage)
	configuration.windowMin = getFirstOrDefaultInt(parsedConfiguration["window_min"], configuration.windowMin)
	configuration.windowMax = getFirstOrDefaultInt(parsedConfiguration["window_max"], configuration.windowMax)
	configuration.stopLossMin = getFirstOrDefaultFloat64(parsedConfiguration["stoploss_min"], configuration.stopLossMin)
	configuration.takeProfitMax = getFirstOrDefaultFloat64(parsedConfiguration["takeprofit_max"], configuration.takeProfitMax)
	configuration.conditionNumberMin = getFirstOrDefaultInt(parsedConfiguration["condition_number_min"], configuration.conditionNumberMin)
	configuration.conditionNumberMax = getFirstOrDefaultInt(parsedConfiguration["condition_number_max"], configuration.conditionNumberMax)
	configuration.sumOrdersThreshold = getFirstOrDefaultInt(parsedConfiguration["sum_orders_threshold"], configuration.sumOrdersThreshold)
	configuration.conditionCountThreshold = getFirstOrDefaultInt(parsedConfiguration["condition_count_threshold"], configuration.conditionCountThreshold)
	configuration.conditionCountSlope = getFirstOrDefaultFloat64(parsedConfiguration["condition_count_slope"], configuration.conditionCountSlope)
	configuration.output = getFirstOrDefaultString(parsedConfiguration["output"], configuration.output)
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
