package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// ListCustomFields lists all custom fields for contacts.
// GET /fields
func (s *RealService) ListCustomFields(ctx context.Context) (*ListFieldsResponse, *client.APIResponse, error) {
	return s.ListCustomFieldsWithOpts(ctx, nil)
}

// ListCustomFieldsWithOpts lists custom fields with optional query parameters
// e.g. opts["limit"] = "100". This is a non-breaking addition; existing
// callers can continue to call ListCustomFields(ctx) which delegates here.
func (s *RealService) ListCustomFieldsWithOpts(ctx context.Context, opts map[string]string) (*ListFieldsResponse, *client.APIResponse, error) {
	out := &ListFieldsResponse{}
	apiResp, err := s.client.Do(ctx, http.MethodGet, "fields", opts, out)
	return out, apiResp, err
}
