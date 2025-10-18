//go:build ignore

package custom_objects

import (
	"context"
	"fmt"
)

// UpdateObjectRecord updates a custom object record.
func (s *service) UpdateObjectRecord(ctx context.Context, objectTypeID, recordID string, req interface{}) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#custom-objects")
}
