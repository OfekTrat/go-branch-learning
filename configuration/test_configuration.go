package configuration

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v3"
)

type TestConfiguration struct {
	DataPath string           `yaml:"data_path"`
	Strategy string           `yaml:"strategy"`
	LogConf  LogConfiguration `yaml:"log"`
}

func ParseTestConfiguration(confFileName string) *TestConfiguration {
	confFile, err := ioutil.ReadFile(confFileName)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	configuration := TestConfiguration{}
	yaml.Unmarshal(confFile, &configuration)
	return &configuration
}
