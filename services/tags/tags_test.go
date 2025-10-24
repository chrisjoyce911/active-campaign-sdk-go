package tags

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTags_NotImplementedPlaceholders(t *testing.T) {
	t.Run("tags placeholders", func(t *testing.T) {
		tests := []struct {
			name string
			fn   func() error
		}{
			{"CreateTag", func() error {
				_, _, err := (&service{}).CreateTag(context.Background(), nil)
				return err
			}},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				require := require.New(t)

				svc := &service{}
				require.NotNil(svc)

				err := tc.fn()
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "service not configured")
			})
		}
	})
}
