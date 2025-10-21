package custom_objects

import (
	"context"
	"errors"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetObjectRecord retrieves a single record by id for a given object type.
// GET /customObjects/schemas/{id}/records/{recordId}
func (s *service) GetObjectRecord(ctx context.Context, objectTypeID, recordID string) (*Record, *client.APIResponse, error) {
	if s.client == nil {
		return nil, nil, errors.New("not implemented")
	}
	var wrapper struct {
		Record Record `json:"record"`
	}
	path := "customObjects/records/" + objectTypeID + "/" + recordID
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &wrapper)
	if err != nil {
		return nil, apiResp, err
	}
	return &wrapper.Record, apiResp, nil
}
