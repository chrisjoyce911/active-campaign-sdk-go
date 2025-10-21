package messages

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetMessage retrieves a message by ID.
func (s *service) GetMessage(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#messages")
}
