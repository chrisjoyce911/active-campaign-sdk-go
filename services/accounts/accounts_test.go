package accounts

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccounts_NotImplementedPlaceholders(t *testing.T) {
	t.Run("accounts placeholders", func(t *testing.T) {
		tests := []struct {
			name string
			fn   func() error
		}{
			{"CreateAccount", func() error {
				_, _, err := (&service{}).CreateAccount(context.Background(), nil)
				return err
			}},
			{"GetAccount", func() error {
				_, _, err := (&service{}).GetAccount(context.Background(), "1")
				return err
			}},
			{"ListAccounts", func() error {
				_, _, err := (&service{}).ListAccounts(context.Background(), nil)
				return err
			}},
			{"DeleteAccount", func() error {
				_, err := (&service{}).DeleteAccount(context.Background(), "1")
				return err
			}},
			{"BulkDeleteAccounts", func() error {
				_, err := (&service{}).BulkDeleteAccounts(context.Background(), []string{"1", "2"})
				return err
			}},
			{"UpdateAccount", func() error {
				_, _, err := (&service{}).UpdateAccount(context.Background(), "1", nil)
				return err
			}},
			{"CreateAccountNote", func() error {
				_, _, err := (&service{}).CreateAccountNote(context.Background(), "1", nil)
				return err
			}},
			{"UpdateAccountNote", func() error {
				_, _, err := (&service{}).UpdateAccountNote(context.Background(), "1", nil)
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
				assert.Contains(t, err.Error(), "not implemented")
			})
		}
	})
}
