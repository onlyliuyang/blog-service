package util

import (
	"reflect"
)

func StructToMap(st interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	t := reflect.TypeOf(st)
	v := reflect.ValueOf(st)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i).Tag.Get("json")
		value := v.Field(i).Interface()
		m[field] = value
	}
	return m
}

func StructToStruct(from interface{}, to interface{}) error {
	fromType := reflect.TypeOf(from)
	fromValue := reflect.ValueOf(from)

	toType := reflect.TypeOf(to)
	toValue := reflect.ValueOf(to)

	for i := 0; i < fromType.NumField(); i++ {
		fieldName := fromType.Field(i).Name
		value := fromValue.Field(i)
		_, hasField := toType.FieldByName(fieldName)
		if hasField {
			if toValue.FieldByName(fieldName).CanSet() {
				switch value.Kind() {
				case reflect.Int:
					toValue.FieldByName(fieldName).SetInt(value.Int())
				}
			}
		}
	}
	return nil
}
