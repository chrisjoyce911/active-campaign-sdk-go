package campaigns

import "github.com/chrisjoyce911/active-campaign-sdk-go/client"

// NewRealService returns a concrete CampaignsService wired to CoreClient.
func NewRealService(c *client.CoreClient) CampaignsService {
	return &service{client: c}
}

// NewRealServiceFromDoer returns a CampaignsService backed by any Doer (useful for tests).
func NewRealServiceFromDoer(d client.Doer) CampaignsService {
	return &service{client: d}
}
