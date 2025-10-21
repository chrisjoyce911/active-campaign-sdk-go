package tracking

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTracking_NotImplementedPlaceholders(t *testing.T) {
	t.Run("tracking placeholders", func(t *testing.T) {
		tests := []struct {
			name string
			fn   func() error
		}{
			{"TrackEvent", func() error {
				_, err := (&service{}).TrackEvent(context.Background(), nil)
				return err
			}},
			{"TrackSite", func() error {
				_, err := (&service{}).TrackSite(context.Background(), nil)
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
