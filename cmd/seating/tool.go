package seating

import (
	"encoding/json"
	"github.com/kAvEh--/TheaterSeating/cmd/model"
	"reflect"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func deepCopy(v model.Hall) (model.Hall, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return model.Hall{}, err
	}

	vptr := reflect.New(reflect.TypeOf(v))
	err = json.Unmarshal(data, vptr.Interface())
	if err != nil {
		return model.Hall{}, err
	}
	return vptr.Elem().Interface().(model.Hall), err
}
