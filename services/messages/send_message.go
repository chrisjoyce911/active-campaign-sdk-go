//go:build ignore

package messages

import (
	"context"
	"fmt"
)

// SendMessage triggers sending a message (preview or send) depending on API.
func (s *service) SendMessage(ctx context.Context, id string, opts map[string]interface{}) (*client.APIResponse, error) {
	return nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#messages")
}
