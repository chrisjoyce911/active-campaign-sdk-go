package messages

import "github.com/chrisjoyce911/active-campaign-sdk-go/client"

// NewRealService returns a concrete MessagesService wired to CoreClient.
func NewRealService(c *client.CoreClient) MessagesService {
	return &service{client: c}
}

// NewRealServiceFromDoer returns a MessagesService backed by any Doer (useful for tests).
func NewRealServiceFromDoer(d client.Doer) MessagesService {
	return &service{client: d}
}
