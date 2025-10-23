package campaigns

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetCampaignLinks retrieves links associated with the given campaign id.
func (s *service) GetCampaignLinks(ctx context.Context, id string) (*CampaignLinksResponse, *client.APIResponse, error) {
	var out CampaignLinksResponse
	path := "campaigns/" + id + "/links"
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &out)
	if err != nil {
		return nil, apiResp, err
	}
	return &out, apiResp, nil
}

// CampaignLinks is a convenience helper that fetches campaign links and
// returns the slice of links; if messageID is non-nil it filters results to
// only links whose Message matches the provided messageID.
func (s *service) CampaignLinks(ctx context.Context, id string, messageID *string) ([]CampaignLink, *client.APIResponse, error) {
	out, apiResp, err := s.GetCampaignLinks(ctx, id)
	if err != nil {
		return nil, apiResp, err
	}
	if out == nil {
		return nil, apiResp, nil
	}
	if messageID == nil {
		return out.Links, apiResp, nil
	}
	var filtered []CampaignLink
	for _, l := range out.Links {
		if l.Message == nil {
			continue
		}
		if *l.Message == *messageID {
			filtered = append(filtered, l)
		}
	}
	return filtered, apiResp, nil
}
