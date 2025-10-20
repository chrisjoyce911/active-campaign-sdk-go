package deals

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// CreateDeal creates a deal record.
//
// What & Why:
//
//	Create a new deal within ActiveCampaign's CRM.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#create-a-deal
//
// Parameters:
//
//	ctx: context
//	req: deal payload
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *RealService) CreateDeal(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) {
	var out interface{}
	apiResp, err := s.client.Do(ctx, http.MethodPost, "deals", req, &out)
	return out, apiResp, err
}
