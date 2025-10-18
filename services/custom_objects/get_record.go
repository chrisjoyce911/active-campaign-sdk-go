//go:build ignore

package custom_objects

import (
	"context"
	"fmt"
)

// GetObjectRecord retrieves a custom object record by ID.
func (s *service) GetObjectRecord(ctx context.Context, objectTypeID, recordID string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#custom-objects")
}
