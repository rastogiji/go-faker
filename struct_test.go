package gofaker

import (
	"math/rand"
	"testing"
)

func TestFaker_Struct(t *testing.T) {
	type fields struct {
		r *rand.Rand
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Faker{
				r: tt.fields.r,
			}
			f.Struct(tt.args.v)
		})
	}
}
