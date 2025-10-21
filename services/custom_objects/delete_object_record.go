package custom_objects

import (
	"context"
	"errors"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// DeleteObjectRecord deletes a record by id for a custom object type.
// DELETE /customObjects/schemas/{id}/records/{recordId}
func (s *service) DeleteObjectRecord(ctx context.Context, objectTypeID, recordID string) (*client.APIResponse, error) {
	if s.client == nil {
		return nil, errors.New("not implemented")
	}
	path := "customObjects/records/" + objectTypeID + "/" + recordID
	apiResp, err := s.client.Do(ctx, http.MethodDelete, path, nil, nil)
	return apiResp, err
}
