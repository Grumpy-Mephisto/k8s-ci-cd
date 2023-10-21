package utils

import (
	"os"
	"reflect"
	"strconv"
)

func GetEnv(key string, defaultValue interface{}) interface{} {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	v := reflect.ValueOf(defaultValue)

	switch v.Kind() {
	case reflect.String:
		return value
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(value)
		if err == nil {
			return boolValue
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue, err := strconv.ParseInt(value, 10, v.Type().Bits())
		if err == nil {
			return intValue
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintValue, err := strconv.ParseUint(value, 10, v.Type().Bits())
		if err == nil {
			return uintValue
		}
	case reflect.Float32, reflect.Float64:
		floatValue, err := strconv.ParseFloat(value, v.Type().Bits())
		if err == nil {
			return floatValue
		}
	}

	return defaultValue
}
