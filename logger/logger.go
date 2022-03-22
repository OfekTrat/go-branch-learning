package logger

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
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
}

func (l *Logger) LogOrder(message string, params ...interface{}) {
	if l.shouldLogOrders && l.orders != nil {
		l.orders.Printf(message, params...)
	}
}

func EnableOrdersLogs() {
	logger.shouldLogOrders = true
	ordersFile = createFilePanicError(ORDERS_FILE)
	ordersLogger := log.New(ordersFile, "", 0)
	logger.orders = ordersLogger
	logger.LogOrder("ticker,time,generation,strategy,type,price\n")
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

func CreateLogger() *Logger {
	if logger != nil {
		return logger
	}
	fmt.Println("Logger is not initialized")
	os.Exit(1)
	return nil
}

func ZipLogs(path string) {
	zipFile, err := os.Create(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	filepath.Walk(TMP_DIR_PATH, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Method = zip.Deflate
		header.Name, err = filepath.Rel(filepath.Dir(TMP_DIR_PATH), path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += "/"
		}

		headerWriter, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(headerWriter, f)
		return err
	})
	closeFiles()
	deleteFolder(TMP_DIR_PATH)
}

func closeFiles() {
	closeFile(logsFile)
	closeFile(strategiesFile)
	closeFile(resultsFile)

	if logger.shouldLogOrders {
		closeFile(ordersFile)
	}
}

func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func deleteFolder(filename string) {
	err := os.RemoveAll(filename)
	if err != nil {
		fmt.Println("Failed to remove file", filename)
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
