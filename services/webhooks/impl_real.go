package webhooks

import "github.com/chrisjoyce911/active-campaign-sdk-go/client"

// NewRealService returns a concrete WebhooksService wired to CoreClient.
func NewRealService(c *client.CoreClient) WebhooksService {
	return &service{client: c}
}

// NewRealServiceFromDoer returns a WebhooksService backed by any Doer (useful for tests).
func NewRealServiceFromDoer(d client.Doer) WebhooksService {
	return &service{client: d}
}
