package webhooks

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebhooks_NotImplementedPlaceholders(t *testing.T) {
	t.Run("webhooks placeholders", func(t *testing.T) {
		tests := []struct {
			name string
			fn   func() error
		}{
			{"CreateWebhook", func() error {
				_, _, err := (&service{}).CreateWebhook(context.Background(), nil)
				return err
			}},
			{"GetWebhook", func() error {
				_, _, err := (&service{}).GetWebhook(context.Background(), "1")
				return err
			}},
			{"ListWebhooks", func() error {
				_, _, err := (&service{}).ListWebhooks(context.Background(), nil)
				return err
			}},
			{"DeleteWebhook", func() error {
				_, err := (&service{}).DeleteWebhook(context.Background(), "1")
				return err
			}},
			{"UpdateWebhook", func() error {
				_, _, err := (&service{}).UpdateWebhook(context.Background(), "1", nil)
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
