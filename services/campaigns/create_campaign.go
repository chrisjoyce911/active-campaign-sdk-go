package campaigns

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// CreateCampaign creates a new campaign resource.
//
// What & Why:
//
//	Campaigns represent email campaigns to be sent. Creating a campaign sets up
//	the content, settings and recipient lists.
//
// Docs:
//
//	Postman: https://www.postman.com/acdevrel/activecampaign-developer-relations/documentation/ju5a59q/activecampaign-api-v3
//	Reference: https://developers.activecampaign.com/reference#create-campaign
//
// Parameters:
//
//	ctx: context for cancellation
//	req: campaign create payload
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *service) CreateCampaign(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#create-campaign")
}
