package candlestream

import (
	candle "branch_learning/candle"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func LoadCandleStreamFromCsv(csvpath string) *CandleStream {
	candleList := []candle.Candle{}
	filename := strings.TrimSuffix(filepath.Base(csvpath), ".csv")
	fileData, err := os.ReadFile(csvpath)

	if err != nil {
		panic(err)
	}
	fileStr := string(fileData)
	lines := strings.Split(fileStr, "\n")
	headers := strings.Split(strings.Replace(strings.ToLower(lines[0]), "\r", "", -1), ",")

	for _, line := range lines[1:] {
		if line == "\r" || line == "" {
			continue
		}
		splittedLine := strings.Split(strings.Replace(line, "\r", "", -1), ",")

		c, err := createCandleFromLines(headers, splittedLine)
		if err != nil {
			panic(err)
		}
		candleList = append(candleList, c)
	}
	return CreateCandleStream(filename, candleList)
}

func createCandleFromLines(headers []string, line []string) (candle.Candle, error) {
	candleMap := make(map[string]float32)

	for i := 0; i < len(headers); i++ {
		val, err := strconv.ParseFloat(line[i], 32)
		if err != nil {
			return candle.Candle{}, err
		}
		candleMap[headers[i]] = float32(val)
	}
	return candle.CreateCandle(candleMap), nil
}
