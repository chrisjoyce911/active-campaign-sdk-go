package webhooks

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// ListWebhooks lists all webhooks with optional filters.
func (s *service) ListWebhooks(ctx context.Context, opts map[string]string) (*ListWebhooksResponse, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("service not configured: ListWebhooks")
	}

	var out ListWebhooksResponse
	apiResp, err := s.client.Do(ctx, "GET", "webhooks", opts, &out)
	return &out, apiResp, err
}

// ListWebhookEvents lists webhook events with optional filters.
func (s *service) ListWebhookEvents(ctx context.Context, opts map[string]string) (*ListWebhookEventsResponse, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("service not configured: ListWebhookEvents")
	}

	var out ListWebhookEventsResponse
	apiResp, err := s.client.Do(ctx, "GET", "webhook/events", opts, &out)
	return &out, apiResp, err
}
