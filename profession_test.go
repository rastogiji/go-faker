package gofaker

import "testing"

func TestFaker_Profession(t *testing.T) {
	f := New()
	profession := f.Profession()

	if profession == "" {
		t.Errorf("Expected profession, got empty string")
	}
}
