package deals

import (
	"context"
	"net/http"
	"net/url"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

//   - filters[group]               : int32   — Filter by pipeline ID
//   - filters[status]              : int32   — Filter by status
//   - filters[owner]               : int32   — Filter by owner ID
//   - filters[nextdate_range]      : string  — Filter by tasks due date range
//   - filters[tag]                 : string  — Tag names on deal's primary contact
//   - filters[tasktype]            : string  — Deals having tasks of given type
//   - filters[created_before]      : date    — Deals created before given date (YYYY-MM-DD)
//   - filters[created_after]       : date    — Deals created on/after given date (YYYY-MM-DD)
//   - filters[updated_before]      : date    — Deals updated before given date (YYYY-MM-DD)
//   - filters[updated_after]       : date    — Deals updated on/after given date (YYYY-MM-DD)
//   - filters[organization]        : int32   — Primary contact's organization ID
//   - filters[minimum_value]       : int32   — USD dollar portion >= given value
//   - filters[maximum_value]       : int32   — USD dollar portion <= given value
//   - filters[score_greater_than]  : string  — "<score_id>:<score_value>"
//   - filters[score_less_than]     : string  — "<score_id>:<score_value>"
//   - filters[score]               : string  — "<score_id>:<score_value>" (equals)
//
// Ordering (defaults to ASC if not specified):
//
//   - orders[title]           : ASC|DESC — Order by deal title
//   - orders[value]           : ASC|DESC — Order by deal value
//   - orders[cdate]           : ASC|DESC — Order by deal created date
//   - orders[contact_name]    : ASC|DESC — Order by primary contact first name
//   - orders[contact_orgname] : ASC|DESC — Order by primary contact org name
//   - orders[next-action]     : ASC|DESC — Order by next task due date (then id)

// Pagination:
//
//   - limit   : string — Maximum number of records to return (API default applies if omitted)
//   - offset  : string — Number of records to skip before starting to return results
//
// Returns:
//
//	(*ListDealsResponse, *client.APIResponse, error)
func (s *RealService) ListDeals(ctx context.Context, opts map[string]string) (*ListDealsResponse, *client.APIResponse, error) {
	out := &ListDealsResponse{}
	base := "deals"
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
