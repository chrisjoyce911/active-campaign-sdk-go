package custom_objects

import (
	"context"
	"errors"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// CreateObjectRecord creates a new record for a custom object type.
// POST /customObjects/schemas/{id}/records
func (s *service) CreateObjectRecord(ctx context.Context, objectTypeID string, req *CreateRecordRequest) (*CreateRecordResponse, *client.APIResponse, error) {
	if s.client == nil {
		return nil, nil, errors.New("not implemented")
	}
	var out CreateRecordResponse
	// POST to /customObjects/records/{schemaId}
	path := "customObjects/records/" + objectTypeID
	// API expects the payload wrapped with a top-level `record` key
	body := map[string]*CreateRecordRequest{"record": req}
	apiResp, err := s.client.Do(ctx, http.MethodPost, path, body, &out)
	return &out, apiResp, err
}
