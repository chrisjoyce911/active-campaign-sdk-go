package webhooks

import (
	"context"
	"fmt"
	"path"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// UpdateWebhook updates an existing webhook.
func (s *service) UpdateWebhook(ctx context.Context, id string, req *UpdateWebhookRequest) (*UpdateWebhookResponse, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("service not configured: UpdateWebhook")
	}

	var out UpdateWebhookResponse
	p := path.Join("webhooks", id)
	apiResp, err := s.client.Do(ctx, "PUT", p, req, &out)
	return &out, apiResp, err
}
