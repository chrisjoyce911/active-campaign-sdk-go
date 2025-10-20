package contacts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// UpdateCustomField updates an existing custom field.
// PUT /fields/{id}
func (s *RealService) UpdateCustomField(ctx context.Context, id string, req *FieldPayload) (*FieldResponse, *client.APIResponse, error) {
	out := &FieldResponse{}
	path := fmt.Sprintf("fields/%s", id)
	apiResp, err := s.client.Do(ctx, http.MethodPut, path, req, out)
	return out, apiResp, err
}
