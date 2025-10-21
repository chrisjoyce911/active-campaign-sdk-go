//go:build ignore

package users

import (
	"context"
	"fmt"
)

// ListUsers lists users in the account.
func (s *service) ListUsers(ctx context.Context, opts map[string]string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#users")
}
