package util

import (
	"errors"
	"math"
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

	sliceSize := len(slice)
	res := make([][]interface{}, 0, chunkNumber)
	chunkSize := int(math.Ceil(float64(sliceSize) / float64(chunkNumber)))

	for i := 0; i < chunkNumber-1; i++ {
		if i+1 > sliceSize {
			break
		}

		res = append(res, slice[i*chunkSize:(i+1)*chunkSize])
	}

	if chunkNumber < sliceSize {
		res = append(res, slice[(chunkNumber-1)*chunkSize:])
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
