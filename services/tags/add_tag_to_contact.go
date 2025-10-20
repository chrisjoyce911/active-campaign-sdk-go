package tags

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// AddTagToContact associates a tag with a contact.
//
// It sends a POST to /contacts/{contactID}/tags with a CreateOrUpdateTagRequest
// payload. The response is returned as a TagResponse and API metadata.
// A nil receiver or missing client returns a not-implemented error which
// helps during migration and tests that exercise zero-value receivers.
// Path: POST contacts/{contactID}/tags
func (s *service) AddTagToContact(ctx context.Context, contactID string, req *CreateOrUpdateTagRequest) (*TagResponse, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#create-contact-tag")
	}
	var out TagResponse
	path := "contacts/" + contactID + "/tags"
	apiResp, err := s.client.Do(ctx, http.MethodPost, path, req, &out)
	return &out, apiResp, err
}
