package messages

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMessages_NotImplementedPlaceholders(t *testing.T) {
	t.Run("messages placeholders", func(t *testing.T) {
		tests := []struct {
			name string
			fn   func() error
		}{
			{"CreateMessage", func() error {
				_, _, err := (&service{}).CreateMessage(context.Background(), nil)
				return err
			}},
			{"GetMessage", func() error {
				_, _, err := (&service{}).GetMessage(context.Background(), "1")
				return err
			}},
			{"ListMessages", func() error {
				_, _, err := (&service{}).ListMessages(context.Background(), nil)
				return err
			}},
			{"DeleteMessage", func() error {
				_, err := (&service{}).DeleteMessage(context.Background(), "1")
				return err
			}},
			{"SendMessage", func() error {
				_, err := (&service{}).SendMessage(context.Background(), "1", nil)
				return err
			}},
			{"UpdateMessage", func() error {
				_, _, err := (&service{}).UpdateMessage(context.Background(), "1", nil)
				return err
			}},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				require := require.New(t)

				svc := &service{}
				require.NotNil(svc)

				// execute the test function using the prebuilt service where appropriate
				err := tc.fn()
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "not implemented")
			})
		}
	})
}
