package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
)

// CheckRequiredFields checks all fielsd of a struct recursiely, and returns an error if any fields
// tagged with required: true, have a zero value
func CheckRequiredFields(originalStruct interface{}) error {
	originalValue := reflect.ValueOf(originalStruct)

	if originalValue.Type().Kind() != reflect.Ptr && originalValue.Type().Kind() != reflect.Struct {
		return errors.New("param is not a struct or pointer to struct")
	}

	var structValue reflect.Value
	if originalValue.Type().Kind() == reflect.Ptr {
		structValue = reflect.Indirect(originalValue)
		if structValue.Kind() != reflect.Struct {
			return errors.New("param is not a struct or pointer to struct")
		}
	} else {
		structValue = originalValue
	}
	structType := structValue.Type()

	for fieldNum := 0; fieldNum < structType.NumField(); fieldNum++ {
		fieldValue := structValue.Field(fieldNum)
		fieldType := structType.Field(fieldNum)

		if fieldType.Type.Kind() == reflect.Struct || (fieldType.Type.Kind() == reflect.Ptr && reflect.Indirect(fieldValue).Type().Kind() == reflect.Struct) {
			err := CheckRequiredFields(fieldValue)
			if err != nil {
				return err
			}
		} else {
			if isRequired := fieldType.Tag.Get("required"); isRequired == "true" {
				if fieldValue == reflect.Zero(fieldType.Type) {
					return fmt.Errorf("Required field %v has a zero-value", fieldType.Name)
				}
			}
		}
	}

	return nil
}

// UpdateStructFromEnv loops through all fields in a struct and replaces the value of any field with a `env:..` tag, with the specified environment variable
func UpdateStructFromEnv(originalStruct interface{}) {
	originalValue := reflect.ValueOf(originalStruct)
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
