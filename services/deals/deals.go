package deals

import (
	"context"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// DealsService defines the public surface for the deals service.
// It matches the methods implemented by RealService and lets
// callers depend on the interface rather than the concrete type.
type DealsService interface {
	CreateDeal(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error)
	CreateDealNote(ctx context.Context, dealID string, req interface{}) (interface{}, *client.APIResponse, error)
	GetDeal(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	GetDealActivities(ctx context.Context, dealID string) (interface{}, *client.APIResponse, error)
	ListDeals(ctx context.Context, opts map[string]string) (*ListDealsResponse, *client.APIResponse, error)
	ListDealStages(ctx context.Context, opts map[string]string) (*ListDealStagesResponse, *client.APIResponse, error)
	UpdateDeal(ctx context.Context, id string, req interface{}) (interface{}, *client.APIResponse, error)
	UpdateDealNote(ctx context.Context, dealID, noteID string, req interface{}) (interface{}, *client.APIResponse, error)
	BulkUpdateDealOwners(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error)
	DeleteDeal(ctx context.Context, id string) (*client.APIResponse, error)
}
