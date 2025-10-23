package webhooks

import (
	"context"
	"fmt"
	"path"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// DeleteWebhook deletes a webhook by ID.
func (s *service) DeleteWebhook(ctx context.Context, id string) (*client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, fmt.Errorf("service not configured: DeleteWebhook")
	}

	p := path.Join("webhooks", id)
	apiResp, err := s.client.Do(ctx, "DELETE", p, nil, nil)
	return apiResp, err
}
