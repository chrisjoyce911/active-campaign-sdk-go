//go:build ignore

package custom_objects

import (
	"context"
	"fmt"
)

// GetObjectType retrieves a custom object type.
func (s *service) GetObjectType(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#custom-objects")
}
