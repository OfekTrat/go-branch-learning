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

	for _, file := range data_files {
		fullPath := path.Join(data_path, file)
		streams = append(streams, LoadCandleStreamFromCsv(fullPath))
	}
	return streams
}

func getDataFiles(data_path string) []string {
	var fileList []string
	fileInfo, err := os.Stat(data_path)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if fileInfo.IsDir() {
		dirList, err := ioutil.ReadDir(data_path)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		for _, f := range dirList {
			fileList = append(fileList, f.Name())
		}
	} else {
		fileList = []string{data_path}
	}
	return fileList
}
