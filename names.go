package gofaker

import "github.com/rastogiji/go-faker/data"

func (f *Faker) FirstName() string {
	return data.FirstNames[f.r.Intn(len(data.FirstNames))]
}

func (f *Faker) LastName() string {
	return data.LastNames[f.r.Intn(len(data.LastNames))]
}

func (f *Faker) FullName() string {
	return f.FirstName() + " " + f.LastName()
}
