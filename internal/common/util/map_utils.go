package util

import "reflect"

func MapKeys(target interface{}) []string {
	keys := reflect.ValueOf(target).MapKeys()

	var resultMap []string
	for _, key := range keys {
		resultMap = append(resultMap, key.String())
	}

	return resultMap
}

func MapKeysInt(target interface{}) []int {
	keys := reflect.ValueOf(target).MapKeys()

	var resultMap []int
	for _, key := range keys {
		resultMap = append(resultMap, int(key.Int()))
	}

	return resultMap
}
