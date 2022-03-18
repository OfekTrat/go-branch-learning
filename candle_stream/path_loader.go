package candlestream

import (
	l "branch_learning/logger"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var logger *l.Logger = l.CreateLogger()

func GetStreamsFromPath(data_path string) []*CandleStream {
	var streams []*CandleStream
	data_files := getDataFiles(data_path)

	logger.Info.Printf("GetStreamsFromPath - Starting to collect data from %s\n", data_path)
	candlesSum := 0

	for _, filepath := range data_files {
		stream := LoadCandleStreamFromCsv(filepath)
		candlesSum += stream.Length()
		streams = append(streams, stream)
	}
	logger.Info.Printf("GetStreamsFromPath - Done. Collected from %d data files. Total of %d candles\n\n", len(streams), candlesSum)
	return streams
}

func getDataFiles(data_path string) []string {
	fileInfo, err := os.Stat(data_path)
	if err != nil {
		log.Println("Please specify an existing data path")
		os.Exit(1)
	}
	if fileInfo.IsDir() {
		return getDirList(data_path)
	} else {
		return []string{data_path}
	}
}

func getDirList(data_path string) []string {
	fileList := []string{}
	dirList, err := ioutil.ReadDir(data_path)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	for _, f := range dirList {
		fileList = append(fileList, path.Join(data_path, f.Name()))
	}
	return fileList
}
