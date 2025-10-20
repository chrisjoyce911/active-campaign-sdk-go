package messages

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// SendMessage triggers sending a message (preview or send) depending on API.
func (s *service) SendMessage(ctx context.Context, id string, opts map[string]interface{}) (*client.APIResponse, error) {
	return nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#messages")
}
