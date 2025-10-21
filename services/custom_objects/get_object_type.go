package custom_objects

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetObjectType retrieves a custom object type by id.
// GET /customObjects/schemas/{id}
func (s *service) GetObjectType(ctx context.Context, id string) (*Schema, *client.APIResponse, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id required")
	}
	if s.client == nil {
		return nil, nil, errors.New("not implemented")
	}
	var wrapper struct {
		Schema Schema `json:"schema"`
	}
	path := "customObjects/schemas/" + id
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &wrapper)
	if err != nil {
		return nil, apiResp, err
	}
	return &wrapper.Schema, apiResp, nil
}
