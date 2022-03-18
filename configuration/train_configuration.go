package configuration

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v3"
)

type TrainConfiguration struct {
	DataPath      string                 `yaml:"data"`
	LogConf       LogConfiguration       `yaml:"log"`
	EvolutionConf EvolutionConfiguration `yaml:"evolution"`
	RandomConf    RandomConfiguration    `yaml:"random"`
}

func ParseTrainConfiguration(confFileName string) *TrainConfiguration {
	confFile, err := ioutil.ReadFile(confFileName)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	configuration := TrainConfiguration{}
	yaml.Unmarshal(confFile, &configuration)
	return &configuration
}
