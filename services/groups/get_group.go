//go:build ignore

package groups

import (
	"context"
	"fmt"
)

// GetGroup retrieves a group by ID.
func (s *service) GetGroup(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#groups")
}
