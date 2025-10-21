package custom_objects

import (
	"context"
	"errors"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// ListObjectRecords lists records for a given custom object type.
// GET /customObjects/records/{schemaId}
// Querying supports filter syntax such as filters[relationships.primary-contact][eq]=22
func (s *service) ListObjectRecords(ctx context.Context, objectTypeID string, opts map[string]string) (*ListRecordsResponse, *client.APIResponse, error) {
	if s.client == nil {
		return nil, nil, errors.New("not implemented")
	}
	var out ListRecordsResponse
	path := "customObjects/records/" + objectTypeID
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, opts, &out)
	return &out, apiResp, err
}
