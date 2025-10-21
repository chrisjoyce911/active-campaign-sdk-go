package custom_objects

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// CreateObjectType creates a custom object type (schema).
// POST /customObjects/schemas
func (s *service) CreateObjectType(ctx context.Context, req *CreateObjectTypeRequest) (*Schema, *client.APIResponse, error) {
	var out CreateObjectTypeResponse
	// API expects the payload wrapped with a top-level `schema` key
	body := map[string]*CreateObjectTypeRequest{"schema": req}
	apiResp, err := s.client.Do(ctx, http.MethodPost, "customObjects/schemas", body, &out)
	if err != nil {
		return nil, apiResp, err
	}
	return &out.Schema, apiResp, nil
}
