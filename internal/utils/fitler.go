package utils

import "reflect"

func Filter[T any](vs []T, pred func(t T, ind int) bool) []T {
	var result []T

	for i, val := range vs {
		if pred(val, i) {
			result = append(result, val)
		}
	}

	return result
}

func Find[T any](vs []T, pred func(t T, ind int) bool) *T {

	var t *T
	for i, val := range vs {
		if pred(val, i) {
			t = &vs[i]
		}
	}
	return t
}

/*
Spread takes to generic structs and returns a "merged" struct where `new` is spread into `old`,
```
	return {...old, ...new} //javscript like
```
*/
func Spread[T any](old, new T) T {

	out := old

	oldVal := reflect.ValueOf(&out).Elem()
	newVal := reflect.ValueOf(new)
	if newVal.Kind() == reflect.Ptr {
		newVal = newVal.Elem() // Dereference if new is a pointer
	}

	for i := 0; i < newVal.NumField(); i++ {
		field := newVal.Type().Field(i)
		srcField := newVal.Field(i)

		if !srcField.IsZero() {
			oldVal.FieldByName(field.Name).Set(srcField)
		}
	}

	return out

}
