package contacts

import (
	"context"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// CreateContactWithTags creates a contact and attaches the provided tag IDs.
// It first calls Create; if that succeeds it attempts to create contactTag
// associations for each tag ID. If any attach call fails, the method returns
// the created contact along with the last attach error and its APIResponse.
func (s *RealService) CreateContactWithTags(ctx context.Context, req *CreateContactRequest, tagIDs []string) (*CreateContactResponse, *client.APIResponse, error) {
	created, apiResp, err := s.Create(ctx, req)
	if err != nil {
		return created, apiResp, err
	}
	if created == nil || created.Contact == nil {
		return created, apiResp, nil
	}

	var lastErr error
	var lastAPIResp *client.APIResponse
	for _, tid := range tagIDs {
		ctReq := &ContactTagRequest{ContactTag: ContactTagPayload{Contact: created.Contact.ID, Tag: tid}}
		_, resp, err := s.CreateContactTag(ctx, ctReq)
		if err != nil {
			lastErr = err
			lastAPIResp = resp
			// continue attempting remaining tags
		}
	}
	return created, lastAPIResp, lastErr
}
