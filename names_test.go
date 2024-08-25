package gofaker

import (
	"strings"
	"testing"
)

func TestFaker_FirstName(t *testing.T) {
	f := New()
	firstName := f.FirstName()

	if firstName == "" {
		t.Errorf("Expected first name, got empty string")
	}
}

func TestFaker_LastName(t *testing.T) {
	f := New()
	lastName := f.LastName()

	if lastName == "" {
		t.Errorf("Expected first name, got empty string")
	}
}

func TestFaker_FullName(t *testing.T) {
	f := New()
	fullName := f.FullName()

	if fullName == "" {
		t.Errorf("Expected full name, got empty string")
	}
	names := strings.Split(fullName, " ")
	if len(names) < 2 {
		t.Errorf("Expected full name to have atleast two strings, got %d", len(names))
	}
}
