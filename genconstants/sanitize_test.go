package genconstants

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeIdentifier_basic(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"spaces", "First name", "FirstName"},
		{"contact id", "contact id", "ContactID"},
		{"percent", "50% off", "_50PctOff"},
		{"rto id", "rto id", "RTOID"},
		{"fallback", "---", "_"},
		{"slash", "something/cool", "SomethingCool"},
		{"postcode", "post-code", "PostCode"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := sanitizeIdentifier(tc.in)
			assert.Equal(t, tc.want, got)
		})
	}
}
