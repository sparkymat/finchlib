package config

import (
	"log"
	"os"
	"reflect"
	"strconv"
)

// UpdateStructFromEnv loops through all fields in a struct and replaces the value of any field with a `env:..` tag, with the specified environment variable. Pass in reflect.ValueOf(pointer to struct)
func UpdateStructFromEnv(originalValue reflect.Value) {
	structValue := reflect.Indirect(originalValue)
	structType := structValue.Type()

	for fieldNum := 0; fieldNum < structType.NumField(); fieldNum++ {
		fieldValue := structValue.Field(fieldNum)
		fieldType := structType.Field(fieldNum)

		if fieldType.Type.Kind() == reflect.Struct {
			UpdateStructFromEnv(fieldValue)
		} else {
			if envName := fieldType.Tag.Get("env"); envName != "" {
				if envValue := os.Getenv(envName); envValue != "" {
					switch fieldType.Type.Kind() {
					case reflect.Int:
						intValue, _ := strconv.Atoi(envValue)
						fieldValue.Set(reflect.ValueOf(intValue))
					case reflect.String:
						fieldValue.Set(reflect.ValueOf(envValue))
					case reflect.Bool:
						if envValue == "true" {
							fieldValue.Set(reflect.ValueOf(true))
						} else {
							fieldValue.Set(reflect.ValueOf(false))
						}
					default:
						log.Fatalf("Error: Unsupported type for field: %v\n", fieldType.Name)
					}
				}
			}
		}
	}
}
