package gofaker

import (
	"fmt"
	"strings"
)

var (
	domains = []string{"gmail.com", "hotmail.com", "outlook.com", "yahoo.com", "aol.com", "protonmail.com", "zoho.com", "icloud.com", "mail.com", "yandex.com", "gmx.com", "tutanota.com", "disroot.org", "riseup.com"}
)

// Email generates a random email address. n is used to generate a random number between 0 and n to be appended to the email address like abc.xyz123@example.com. if n <= 0 || 1, no number will be appended.
func (f *Faker) Email(n int) string {
	if n <= 0 || n == 1 {
		return fmt.Sprintf("%s.%s@%s", strings.ToLower(f.FirstName()), strings.ToLower(f.LastName()), domains[f.r.Intn(len(domains))])
	}
	return fmt.Sprintf("%s.%s%d@%s", strings.ToLower(f.FirstName()), strings.ToLower(f.LastName()), f.r.Intn(n), domains[f.r.Intn(len(domains))])
}
