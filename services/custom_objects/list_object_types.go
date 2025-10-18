//go:build ignore

package custom_objects

import (
	"context"
	"fmt"
)

// ListObjectTypes lists available custom object types.
func (s *service) ListObjectTypes(ctx context.Context, opts map[string]string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#custom-objects")
}
