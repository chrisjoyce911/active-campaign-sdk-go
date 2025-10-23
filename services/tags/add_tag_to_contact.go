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
// When the service is not configured this method returns an error indicating
// the service is not configured.
//
// API reference: https://developers.activecampaign.com/reference#create-contact-tag
// Path: POST contacts/{contactID}/tags
func (s *service) AddTagToContact(ctx context.Context, contactID string, req *CreateOrUpdateTagRequest) (*TagResponse, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("service not configured: AddTagToContact")
	}
	var out TagResponse
	path := "contacts/" + contactID + "/tags"
	apiResp, err := s.client.Do(ctx, http.MethodPost, path, req, &out)
	return &out, apiResp, err
}
