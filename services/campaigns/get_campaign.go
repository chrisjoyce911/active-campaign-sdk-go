package campaigns

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetCampaign retrieves a campaign by ID.
//
// Docs: see Postman and reference links in createCampaign.go
//
// Parameters:
//
//	ctx: context
//	id: campaign ID
//
// Returns: (interface{}, *client.APIResponse, error)
func (s *service) GetCampaign(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#get-campaign")
}
