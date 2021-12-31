package log_init

import (
	"branch_learning/output"
	"log"
	"os"
)

func LogInitialize(output_config *output.OutputConfig) {
	log.SetFlags(0)

	if output_config.LogFile != "" {
		f, err := os.Create(output_config.LogFile)
		if err != nil {
			f = os.Stdout
		}
		log.SetOutput(f)
	}
}
