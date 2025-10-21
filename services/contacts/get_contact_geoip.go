package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetContactGeoIP returns a specific geo IP record for a contact.
//
// What & Why:
//
//	Provide a single geo IP record by IP address for inspection or debugging.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#get-contact-geoip
//
// Parameters:
//
//	ctx: context
//	id: contact id
//	ip: ip address
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *RealService) GetContactGeoIP(ctx context.Context, id, ip string) (interface{}, *client.APIResponse, error) {
	var out interface{}
	path := "contacts/" + id + "/geoips/" + ip
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &out)
	return out, apiResp, err
}
