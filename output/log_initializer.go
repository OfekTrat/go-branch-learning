package output

import (
	"fmt"
	"log"
	"os"
)

func LogInitialize(output_config *OutputConfig) {
	log.SetFlags(0)

	if output_config.LogFile != "" {
		f, err := os.Create(output_config.LogFile)
		if err != nil {
			fmt.Println("Something went wrong with creating the log file. Printing to std out")
			f = os.Stdout
		}
		log.SetOutput(f)
	} else {
		log.SetOutput(os.Stdout)
	}
}
