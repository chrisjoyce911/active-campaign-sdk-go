//go:build ignore

package webhooks

import (
	"context"
	"fmt"
)

// UpdateWebhook updates a webhook.
func (s *service) UpdateWebhook(ctx context.Context, id string, req interface{}) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#webhooks")
}
