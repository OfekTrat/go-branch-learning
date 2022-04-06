package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var logger *Logger
var logsFile *os.File
var ordersFile *os.File
var strategiesFile *os.File
var resultsFile *os.File

func init() {
	// Default behavior is not to log orders
	InitLoggers(false)
}

func InitLoggers(shouldLogOrders bool) {
	makeDir(TMP_DIR_PATH)
	logsFile = createFilePanicError(LOGS_FILE)
	strategiesFile = createFilePanicError(STRATEGIES_FILE)
	resultsFile = createFilePanicError(RESULTS_FILE)

	multiWriter := io.MultiWriter(logsFile, os.Stdout)

	strategiesLogger := log.New(strategiesFile, "", 0)
	resultsLogger := log.New(resultsFile, "", 0)
	infoLogger := log.New(multiWriter, "", 0)
	errorLogger := log.New(os.Stderr, "", 0)

	logger = &Logger{}
	logger.Strategies = strategiesLogger
	logger.Results = resultsLogger
	logger.Info = infoLogger
	logger.Error = errorLogger

	if configuration.ShouldLogOrders() {
		ordersFile = createFilePanicError(ORDERS_FILE)
		logger.orders = log.New(ordersFile, "", 0)
		logger.orders.Println("coin,buy_time,generation,strategy,buy_sell,price,close_time")
	}

	logger.Results.Println("generation,id,conditionCount,wins,losses,winRate,takeProfit,stopLoss,Score")
}

func makeDir(path string) {
	err := os.Mkdir(path, 0755)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func createFilePanicError(filename string) *os.File {
	path := filepath.Join(TMP_DIR_PATH, filename)
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return file
}
