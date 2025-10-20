package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetBulkImportStatus returns the status of a bulk import job by id.
//
// What & Why:
//
//	Allow callers to query the state of an import job started via BulkImportContacts.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#bulk-import-status
//
// Parameters:
//
//	ctx: context
//	id: bulk import job id
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *RealService) GetBulkImportStatus(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	var out interface{}
	path := "bulkImport/" + id
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &out)
	return out, apiResp, err
}
