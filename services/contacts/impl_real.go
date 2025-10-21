package contacts

import (
	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// RealService is a minimal concrete implementation of ContactsService that
// uses client.CoreClient to call the ActiveCampaign API.
type RealService struct {
	client client.Doer
}

// NewRealService returns a RealService backed by a CoreClient.
// NOTE: returns the concrete *RealService during the migration so callers
// that need to assert the concrete type in tests continue to work. Once the
// migration completes we can switch back to returning the ContactsService
// interface if desired.
func NewRealService(c *client.CoreClient) *RealService {
	return &RealService{client: c}
}

// NewRealServiceFromDoer returns a RealService backed by any Doer (useful for tests).
func NewRealServiceFromDoer(d client.Doer) *RealService {
	return &RealService{client: d}
}

// NOTE: endpoint implementations were moved into individual files under
// services/contacts/*. Each endpoint now has its own source and test file to
// simplify incremental migration and clearer diffs.
