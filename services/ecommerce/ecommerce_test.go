package ecommerce

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEcommerce_NotImplementedPlaceholders(t *testing.T) {
	t.Run("ecommerce placeholders", func(t *testing.T) {
		tests := []struct {
			name string
			fn   func() error
		}{
			{"CreateOrder", func() error {
				_, _, err := (&service{}).CreateOrder(context.Background(), nil)
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
