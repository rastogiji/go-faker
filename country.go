package gofaker

import (
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
		return data.AfricanCountries[f.r.Intn(len(data.AfricanCountries))]
	case "asia":
		return data.AsianCountries[f.r.Intn(len(data.AsianCountries))]
	case "europe":
		return data.EuropeanCountries[f.r.Intn(len(data.EuropeanCountries))]
	case "americas":
		return data.AmericanCountries[f.r.Intn(len(data.AmericanCountries))]
	case "oceania":
		return data.OceanianCountries[f.r.Intn(len(data.OceanianCountries))]
	default:
		return data.AllCountries[f.r.Intn(len(data.AllCountries))]
	}

}
