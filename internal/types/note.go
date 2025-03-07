package types

import "reflect"

type Note struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Text   string `json:"text"`
	Author string `json:"author"`
}

func (n Note) Validate() validationError {

	v := reflect.ValueOf(n)
	t := reflect.TypeOf(n)

	var response validationError

	for i := range v.NumField() {
		fieldName := t.Field(i).Name
		value := v.Field(i)

		if value.Kind() == reflect.String && value.String() == "" {
			response.Ok = false
			response.ErrorFields = append(response.ErrorFields, fieldName)
		}
	}

	return response
}
