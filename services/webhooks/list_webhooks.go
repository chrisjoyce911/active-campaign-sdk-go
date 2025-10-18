//go:build ignore

package webhooks

import (
	"context"
	"fmt"
)

// ListWebhooks lists registered webhooks.
func (s *service) ListWebhooks(ctx context.Context, opts map[string]string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#webhooks")
}
