package utils

import (
	"encoding/json"
	"reflect"
	"strconv"
)

func ToJson(data map[string]interface{}) string {
	body, _ := json.Marshal(data)
	return string(body[:])
}

func GetTestObjects(count int) map[string]interface{} {
	var objects []interface{}
	for i := 0; i < count; i++ {
		objects = append(objects, Service{i, "service number: " + strconv.Itoa(i)})
	}
	return BuildIndex(objects, "name")
}

func BuildIndex(collection []interface{}, key string) map[string]interface{} {
	result := make(map[string]interface{}, len(collection))
	for _, object := range collection {
		value := reflect.ValueOf(object)
		keyValue := reflect.Indirect(value).FieldByName(key).String()
		result[keyValue] = object
	}
	return result
}

type Service struct {
	id   int
	name string
}
