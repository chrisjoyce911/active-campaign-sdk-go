package dealsmock

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/deals"
)

func TestService_ListDealsFuncCalled(t *testing.T) {
	called := false
	m := &Service{ListDealsFunc: func(ctx context.Context, opts map[string]string) (*deals.ListDealsResponse, *client.APIResponse, error) {
		called = true
		return &deals.ListDealsResponse{Deals: []deals.Deal{{ID: "1"}}, Meta: &deals.DealsListMeta{Total: 1}}, &client.APIResponse{StatusCode: 200}, nil
	}}

	resp, api, err := m.ListDeals(context.Background(), map[string]string{"offset": "0"})
	if err != nil || api == nil || resp == nil || len(resp.Deals) != 1 || !called {
		t.Fatalf("unexpected result: resp=%v api=%v err=%v called=%v", resp, api, err, called)
	}
}

func TestService_Defaults(t *testing.T) {
	m := &Service{}
	// Defaults should not panic and should return non-nil zero values where applicable
	resp, api, err := m.ListDeals(context.Background(), nil)
	if err != nil || resp == nil || api == nil {
		t.Fatalf("expected non-nil defaults, got resp=%v api=%v err=%v", resp, api, err)
	}
	delAPI, delErr := m.DeleteDeal(context.Background(), "x")
	if delErr != nil || delAPI == nil {
		t.Fatalf("expected default delete non-nil api, got %v %v", delAPI, delErr)
	}
}

func TestService_AllFuncs(t *testing.T) {
	called := 0
	m := &Service{
		CreateDealFunc: func(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) {
			called++
			return struct{}{}, &client.APIResponse{StatusCode: 200}, nil
		},
		CreateDealNoteFunc: func(ctx context.Context, dealID string, req interface{}) (interface{}, *client.APIResponse, error) {
			called++
			return struct{}{}, &client.APIResponse{StatusCode: 200}, nil
		},
		GetDealFunc: func(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
			called++
			return struct{}{}, &client.APIResponse{StatusCode: 200}, nil
		},
		GetDealActivitiesFunc: func(ctx context.Context, dealID string) (interface{}, *client.APIResponse, error) {
			called++
			return []string{"a"}, &client.APIResponse{StatusCode: 200}, nil
		},
		ListDealsFunc: func(ctx context.Context, opts map[string]string) (*deals.ListDealsResponse, *client.APIResponse, error) {
			called++
			return &deals.ListDealsResponse{}, &client.APIResponse{StatusCode: 200}, nil
		},
		ListDealStagesFunc: func(ctx context.Context, opts map[string]string) (*deals.ListDealStagesResponse, *client.APIResponse, error) {
			called++
			return &deals.ListDealStagesResponse{}, &client.APIResponse{StatusCode: 200}, nil
		},
		UpdateDealFunc: func(ctx context.Context, id string, req interface{}) (interface{}, *client.APIResponse, error) {
			called++
			return struct{}{}, &client.APIResponse{StatusCode: 200}, nil
		},
		UpdateDealNoteFunc: func(ctx context.Context, dealID, noteID string, req interface{}) (interface{}, *client.APIResponse, error) {
			called++
			return struct{}{}, &client.APIResponse{StatusCode: 200}, nil
		},
		BulkUpdateDealOwnersFunc: func(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) {
			called++
			return struct{}{}, &client.APIResponse{StatusCode: 200}, nil
		},
		DeleteDealFunc: func(ctx context.Context, id string) (*client.APIResponse, error) {
			called++
			return &client.APIResponse{StatusCode: 204}, nil
		},
	}

	// Call each method once
	_, _, _ = m.CreateDeal(context.Background(), nil)
	_, _, _ = m.CreateDealNote(context.Background(), "d1", nil)
	_, _, _ = m.GetDeal(context.Background(), "d1")
	_, _, _ = m.GetDealActivities(context.Background(), "d1")
	_, _, _ = m.ListDeals(context.Background(), nil)
	_, _, _ = m.ListDealStages(context.Background(), nil)
	_, _, _ = m.UpdateDeal(context.Background(), "d1", nil)
	_, _, _ = m.UpdateDealNote(context.Background(), "d1", "n1", nil)
	_, _, _ = m.BulkUpdateDealOwners(context.Background(), nil)
	_, _ = m.DeleteDeal(context.Background(), "d1")

	if called != 10 {
		t.Fatalf("expected 10 calls, got %d", called)
	}
}

func TestService_AllDefaults(t *testing.T) {
	m := &Service{}
	ctx := context.Background()
	// invoke every method with nil funcs to cover default returns
	_, _, _ = m.CreateDeal(ctx, nil)
	_, _, _ = m.CreateDealNote(ctx, "d1", nil)
	_, _, _ = m.GetDeal(ctx, "d1")
	_, _, _ = m.GetDealActivities(ctx, "d1")
	_, _, _ = m.ListDeals(ctx, nil)
	_, _, _ = m.ListDealStages(ctx, nil)
	_, _, _ = m.UpdateDeal(ctx, "d1", nil)
	_, _, _ = m.UpdateDealNote(ctx, "d1", "n1", nil)
	_, _, _ = m.BulkUpdateDealOwners(ctx, nil)
	_, _ = m.DeleteDeal(ctx, "d1")
}
