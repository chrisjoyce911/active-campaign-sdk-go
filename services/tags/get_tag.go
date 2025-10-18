//go:build ignore

package tags

import (
	"context"
	"fmt"
)

// GetTag retrieves a tag by ID.
//
// Docs:
//
//	Postman: https://www.postman.com/acdevrel/activecampaign-developer-relations/documentation/ju5a59q/activecampaign-api-v3
//	Reference: https://developers.activecampaign.com/reference#get-tag
//
// Parameters:
//
//	ctx: context
//	id: tag id
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *service) GetTag(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#tags")
}
