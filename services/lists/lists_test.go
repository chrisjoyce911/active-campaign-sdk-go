package lists

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLists_NotImplementedPlaceholders(t *testing.T) {
	t.Run("lists placeholders", func(t *testing.T) {
		tests := []struct {
			name string
			fn   func() error
		}{
			{"AddContactToList", func() error {
				_, _, err := (&service{}).AddContactToList(context.Background(), nil)
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
