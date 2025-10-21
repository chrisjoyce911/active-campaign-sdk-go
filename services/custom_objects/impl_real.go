package custom_objects

import "github.com/chrisjoyce911/active-campaign-sdk-go/client"

// NewRealService returns a concrete CustomObjectsService wired to CoreClient.
func NewRealService(c *client.CoreClient) CustomObjectsService {
	return &service{client: c}
}

// NewRealServiceFromDoer returns a CustomObjectsService backed by any Doer (useful for tests).
func NewRealServiceFromDoer(d client.Doer) CustomObjectsService {
	return &service{client: d}
}
