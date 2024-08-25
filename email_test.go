package gofaker

import (
	"net/mail"
	"regexp"
	"strconv"
	"testing"
)

func TestFaker_Email(t *testing.T) {
	f := New()
	tests := []struct {
		n int
	}{
		{n: 1},
		{n: 100},
		{n: -10},
		{n: 0},
	}

	for _, tt := range tests {
		email := f.Email(tt.n)
		t.Run("Proper Email", func(t *testing.T) {
			_, err := mail.ParseAddress(email)
			if err != nil {
				t.Errorf("Expected valid email address, got %s", email)
			}
		})
		if tt.n > 2 {
			re := regexp.MustCompile("[0-9]+")
			num := re.FindString(email)
			intNum, err := strconv.Atoi(num)
			if err != nil {
				t.Errorf("Expected number in email, got %s", email)
			}
			if intNum >= tt.n {
				t.Errorf("Expected number in email to be less than %d, got %d", tt.n, intNum)
			}
		} else {
			re := regexp.MustCompile("[0-9]+")
			num := re.FindString(email)
			if num != "" {
				t.Errorf("Expected no number in email, got %s", email)
			}
		}
	}
}
