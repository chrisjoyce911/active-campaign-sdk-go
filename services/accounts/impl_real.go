package accounts

import "github.com/chrisjoyce911/active-campaign-sdk-go/client"

// NewRealService returns a concrete service implementation wired to CoreClient.
func NewRealService(c *client.CoreClient) AccountsService {
	return &service{client: c}
}

// NewRealServiceFromDoer returns a service backed by any Doer (useful for tests).
func NewRealServiceFromDoer(d client.Doer) AccountsService {
	return &service{client: d}
}
