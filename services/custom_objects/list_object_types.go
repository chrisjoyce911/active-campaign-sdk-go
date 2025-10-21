package custom_objects

import (
	"context"
	"errors"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// ListObjectTypes lists all schemas (custom object types) available within an account.
//
// https://developers.activecampaign.com/reference/list-all-schemas
// API: GET /customObjects/schemas
//
// Note: callers may pass query opts such as `showFields=all` to include full
// field definitions which is useful when programmatically creating or
// synchronizing schema fields (it helps ensure field id uniqueness). Fields
// that were deleted may include a `status: marked_for_deletion` attribute.
func (s *service) ListObjectTypes(ctx context.Context, opts map[string]string) (*ListSchemasResponse, *client.APIResponse, error) {
	if s.client == nil {
		return nil, nil, errors.New("not implemented")
	}
	var out ListSchemasResponse
	apiResp, err := s.client.Do(ctx, http.MethodGet, "customObjects/schemas", opts, &out)
	return &out, apiResp, err
}
