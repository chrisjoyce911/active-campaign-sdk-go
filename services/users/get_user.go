//go:build ignore

package users

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetUser retrieves a user by ID.
func (s *service) GetUser(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#users")
}
