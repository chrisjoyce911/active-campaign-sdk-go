//go:build ignore

package webhooks

import (
	"context"
	"fmt"
)

// CreateWebhook registers a webhook.
//
// Docs: https://developers.activecampaign.com/reference#webhooks
func (s *service) CreateWebhook(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#webhooks")
}
