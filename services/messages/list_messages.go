package messages

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// ListMessages lists messages.
func (s *service) ListMessages(ctx context.Context, opts map[string]string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#messages")
}
