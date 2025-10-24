package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

func (s *RealService) Create(ctx context.Context, req *CreateContactRequest) (*CreateContactResponse, *client.APIResponse, error) {
	out := &CreateContactResponse{}
	apiResp, err := s.client.Do(ctx, http.MethodPost, "contacts", req, out)
	return out, apiResp, err
}
