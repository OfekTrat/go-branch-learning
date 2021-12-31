package output

import (
	"log"
	"os"
)

func LogInitialize(output_config *OutputConfig) {
	log.SetFlags(0)

	if output_config.LogFile != "" {
		f, err := os.Create(output_config.LogFile)
		if err != nil {
			f = os.Stdout
		}
		log.SetOutput(f)
	}
}
