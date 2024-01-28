package reflection

import "reflect"

func Walk(x interface{}, fn func(input string)) {
	value := getValue(x)

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			Walk(field.Interface(), fn)
		}
	}
}

func getValue(x any) reflect.Value {
	value := reflect.ValueOf(x)

	if value.Kind() == reflect.Pointer {
		return value.Elem()
	}

	return value
}
