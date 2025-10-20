package webhooks

import (
	"context"
	"fmt"
	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// UpdateWebhook updates an existing webhook.
//
// What & Why:
//
//	Update webhook endpoint or events.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#update-webhook
//
// Parameters:
//
//	ctx: context
//	id: webhook ID
//	req: payload
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *service) UpdateWebhook(ctx context.Context, id string, req interface{}) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#update-webhook")
}
