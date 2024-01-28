package reflection

import "reflect"

func Walk(x interface{}, fn func(input string)) {
	value := getValue(x)

	numberOfValues := 0
	var getField func(index int) reflect.Value

	switch value.Kind() {
	case reflect.String:
		fn(value.String())
	case reflect.Struct:
		numberOfValues = value.NumField()
		getField = value.Field
	case reflect.Array, reflect.Slice:
		numberOfValues = value.Len()
		getField = value.Index
	}

	for i := 0; i < numberOfValues; i++ {
		Walk(getField(i).Interface(), fn)
	}
}

func getValue(x any) reflect.Value {
	value := reflect.ValueOf(x)

	if value.Kind() == reflect.Pointer {
		return value.Elem()
	}

	return value
}
