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
