package csv

import (
	"encoding/csv"
	"errors"
	"os"
)

// FileContent is interface that must be implemented by file reader
type FileContent interface {
	ParseToJSONString() (string, error)
	ParseToSliceMapInterface() []map[string]interface{}
	ParseToStruct(target interface{}) error
	IsEmptyFile() bool
	IsEmptyRow() bool
}

const (
	DefaultHeaderIndex   = 0
	DefaultStartRowIndex = 1
)

// ReadFile for initialize csv reader
func ReadFile(fileName string) (FileContent, error) {
	if fileName == "" {
		return nil, errors.New("file name can not be empty")
	}

	csvLib := &CSV{
		fileName: fileName,
	}

	err := csvLib.readFile()
	if err != nil {
		return nil, err
	}

	csvLib.parseRowsToSliceMapInterface()

	return csvLib, nil
}

// readFile for read csv data from file
func (cc *CSV) readFile() error {
	csvFile, err := os.Open(cc.fileName)
	if err != nil {
		return err
	}

	defer csvFile.Close()

	lines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return err
	}

	cc.fillRawData(lines)

	return nil
}

// fillRawData for separated header and rows from csv file
func (cc *CSV) fillRawData(lines [][]string) {
	lenLine := len(lines)

	cc.isEmptyFile = lenLine == DefaultHeaderIndex
	cc.isEmptyRow = lenLine == DefaultStartRowIndex

	if cc.isEmptyFile {
		return
	}

	cc.rawHeader = lines[DefaultHeaderIndex]

	if cc.isEmptyRow {
		return
	}

	cc.rawRows = lines[DefaultStartRowIndex:]
}

// parseRowsToSliceMapInterface for convert rows data to map interface with key based on header
func (cc *CSV) parseRowsToSliceMapInterface() {
	for i := 0; i < len(cc.rawRows); i++ {
		row := map[string]interface{}{}

		for j := 0; j < len(cc.rawRows[i]); j++ {
			row[cc.rawHeader[j]] = cc.rawRows[i][j]
		}

		cc.rows = append(cc.rows, row)
	}
}
