package ecommerce

import "github.com/chrisjoyce911/active-campaign-sdk-go/client"

// NewRealService returns a concrete EcommerceService wired to CoreClient.
func NewRealService(c *client.CoreClient) EcommerceService {
	return &service{client: c}
}

// NewRealServiceFromDoer returns an EcommerceService backed by any Doer (useful for tests).
func NewRealServiceFromDoer(d client.Doer) EcommerceService {
	return &service{client: d}
}
