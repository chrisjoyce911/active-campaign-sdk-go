package deals

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetDealActivities retrieves activities for a given deal.
//
// What & Why:
//
//	Fetch the activity feed (timeline) related to a specific deal.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#get-deal-activities
//
// Parameters:
//
//	ctx: context
//	dealID: ID of the deal
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *RealService) GetDealActivities(ctx context.Context, dealID string) (interface{}, *client.APIResponse, error) {
	var out interface{}
	path := "deals/" + dealID + "/activities"
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &out)
	return out, apiResp, err
}
