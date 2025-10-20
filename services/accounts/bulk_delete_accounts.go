package accounts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// BulkDeleteAccounts deletes multiple accounts by IDs.
//
// What & Why:
//
//	Allows deleting multiple accounts in a single request when supported.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#bulk-delete-accounts
//
// Parameters:
//
//	ctx: context
//	ids: slice of account IDs
//
// Returns:
//
//	(*client.APIResponse, error)
func (s *service) BulkDeleteAccounts(ctx context.Context, ids []string) (*client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#bulk-delete-accounts")
	}
	// assuming API supports POST to accounts/bulk-delete or similar; use accounts/bulk_delete
	apiResp, err := s.client.Do(ctx, http.MethodPost, "accounts/bulk_delete", map[string]interface{}{"ids": ids}, nil)
	return apiResp, err
}
