package tracking

import "github.com/chrisjoyce911/active-campaign-sdk-go/client"

// NewRealService returns a concrete TrackingService wired to CoreClient.
func NewRealService(c *client.CoreClient) TrackingService {
	return &service{client: c}
}

// NewRealServiceFromDoer returns a TrackingService backed by any Doer (useful for tests).
func NewRealServiceFromDoer(d client.Doer) TrackingService {
	return &service{client: d}
}
