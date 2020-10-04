package service

import (
	"reflect"
	"strconv"
)

type Utils struct {

}

func (utils *Utils) GetTestObjects(count int) map[string]interface{} {
	var objects []interface{}
	for i := 0; i < count; i++ {
		objects = append(objects, Service{i, "service number: " + strconv.Itoa(i)})
	}
	return buildIndex(objects, "name")
}

type Service struct {
	id int
	name string
}

func buildIndex(collection []interface{}, key string) map[string]interface{} {
	result := make(map[string] interface{}, len(collection))
	for _, object := range collection {
		value := reflect.ValueOf(object)
		keyValue := reflect.Indirect(value).FieldByName(key).String()
		result[keyValue] = object
	}
	return result
}