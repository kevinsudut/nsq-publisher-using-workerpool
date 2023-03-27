package csv

import (
	jsoniter "github.com/json-iterator/go"
)

// CSV is struct for save config and content of csv file
type CSV struct {
	fileName string

	rawHeader []string
	rawRows   [][]string

	rows []map[string]interface{}

	isEmptyFile bool
	isEmptyRow  bool
}

// ToJSONString for convert csv data into json string
func (cc *CSV) ParseToJSONString() (string, error) {
	return jsoniter.MarshalToString(cc.rows)
}

// ToSliceMapInterface for convert csv data into slice map interface
func (cc *CSV) ParseToSliceMapInterface() []map[string]interface{} {
	return cc.rows
}

// ToStruct for convert csv data into struct
func (cc *CSV) ParseToStruct(target interface{}) error {
	json, err := cc.ParseToJSONString()
	if err != nil {
		return err
	}

	return jsoniter.UnmarshalFromString(json, target)
}

// IsEmptyFile for check file is empty of not
func (cc *CSV) IsEmptyFile() bool {
	return cc.isEmptyFile
}

// IsEmptyRow for check file has row or not
func (cc *CSV) IsEmptyRow() bool {
	return cc.isEmptyRow
}
