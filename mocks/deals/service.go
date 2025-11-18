package dealsmock

import (
	"context"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/deals"
)

// Service is a function-field mock that implements deals.DealsService.
// Tests can set the desired function fields per test case.
type Service struct {
	CreateDealFunc           func(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error)
	CreateDealNoteFunc       func(ctx context.Context, dealID string, req interface{}) (interface{}, *client.APIResponse, error)
	GetDealFunc              func(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	GetDealActivitiesFunc    func(ctx context.Context, dealID string) (interface{}, *client.APIResponse, error)
	ListDealsFunc            func(ctx context.Context, opts map[string]string) (*deals.ListDealsResponse, *client.APIResponse, error)
	ListDealStagesFunc       func(ctx context.Context, opts map[string]string) (*deals.ListDealStagesResponse, *client.APIResponse, error)
	UpdateDealFunc           func(ctx context.Context, id string, req interface{}) (interface{}, *client.APIResponse, error)
	UpdateDealNoteFunc       func(ctx context.Context, dealID, noteID string, req interface{}) (interface{}, *client.APIResponse, error)
	BulkUpdateDealOwnersFunc func(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error)
	DeleteDealFunc           func(ctx context.Context, id string) (*client.APIResponse, error)
}

var _ deals.DealsService = (*Service)(nil)

func (m *Service) CreateDeal(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) {
	if m.CreateDealFunc != nil {
		return m.CreateDealFunc(ctx, req)
	}
	return nil, nil, nil
}
func (m *Service) CreateDealNote(ctx context.Context, dealID string, req interface{}) (interface{}, *client.APIResponse, error) {
	if m.CreateDealNoteFunc != nil {
		return m.CreateDealNoteFunc(ctx, dealID, req)
	}
	return nil, nil, nil
}
func (m *Service) GetDeal(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	if m.GetDealFunc != nil {
		return m.GetDealFunc(ctx, id)
	}
	return nil, nil, nil
}
func (m *Service) GetDealActivities(ctx context.Context, dealID string) (interface{}, *client.APIResponse, error) {
	if m.GetDealActivitiesFunc != nil {
		return m.GetDealActivitiesFunc(ctx, dealID)
	}
	return nil, nil, nil
}
func (m *Service) ListDeals(ctx context.Context, opts map[string]string) (*deals.ListDealsResponse, *client.APIResponse, error) {
	if m.ListDealsFunc != nil {
		return m.ListDealsFunc(ctx, opts)
	}
	return &deals.ListDealsResponse{}, &client.APIResponse{}, nil
}
func (m *Service) ListDealStages(ctx context.Context, opts map[string]string) (*deals.ListDealStagesResponse, *client.APIResponse, error) {
	if m.ListDealStagesFunc != nil {
		return m.ListDealStagesFunc(ctx, opts)
	}
	return &deals.ListDealStagesResponse{}, &client.APIResponse{}, nil
}
func (m *Service) UpdateDeal(ctx context.Context, id string, req interface{}) (interface{}, *client.APIResponse, error) {
	if m.UpdateDealFunc != nil {
		return m.UpdateDealFunc(ctx, id, req)
	}
	return nil, nil, nil
}
func (m *Service) UpdateDealNote(ctx context.Context, dealID, noteID string, req interface{}) (interface{}, *client.APIResponse, error) {
	if m.UpdateDealNoteFunc != nil {
		return m.UpdateDealNoteFunc(ctx, dealID, noteID, req)
	}
	return nil, nil, nil
}
func (m *Service) BulkUpdateDealOwners(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) {
	if m.BulkUpdateDealOwnersFunc != nil {
		return m.BulkUpdateDealOwnersFunc(ctx, req)
	}
	return nil, nil, nil
}
func (m *Service) DeleteDeal(ctx context.Context, id string) (*client.APIResponse, error) {
	if m.DeleteDealFunc != nil {
		return m.DeleteDealFunc(ctx, id)
	}
	return &client.APIResponse{}, nil
}
