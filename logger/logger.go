package logger

import (
	"branch_learning/configuration"
	"io"
	"log"
	"os"
)

type Logger struct {
	Logs       *log.Logger
	Strategies *log.Logger
	Results    *log.Logger
	Info       *log.Logger
	Error      *log.Logger
}

var logger *Logger = &Logger{}

func InitLoggers(logConf configuration.LogConfiguration) {
	logsFile := createFilePanicError(logConf.LogsFile)
	strategiesFile := createFilePanicError(logConf.StrategiesFile)
	resultsFile := createFilePanicError(logConf.ResultsFile)

	multiWriter := io.MultiWriter(logsFile, os.Stdout)

	logsLogger := log.New(logsFile, "", 0)
	strategiesLogger := log.New(strategiesFile, "", 0)
	resultsLogger := log.New(resultsFile, "", 0)
	infoLogger := log.New(multiWriter, "", 0)
	errorLogger := log.New(os.Stderr, "", 0)

	logger.Logs = logsLogger
	logger.Strategies = strategiesLogger
	logger.Results = resultsLogger
	logger.Info = infoLogger
	logger.Error = errorLogger
}

func createFilePanicError(filename string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	return file
}

func CreateLogger() *Logger {
	return logger
}
