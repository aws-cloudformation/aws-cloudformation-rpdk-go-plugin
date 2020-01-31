package encoding

import (
	"fmt"
	"reflect"
	"strings"
)

// Unstringify takes a stringified representation of a value
// and populates it into the supplied interface
func Unstringify(data map[string]interface{}, v interface{}) error {
	t := reflect.TypeOf(v).Elem()

	val := reflect.ValueOf(v).Elem()

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		jsonName := f.Name
		jsonTag := strings.Split(f.Tag.Get("json"), ",")
		if len(jsonTag) > 0 && jsonTag[0] != "" {
			jsonName = jsonTag[0]
		}

		var newValue reflect.Value

		if value, ok := data[jsonName]; ok {
			switch f.Type.Kind() {
			case reflect.Ptr:
				switch f.Type.Elem().Kind() {
				case reflect.String:
					newValue = toStringPtrValue(value.(string))
				case reflect.Bool:
					newValue = toBoolPtrValue(value.(string))
				case reflect.Int:
					newValue = toIntPtrValue(value.(string))
				case reflect.Float64:
					newValue = toFloat64PtrValue(value.(string))
				case reflect.Struct:
					newValue = reflect.New(f.Type.Elem())
					Unstringify(value.(map[string]interface{}), newValue.Interface())
				}
			case reflect.String:
				newValue = toStringValue(value.(string))
			case reflect.Bool:
				newValue = toBoolValue(value.(string))
			case reflect.Int:
				newValue = toIntValue(value.(string))
			case reflect.Float64:
				newValue = toFloat64Value(value.(string))
			case reflect.Struct:
				newValue = reflect.New(f.Type)
				Unstringify(value.(map[string]interface{}), newValue.Interface())
				newValue = newValue.Elem()
			default:
				return fmt.Errorf("Unsupported type: '%v'", f.Type)
			}

			val.FieldByName(f.Name).Set(newValue)
		}
	}

	return nil
}
