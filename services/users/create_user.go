package users

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// CreateUser creates a user (account user) in ActiveCampaign.
func (s *service) CreateUser(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#users")
}
