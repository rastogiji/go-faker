package gofaker

import "fmt"

func (f *Faker) Phone() string {
	return fmt.Sprintf("%d-%d-%d", f.r.Intn(899)+100, f.r.Intn(899)+100, f.r.Intn(8999)+1000)
}
