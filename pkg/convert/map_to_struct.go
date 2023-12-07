package convert

import (
	"fmt"
	"reflect"
	"strings"
)

func MapToStruct(mapData map[string]interface{}, resultStruct interface{}) error {
	fmt.Println(mapData)
	//使用反射将哈希数据映射到结构体
	for key, value := range mapData {
		field := toCamelCase(key)
		err := setField(resultStruct, field, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func setField(resultStruct interface{}, fieldName string, value interface{}) error {
	//获取结构体字段
	//fmt.Println(resultStruct, fieldName, value)
	//field := reflect.ValueOf(resultStruct).Elem().FieldByName(fieldName)
	//ptrValue := reflect.ValueOf(resultStruct)
	//if ptrValue.Kind() != reflect.Ptr {
	//	fmt.Println("Not a pointer to struct")
	//	return nil
	//}
	//
	//field := reflect.ValueOf(resultStruct).Elem().FieldByName(fieldName)
	//if !field.IsValid() {
	//	return fmt.Errorf("Field %s not found in struct", fieldName)
	//}

	// Use reflection to get the pointer to the struct
	ptrValue := reflect.ValueOf(resultStruct)
	if ptrValue.Kind() != reflect.Ptr {
		fmt.Println("Not a pointer to struct")
		return nil
	}

	// Dereference the pointer to get the struct value
	structValue := ptrValue.Elem()
	if structValue.Kind() != reflect.Struct {
		fmt.Println("Not a pointer to struct")
		return nil
	}

	// Get the field by name
	field := structValue.FieldByName(fieldName)

	//获取字段类型
	fieldType := field.Type()
	switch fieldType.Kind() {
	case reflect.String:
		field.SetString(value.(string))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue := StrTo(value.(string)).MustInt64()
		field.SetInt(intValue)
	case reflect.Struct:
		//field.Set(reflect.ValueOf(value))
	default:
		//return fmt.Errorf("Unsupported field type: %v", fieldType.Kind())
	}
	return nil
}

// 将下划线分隔的字符串转换为驼峰命名
func toCamelCase(str string) string {
	parts := strings.Split(str, "_")
	for i := 0; i < len(parts); i++ {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}

//// 递归设置结构体字段值
//func setValues(v interface{}, values map[string]interface{}) {
//	val := reflect.ValueOf(v).Elem()
//
//	// 遍历结构体字段
//	for i := 0; i < val.NumField(); i++ {
//		field := val.Type().Field(i).Name
//		if value, ok := values[field]; ok {
//			fieldValue := val.Field(i)
//
//			// 如果字段是嵌套结构体，则递归调用setValues
//			if fieldValue.Kind() == reflect.Struct {
//				setValues(fieldValue.Addr().Interface(), value.(map[string]string))
//			} else {
//				// 否则直接设置字段值
//				fieldValue.Set(reflect.ValueOf(value))
//			}
//		}
//	}
//}
