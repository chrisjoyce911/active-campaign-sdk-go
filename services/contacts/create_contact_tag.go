package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// ContactTagPayload describes the body for creating a contactTag
type ContactTagPayload struct {
	Contact string `json:"contact"`
	Tag     string `json:"tag"`
}

// ContactTagRequest is the envelope to create a contactTag association
type ContactTagRequest struct {
	ContactTag ContactTagPayload `json:"contactTag"`
}

// ContactTagResponse is the envelope returned by POST /contactTags
type ContactTagResponse struct {
	ContactTag ContactTagPayload `json:"contactTag"`
}

// CreateContactTag creates a contactTag association (POST /contactTags)
// Implemented on RealService so callers can use the contacts service for
// contact-related operations instead of reaching into tags or core.
func (s *RealService) CreateContactTag(ctx context.Context, req *ContactTagRequest) (*ContactTagResponse, *client.APIResponse, error) {
	out := &ContactTagResponse{}
	apiResp, err := s.client.Do(ctx, http.MethodPost, "contactTags", req, out)
	return out, apiResp, err
}
