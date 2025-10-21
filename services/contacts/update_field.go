package contacts

import (
	"context"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// UpdateField updates an existing custom field.
// PUT /fields/{id}
func (s *RealService) UpdateField(ctx context.Context, id string, req *FieldPayload) (*FieldResponse, *client.APIResponse, error) {
	return s.UpdateCustomField(ctx, id, req)
}
