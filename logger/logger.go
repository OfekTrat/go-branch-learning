package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	TMP_DIR_PATH    = "tmp_dir"
	ORDERS_FILE     = "orders"
	STRATEGIES_FILE = "strategies"
	RESULTS_FILE    = "results"
	LOGS_FILE       = "logs"
)

type Logger struct {
	shouldLogOrders bool
	orders          *log.Logger
	Strategies      *log.Logger
	Results         *log.Logger
	Info            *log.Logger
	Error           *log.Logger
}

func CreateLogger() *Logger {
	if logger != nil {
		return logger
	}
	fmt.Println("Logger is not initialized")
	os.Exit(1)
	return nil
}

func (l *Logger) LogOrder(message string, params ...interface{}) {
	if l.shouldLogOrders && l.orders != nil {
		l.orders.Printf(message, params...)
	}
}
