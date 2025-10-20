package messages // Ensure package header consistency

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// CreateMessage creates a message resource used by campaigns or automations.
//
// Docs:
//
//	Postman & Reference (see global links)
func (s *service) CreateMessage(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#messages")
}
