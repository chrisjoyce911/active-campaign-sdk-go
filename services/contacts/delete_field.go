package contacts

import (
	"context"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// DeleteField deletes a custom field by id.
// DEL /fields/{id}
func (s *RealService) DeleteField(ctx context.Context, id string) (*client.APIResponse, error) {
	return s.DeleteCustomField(ctx, id)
}
