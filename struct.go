package gofaker

import (
	"reflect"
	"strconv"
	"strings"
)

// Inserts Fake data into a custom strict using struct tag of faker
func (f *Faker) Struct(v interface{}) {
	val := reflect.Indirect(reflect.ValueOf(v))
	for i := 0; i < val.NumField(); i++ {
		tag := val.Type().Field(i).Tag.Get("faker")
		if tag == "" {
			continue
		}
		tagVal := strings.Split(tag, ",")
		if len(tagVal) == 0 {
			continue
		}
		switch tagVal[0] {
		case "first_name":
			val.Field(i).SetString(f.FirstName())
		case "last_name":
			val.Field(i).SetString(f.LastName())
		case "phone":
			val.Field(i).SetString(f.Phone())
		case "profession":
			val.Field(i).SetString(f.Profession())
		case "country":
			if len(tagVal) == 1 || len(tagVal) >= 2 {
				val.Field(i).SetString(f.Country(tagVal[1]))
			} else {
				continue
			}
		case "email":
			if len(tagVal) >= 2 {
				num, err := strconv.Atoi(tagVal[1])
				if err != nil {
					continue
				}
				val.Field(i).SetString(f.Email(num))
			} else {
				val.Field(i).SetString(f.Email(1))
			}
		}
	}
}
