package deals

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// UpdateDeal updates a deal.
//
// What & Why:
//
//	Modify fields on an existing deal resource.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#update-a-deal
//
// Parameters:
//
//	ctx: context
//	id: deal ID
//	req: updated deal payload
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *RealService) UpdateDeal(ctx context.Context, id string, req interface{}) (interface{}, *client.APIResponse, error) {
	var out interface{}
	path := "deals/" + id
	apiResp, err := s.client.Do(ctx, http.MethodPut, path, req, &out)
	return out, apiResp, err
}
