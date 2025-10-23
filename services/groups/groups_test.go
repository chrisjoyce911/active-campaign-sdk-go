package groups

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroups_NotImplementedPlaceholders(t *testing.T) {
	t.Run("groups placeholders", func(t *testing.T) {
		tests := []struct {
			name string
			fn   func() error
		}{
			{"CreateGroup", func() error {
				_, _, err := (&service{}).CreateGroup(context.Background(), nil)
				return err
			}},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				err := tc.fn()
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "service not configured")
			})
		}
	})
}
