package logger

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

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
