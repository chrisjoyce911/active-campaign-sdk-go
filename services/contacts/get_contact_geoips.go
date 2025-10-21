package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetContactGeoIPs returns geo IP information for a contact.
//
// What & Why:
//
//	Provide IP-based geo information for analytics or enrichment.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#get-contact-geoips
//
// Parameters:
//
//	ctx: context
//	id: contact id
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *RealService) GetContactGeoIPs(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	var out interface{}
	path := "contacts/" + id + "/geoips"
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &out)
	return out, apiResp, err
}
