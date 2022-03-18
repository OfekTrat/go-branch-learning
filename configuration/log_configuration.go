package configuration

type LogConfiguration struct {
	LogsFile       string `yaml:"logs"`
	StrategiesFile string `yaml:"strategies"`
	ResultsFile    string `yaml:"results"`
}
