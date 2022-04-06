package logger

import (
	c "branch_learning/configuration"
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

var configuration *c.Configuration = c.GetConfiguration()

type Logger struct {
	orders     *log.Logger
	Strategies *log.Logger
	Results    *log.Logger
	Info       *log.Logger
	Error      *log.Logger
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
	if configuration.ShouldLogOrders() && l.orders != nil {
		l.orders.Printf(message, params...)
	}
}
