package webhooks

import (
	"context"
	"fmt"
	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetWebhook retrieves a webhook by ID.
//
// What & Why:
//
//	Fetches webhook details including endpoint and subscribed events.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#get-webhook
//
// Parameters:
//
//	ctx: context
//	id: webhook ID
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *service) GetWebhook(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#get-webhook")
}
