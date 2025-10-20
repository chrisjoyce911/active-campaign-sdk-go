package users

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsers_NotImplementedPlaceholders(t *testing.T) {
	t.Run("users placeholders", func(t *testing.T) {
		tests := []struct {
			name string
			fn   func() error
		}{
			{"CreateUser", func() error {
				_, _, err := (&service{}).CreateUser(context.Background(), nil)
				return err
			}},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				err := tc.fn()
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "not implemented")
			})
		}
	})
}
