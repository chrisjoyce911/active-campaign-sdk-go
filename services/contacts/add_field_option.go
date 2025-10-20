package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// AddFieldOption adds an option to a field (for select/multiselect types).
// POST /fieldOptions
func (s *RealService) AddFieldOption(ctx context.Context, req *FieldOptionPayload) (*FieldOptionResponse, *client.APIResponse, error) {
	out := &FieldOptionResponse{}
	apiResp, err := s.client.Do(ctx, http.MethodPost, "fieldOptions", req, out)
	return out, apiResp, err
}
