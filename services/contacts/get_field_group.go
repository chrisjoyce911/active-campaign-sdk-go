package contacts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetFieldGroup retrieves a field group by id and returns a typed response.
// GET /fieldGroups/{id}
func (s *RealService) GetFieldGroup(ctx context.Context, id string) (*FieldGroupResponse, *client.APIResponse, error) {
	path := fmt.Sprintf("fieldGroups/%s", id)
	out := &FieldGroupResponse{}
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, out)
	return out, apiResp, err
}
