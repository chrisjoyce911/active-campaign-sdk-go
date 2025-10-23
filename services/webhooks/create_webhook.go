package webhooks

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// CreateWebhook creates a webhook.
func (s *service) CreateWebhook(ctx context.Context, req *CreateWebhookRequest) (*CreateWebhookResponse, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("service not configured: CreateWebhook")
	}

	var out CreateWebhookResponse
	apiResp, err := s.client.Do(ctx, "POST", "webhooks", req, &out)
	return &out, apiResp, err
}
