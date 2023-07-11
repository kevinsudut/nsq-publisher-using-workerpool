package util

import (
	"errors"
	"reflect"

	jsoniter "github.com/json-iterator/go"
)

// ConvertInterfaceToSliceOfInterface will convert interface to slice of interface
func ConvertInterfaceToSliceOfInterface(data interface{}) ([]interface{}, error) {
	s := reflect.ValueOf(data)
	if s.Kind() != reflect.Slice {
		return nil, errors.New("data must be type of slice")
	}

	if s.IsNil() {
		return nil, errors.New("data is nil")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret, nil
}

// ChunkSliceOfInterface will split interface to each slice based on number of chunk
func ChunkSliceOfInterface(data interface{}, chunkNumber int) ([][]interface{}, error) {
	slice, err := ConvertInterfaceToSliceOfInterface(data)
	if err != nil {
		return nil, err
	}

	var res [][]interface{}

	chunkSize := (len(slice) + chunkNumber - 1) / chunkNumber

	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		if end > len(slice) {
			end = len(slice)
		}

		res = append(res, slice[i:end])
	}

	return res, nil
}

// ConvertSliceOfInterfaceToStruct for convert slice of interface to real target
func ConvertSliceOfInterfaceToStruct(data []interface{}, v interface{}) error {
	byt, err := jsoniter.Marshal(data)
	if err != nil {
		return err
	}

	return jsoniter.Unmarshal(byt, v)
}
