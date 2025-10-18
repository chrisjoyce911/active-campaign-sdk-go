//go:build ignore

package messages

import (
	"context"
	"fmt"
)

// UpdateMessage updates a message.
func (s *service) UpdateMessage(ctx context.Context, id string, req interface{}) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#messages")
}
