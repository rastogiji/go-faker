package gofaker

import (
	"regexp"
	"testing"
)

func TestFaker_Phone(t *testing.T) {
	f := New()
	phone := f.Phone()

	if len(phone) != 12 {
		t.Errorf("Expected phone length 11, got %d", len(phone))
	}

	mat, _ := regexp.MatchString(`^\d{3}-\d{3}-\d{4}$`, phone)
	if !mat {
		t.Errorf("Expected phone format XXX-XXX-XXXX, got %s", phone)
	}
}
