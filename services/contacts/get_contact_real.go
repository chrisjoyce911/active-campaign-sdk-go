package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

func (s *RealService) GetContact(ctx context.Context, id string) (*CreateContactResponse, *client.APIResponse, error) {
	out := &CreateContactResponse{}
	path := "contacts/" + id
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, out)
	return out, apiResp, err
}
