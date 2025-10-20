package webhooks

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// DeleteWebhook deletes a webhook by ID.
func (s *service) DeleteWebhook(ctx context.Context, id string) (*client.APIResponse, error) {
	return nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#webhooks")
}
