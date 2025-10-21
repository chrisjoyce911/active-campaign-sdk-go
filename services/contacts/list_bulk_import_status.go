package contacts

import (
	"context"
	"net/http"
	"net/url"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// ListBulkImportStatus lists bulk import statuses.
//
// What & Why:
//
//	Returns a list of bulk import job statuses; useful for monitoring several
//	imports.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#bulk-import-status-list
//
// Parameters:
//
//	ctx: context for cancellation/timeouts
//	opts: optional query parameters
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *RealService) ListBulkImportStatus(ctx context.Context, opts map[string]string) (interface{}, *client.APIResponse, error) {
	var out interface{}
	base := "bulkImport"
	if len(opts) > 0 {
		vals := url.Values{}
		for k, v := range opts {
			vals.Add(k, v)
		}
		base = base + "?" + vals.Encode()
	}
	apiResp, err := s.client.Do(ctx, http.MethodGet, base, nil, &out)
	return out, apiResp, err
}
