package webhooks

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// ListWebhooks lists all webhooks with optional filters.
//
// What & Why:
//
//	Returns a paginated list of registered webhooks.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#list-webhooks
//
// Parameters:
//
//	ctx: context
//	opts: query options
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *service) ListWebhooks(ctx context.Context, opts map[string]string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#list-webhooks")
}

// ListWebhookEvents lists webhook events with optional filters.
//
// What & Why:
//
//	Returns a list of webhook events sent to endpoints for debugging/audit.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#list-webhook-events
//
// Parameters:
//
//	ctx: context
//	opts: query options
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *service) ListWebhookEvents(ctx context.Context, opts map[string]string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#list-webhook-events")
}
