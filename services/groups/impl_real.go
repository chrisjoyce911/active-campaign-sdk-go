package groups

import "github.com/chrisjoyce911/active-campaign-sdk-go/client"

// NewRealService returns a concrete GroupsService wired to CoreClient.
func NewRealService(c *client.CoreClient) GroupsService {
	return &service{client: c}
}

// NewRealServiceFromDoer returns a GroupsService backed by any Doer (useful for tests).
func NewRealServiceFromDoer(d client.Doer) GroupsService {
	return &service{client: d}
}
