package gofaker

import "testing"

func TestFaker_Country(t *testing.T) {
	f := New()

	tests := []struct {
		continent string
	}{
		{continent: "americas"},
		{continent: "europe"},
		{continent: "africa"},
		{continent: "asia"},
		{continent: "oceania"},
		{continent: "random"},
		{continent: ""},
		{continent: "Africa"},
		{continent: "1"},
	}

	for _, tt := range tests {
		t.Run("Country Test", func(t *testing.T) {
			country := f.Country(tt.continent)
			if country == "" {
				t.Errorf("Expected country name, got empty string")
			}
		})
	}
}
