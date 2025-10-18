package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// RealService is a minimal concrete implementation of ContactsService that
// uses client.CoreClient to call the ActiveCampaign API.
type RealService struct {
	client *client.CoreClient
}

func NewRealService(c *client.CoreClient) ContactsService {
	return &RealService{client: c}
}

func (s *RealService) Create(ctx context.Context, req *CreateContactRequest) (*CreateContactResponse, *client.APIResponse, error) {
	out := &CreateContactResponse{}
	apiResp, err := s.client.Do(ctx, http.MethodPost, "contacts", req, out)
	return out, apiResp, err
}

func (s *RealService) SearchByEmail(ctx context.Context, email string) (*ContactSearchResponse, *client.APIResponse, error) {
	out := &ContactSearchResponse{}
	path := client.BuildContactsSearchPath(email)
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, out)
	return out, apiResp, err
}

func (s *RealService) UpdateListStatus(ctx context.Context, req *UpdateListStatusForContactRequest) (*UpdateContactListStatusResponse, *client.APIResponse, error) {
	out := &UpdateContactListStatusResponse{}
	apiResp, err := s.client.Do(ctx, http.MethodPost, "contactLists", req, out)
	return out, apiResp, err
}
