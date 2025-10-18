//go:build ignore

package custom_objects

import (
	"context"
	"fmt"
)

// DeleteObjectRecord deletes a record for a custom object type.
func (s *service) DeleteObjectRecord(ctx context.Context, objectTypeID, recordID string) (*client.APIResponse, error) {
	return nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#custom-objects")
}
