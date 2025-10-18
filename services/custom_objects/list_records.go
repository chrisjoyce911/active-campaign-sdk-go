//go:build ignore

package custom_objects

import (
	"context"
	"fmt"
)

// ListObjectRecords lists records for a custom object type.
func (s *service) ListObjectRecords(ctx context.Context, objectTypeID string, opts map[string]string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#custom-objects")
}
