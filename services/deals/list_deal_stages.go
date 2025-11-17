package deals

import (
	"context"
	"net/http"
	"net/url"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// ListDealStages retrieves all deal stages (GET /api/3/dealStages).
//
// Overview:
//
// 	Returns all existing deal stages with optional query parameters for filtering
// 	and ordering.
//
// Permissions required:
//
// 	- Deal permission: user must have permission to manage deals.
// 	- Pipeline-specific permission: user must have permission to manage the
// 	  pipelines that the stages belong to. Stages from pipelines the user
// 	  cannot manage are not returned.
//
// Query parameters (opts):
//
// 	- filters[title]     : string — Filter by stage title (partial match)
// 	- filters[d_groupid] : string — Filter by pipeline ID
// 	- orders[title]      : ASC|DESC — Order by stage title
//
// Pagination:
//
// 	- limit  : string — Maximum number of records to return (API default if omitted)
// 	- offset : string — Number of records to skip before returning results
//
// Example:
//
// 	opts := map[string]string{
// 		"filters[d_groupid]": "2",
// 		"orders[title]":      "ASC",
// 	}
// 	stages, resp, err := svc.ListDealStages(ctx, opts)
//
// Docs:
//
// 	Reference: https://developers.activecampaign.com/reference/list-all-stages
// 	Example curl:
// 	curl --request GET \
// 	  --url 'https://youraccountname.api-us1.com/api/3/dealStages?filters[d_groupid]=2&orders[title]=ASC' \
// 	  --header 'accept: application/json'
//
// Parameters:
//
// 	ctx  : context.Context
// 	opts : optional query parameters (use exact keys like "filters[d_groupid]", "filters[title]", "orders[title]", "limit", "offset"])
//
// Returns:
//
// 	(*ListDealStagesResponse, *client.APIResponse, error)
func (s *RealService) ListDealStages(ctx context.Context, opts map[string]string) (*ListDealStagesResponse, *client.APIResponse, error) {
	out := &ListDealStagesResponse{}
	base := "dealStages"
	if len(opts) > 0 {
		vals := url.Values{}
		for k, v := range opts {
			vals.Add(k, v)
		}
		base = base + "?" + vals.Encode()
	}
	apiResp, err := s.client.Do(ctx, http.MethodGet, base, nil, out)
	if err != nil {
		return nil, apiResp, err
	}
	return out, apiResp, nil
}
