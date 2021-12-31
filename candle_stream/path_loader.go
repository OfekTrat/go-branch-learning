package candlestream

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

func GetStreamsFromPath(data_path string) []*CandleStream {
	var streams []*CandleStream
	data_files := getDataFiles(data_path)

	for _, filepath := range data_files {
		streams = append(streams, LoadCandleStreamFromCsv(filepath))
	}
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
