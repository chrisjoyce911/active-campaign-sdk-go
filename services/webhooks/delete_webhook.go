//go:build ignore

package webhooks

import (
	"context"
	"fmt"
)

// DeleteWebhook deletes a webhook by ID.
func (s *service) DeleteWebhook(ctx context.Context, id string) (*client.APIResponse, error) {
	return nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#webhooks")
}
