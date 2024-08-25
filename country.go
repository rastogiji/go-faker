package gofaker

import (
	"fmt"
	"github.com/rastogiji/go-faker/data"
	"strings"
)

/*
	Country generates a random country name. s represents the continent name.

Valid values for s are "africa", "asia", "europe", "americas", "oceania". If s is not one of these values, a random country name from anywhere is returned.
*/
func (f *Faker) Country(s string) string {
	switch strings.ToLower(s) {
	case "africa":
		fmt.Println(len(data.AfricanCountries))
		return data.AfricanCountries[f.r.Intn(len(data.AfricanCountries))]
	case "asia":
		fmt.Println(len(data.AsianCountries))
		return data.AsianCountries[f.r.Intn(len(data.AsianCountries))]
	case "europe":
		fmt.Println(len(data.EuropeanCountries))
		return data.EuropeanCountries[f.r.Intn(len(data.EuropeanCountries))]
	case "americas":
		fmt.Println(len(data.AmericanCountries))
		return data.AmericanCountries[f.r.Intn(len(data.AmericanCountries))]
	case "oceania":
		fmt.Println(len(data.OceanianCountries))
		return data.OceanianCountries[f.r.Intn(len(data.OceanianCountries))]
	default:
		fmt.Println(len(data.AllCountries))
		return data.AllCountries[f.r.Intn(len(data.AllCountries))]
	}

}
