package contactautomation

import "github.com/chrisjoyce911/active-campaign-sdk-go/client"

// NewRealService returns a concrete ContactAutomationService wired to CoreClient.
func NewRealService(c *client.CoreClient) ContactAutomationService {
	return &service{client: c}
}

// NewRealServiceFromDoer returns a ContactAutomationService backed by any Doer (useful for tests).
func NewRealServiceFromDoer(d client.Doer) ContactAutomationService {
	return &service{client: d}
}
