package custom_objects

import (
	"context"
	"errors"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// DeleteObjectType deletes a custom object type by id.
// DELETE /customObjects/schemas/{id}
func (s *service) DeleteObjectType(ctx context.Context, id string) (*client.APIResponse, error) {
	if s.client == nil {
		return nil, errors.New("not implemented")
	}
	path := "customObjects/schemas/" + id
	apiResp, err := s.client.Do(ctx, http.MethodDelete, path, nil, nil)
	return apiResp, err
}
