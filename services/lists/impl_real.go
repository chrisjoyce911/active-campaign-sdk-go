package lists

import "github.com/chrisjoyce911/active-campaign-sdk-go/client"

// NewRealService returns a concrete ListsService wired to CoreClient.
func NewRealService(c *client.CoreClient) ListsService {
	return &service{client: c}
}

// NewRealServiceFromDoer returns a ListsService backed by any Doer (useful for tests).
func NewRealServiceFromDoer(d client.Doer) ListsService {
	return &service{client: d}
}
