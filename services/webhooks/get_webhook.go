package webhooks

import (
	"context"
	"fmt"
	"path"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetWebhook retrieves a webhook by ID.
func (s *service) GetWebhook(ctx context.Context, id string) (*GetWebhookResponse, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("service not configured: GetWebhook")
	}

	var out GetWebhookResponse
	p := path.Join("webhooks", id)
	apiResp, err := s.client.Do(ctx, "GET", p, nil, &out)
	return &out, apiResp, err
}
