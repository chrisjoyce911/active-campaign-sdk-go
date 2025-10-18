//go:build ignore

package webhooks

import (
	"context"
	"fmt"
)

// GetWebhook retrieves a webhook by ID.
func (s *service) GetWebhook(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#webhooks")
}
