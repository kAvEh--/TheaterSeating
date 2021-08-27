package seating

import (
	"encoding/json"
	"github.com/kAvEh--/TheaterSeating/cmd/database"
	"reflect"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func deepCopy(v database.Hall) (database.Hall, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return database.Hall{}, err
	}

	vptr := reflect.New(reflect.TypeOf(v))
	err = json.Unmarshal(data, vptr.Interface())
	if err != nil {
		return database.Hall{}, err
	}
	return vptr.Elem().Interface().(database.Hall), err
}
