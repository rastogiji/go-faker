package gofaker

import (
	"reflect"
	"testing"
)

func TestFaker_Struct(t *testing.T) {
	f := New()

	tests := []struct {
		tag  string
		name string
	}{
		{tag: "first_name", name: "FirstName"},
		{tag: "last_name", name: "LastName"},
		{tag: "phone", name: "Phone"},
		{tag: "profession", name: "Profession"},
		{tag: "country,americas", name: "Country"},
		{tag: "country,123", name: "Country"},
		{tag: "email,32", name: "Email"},
		{tag: "email", name: "Email"},
		{tag: "email,12,34", name: "Email"},
		{tag: "email,qw", name: "Email"},
	}

	for _, tt := range tests {
		t.Run("Struct Test", func(t *testing.T) {
			field := []reflect.StructField{
				{
					Name: tt.name,
					Type: reflect.TypeOf(""),
					Tag:  reflect.StructTag("faker:\"" + tt.tag + "\""),
				},
			}
			testType := reflect.StructOf(field)
			testVal := reflect.New(testType).Interface()

			f.Struct(testVal)
			if reflect.ValueOf(testVal).Elem().Field(0).String() == "" {
				t.Errorf("Expected non-empty string, got empty string")
			}
		})
	}
}
