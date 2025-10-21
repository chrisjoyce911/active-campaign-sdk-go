package messages

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// DeleteMessage deletes a message.
func (s *service) DeleteMessage(ctx context.Context, id string) (*client.APIResponse, error) {
	return nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#messages")
}
