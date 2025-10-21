package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// UpdateFieldValueByID updates an existing fieldValue by id.
// PUT /fieldValues/{id}
func (s *RealService) UpdateFieldValueByID(ctx context.Context, id string, req *FieldValuePayload) (*FieldValueResponse, *client.APIResponse, error) {
	out := &FieldValueResponse{}
	path := "fieldValues/" + id
	apiResp, err := s.client.Do(ctx, http.MethodPut, path, req, out)
	return out, apiResp, err
}
