package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// BulkImportContacts uploads a bulk import job for contacts.
//
// What & Why:
//
//	Start a bulk import to ingest many contacts.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#bulk-import-contacts
//
// Parameters:
//
//	ctx: context
//	req: request payload
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *RealService) BulkImportContacts(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) {
	var out interface{}
	path := "bulkImport"
	apiResp, err := s.client.Do(ctx, http.MethodPost, path, req, &out)
	return out, apiResp, err
}
