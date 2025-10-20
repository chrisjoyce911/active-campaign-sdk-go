package deals

import (
	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// RealService is the concrete implementation for deals.
type RealService struct {
	client client.Doer
}

// NewRealService returns a DealsService wired to CoreClient.
func NewRealService(c *client.CoreClient) DealsService {
	return &RealService{client: c}
}

// NewRealServiceFromDoer returns a DealsService backed by any Doer (for tests).
func NewRealServiceFromDoer(d client.Doer) DealsService {
	return &RealService{client: d}
}
