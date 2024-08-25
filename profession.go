package gofaker

import "github.com/rastogiji/go-faker/data"

func (f *Faker) Profession() string {
	return data.Professions[f.r.Intn(len(data.Professions))]
}
