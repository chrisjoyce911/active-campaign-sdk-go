package contacts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// UpdateFieldGroup updates a field group by id.
// PUT /fieldGroups/{id}
func (s *RealService) UpdateFieldGroup(ctx context.Context, id string, req interface{}) (*client.APIResponse, error) {
	path := fmt.Sprintf("fieldGroups/%s", id)
	apiResp, err := s.client.Do(ctx, http.MethodPut, path, req, nil)
	return apiResp, err
}
