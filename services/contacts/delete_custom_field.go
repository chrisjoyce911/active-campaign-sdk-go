package contacts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// DeleteCustomField deletes a custom field by id.
// DEL /fields/{id}
func (s *RealService) DeleteCustomField(ctx context.Context, id string) (*client.APIResponse, error) {
	path := fmt.Sprintf("fields/%s", id)
	apiResp, err := s.client.Do(ctx, http.MethodDelete, path, nil, nil)
	return apiResp, err
}
