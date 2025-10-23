package groups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// DeleteGroup deletes a group.
func (s *service) DeleteGroup(ctx context.Context, id string) (*client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, fmt.Errorf("service not configured: DeleteGroup")
	}
	path := "groups/" + id
	apiResp, err := s.client.Do(ctx, http.MethodDelete, path, nil, nil)
	return apiResp, err
}
