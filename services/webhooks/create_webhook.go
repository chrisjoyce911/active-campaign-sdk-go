package webhooks

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// CreateWebhook creates a webhook.
//
// What & Why:
//
//	Registers a webhook endpoint to receive ActiveCampaign events.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#create-webhook
//
// Parameters:
//
//	ctx: context
//	req: payload
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *service) CreateWebhook(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#create-webhook")
}
