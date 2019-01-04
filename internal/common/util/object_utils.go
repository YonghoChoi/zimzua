package util

import "github.com/mitchellh/mapstructure"

func MapToObj(input map[string]interface{}, output interface{}) error {
	return mapstructure.Decode(input, output)
}
