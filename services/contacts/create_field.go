package contacts

import (
	"context"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// CreateField creates a custom field for contacts.
// POST /fields
func (s *RealService) CreateField(ctx context.Context, req *FieldPayload) (*FieldResponse, *client.APIResponse, error) {
	return s.CreateCustomField(ctx, req)
}
