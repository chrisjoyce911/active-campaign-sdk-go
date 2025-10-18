//go:build ignore

package messages

import (
	"context"
	"fmt"
)

// DeleteMessage deletes a message.
func (s *service) DeleteMessage(ctx context.Context, id string) (*client.APIResponse, error) {
	return nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#messages")
}
