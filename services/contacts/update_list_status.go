package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

func (s *RealService) UpdateListStatus(ctx context.Context, req *UpdateListStatusForContactRequest) (*UpdateContactListStatusResponse, *client.APIResponse, error) {
	out := &UpdateContactListStatusResponse{}
	apiResp, err := s.client.Do(ctx, http.MethodPost, "contactLists", req, out)
	return out, apiResp, err
}
