//go:build ignore

package tags

import "context"

// CreateTag creates a new tag resource.
//
// What & Why:
//
//	Tags are used to categorise contacts and trigger automations. Use this to
//	create and manage tags programmatically.
//
// Docs:
//
//	Postman: https://www.postman.com/acdevrel/activecampaign-developer-relations/documentation/ju5a59q/activecampaign-api-v3
//	Reference: https://developers.activecampaign.com/reference#create-tag
//
// Parameters:
//
//	ctx: context
//	req: payload
//
// Returns:
//
//	(*CreateTagResponse, *client.APIResponse, error)
func (s *service) CreateTag(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#tags")
}
}
