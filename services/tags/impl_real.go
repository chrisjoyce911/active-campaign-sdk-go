package tags

import "github.com/chrisjoyce911/active-campaign-sdk-go/client"

// NewRealService returns a concrete TagsService wired to CoreClient.
func NewRealService(c *client.CoreClient) TagsService {
	return &service{client: c}
}

// NewRealServiceFromDoer returns a TagsService backed by any Doer (useful for tests).
func NewRealServiceFromDoer(d client.Doer) TagsService {
	return &service{client: d}
}
