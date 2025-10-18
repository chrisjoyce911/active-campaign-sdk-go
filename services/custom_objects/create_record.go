//go:build ignore

package custom_objects

import (
	"context"
	"fmt"
)

// CreateObjectRecord creates a record for a custom object type.
func (s *service) CreateObjectRecord(ctx context.Context, objectTypeID string, req interface{}) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#custom-objects")
}
