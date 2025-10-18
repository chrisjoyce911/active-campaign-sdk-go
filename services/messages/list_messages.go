//go:build ignore

package messages

import (
	"context"
	"fmt"
)

// ListMessages lists messages.
func (s *service) ListMessages(ctx context.Context, opts map[string]string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#messages")
}
