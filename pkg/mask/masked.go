package mask

import (
	"reflect"
	"strconv"
	"time"
)

const AccessLevel = "accessLevel"

// MaskedMapResponse function that masks fields, including pointers, nested structs, and slices, based on access level tags
func MaskedMapResponse(data interface{}, accessLevel int) interface{} {
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	// Dereference pointers to get the underlying value
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil // Handle nil pointers as null
		}
		val = reflect.Indirect(val)
		typ = typ.Elem()
	}

	if val.Kind() == reflect.Struct {
		result := make(map[string]interface{})

		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fieldType := typ.Field(i)
			jsonKey := fieldType.Tag.Get("json") // Get the JSON key name

			// Check if there's an access level tag on the field
			tag := fieldType.Tag.Get("access_level")
			if tag != "" {
				requiredAccess, _ := strconv.Atoi(tag)
				if accessLevel < requiredAccess {
					// Mask the field if access level is insufficient
					result[jsonKey] = "*****"
					continue
				}
			}

			// Handle time.Time or *time.Time types without masking
			if field.Type() == reflect.TypeOf(time.Time{}) || field.Type() == reflect.TypeOf(&time.Time{}) {
				result[jsonKey] = field.Interface()
				continue
			}

			// Handle nested structs, slices, and pointers recursively
			result[jsonKey] = MaskedMapResponse(field.Interface(), accessLevel)
		}
		return result

	} else if val.Kind() == reflect.Slice {
		sliceResult := make([]interface{}, val.Len())
		for i := 0; i < val.Len(); i++ {
			sliceResult[i] = MaskedMapResponse(val.Index(i).Interface(), accessLevel)
		}
		return sliceResult
	}

	// For other data types, return the value directly
	return val.Interface()
}
