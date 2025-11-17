package main

import (
	"bytes"
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/deals"
)

type fakeDealsSvc struct{}

func (f *fakeDealsSvc) ListDeals(ctx context.Context, opts map[string]string) (*deals.ListDealsResponse, *client.APIResponse, error) {
	// Simulate two pages when limit=1, using offset detection
	limit := opts["limit"]
	offset := opts["offset"]
	if limit == "1" && offset == "0" {
		return &deals.ListDealsResponse{Deals: []deals.Deal{{ID: "1", Title: "A", Group: "2", Stage: "7"}}, Meta: &deals.DealsListMeta{Total: 2}}, &client.APIResponse{StatusCode: 200}, nil
	}
	return &deals.ListDealsResponse{Deals: []deals.Deal{{ID: "2", Title: "B", Group: "2", Stage: "7"}}, Meta: &deals.DealsListMeta{Total: 2}}, &client.APIResponse{StatusCode: 200}, nil
}

// Stub unused interface methods
func (f *fakeDealsSvc) CreateDeal(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) { return nil, nil, nil }
func (f *fakeDealsSvc) CreateDealNote(ctx context.Context, dealID string, req interface{}) (interface{}, *client.APIResponse, error) {
	return nil, nil, nil
}
func (f *fakeDealsSvc) GetDeal(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	return nil, nil, nil
}
func (f *fakeDealsSvc) GetDealActivities(ctx context.Context, dealID string) (interface{}, *client.APIResponse, error) {
	return nil, nil, nil
}
func (f *fakeDealsSvc) ListDealStages(ctx context.Context, opts map[string]string) (*deals.ListDealStagesResponse, *client.APIResponse, error) {
	return nil, nil, nil
}
func (f *fakeDealsSvc) UpdateDeal(ctx context.Context, id string, req interface{}) (interface{}, *client.APIResponse, error) {
	return nil, nil, nil
}
func (f *fakeDealsSvc) UpdateDealNote(ctx context.Context, dealID, noteID string, req interface{}) (interface{}, *client.APIResponse, error) {
	return nil, nil, nil
}
func (f *fakeDealsSvc) BulkUpdateDealOwners(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) {
	return nil, nil, nil
}
func (f *fakeDealsSvc) DeleteDeal(ctx context.Context, id string) (*client.APIResponse, error) { return nil, nil }

func TestRun_PrintsDealsAcrossPages(t *testing.T) {
	var buf bytes.Buffer
	svc := &fakeDealsSvc{}
	if err := Run(context.Background(), svc, &buf); err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	out := buf.String()
	if out == "" || !bytes.Contains(buf.Bytes(), []byte("deal 1")) || !bytes.Contains(buf.Bytes(), []byte("deal 2")) {
		t.Fatalf("unexpected output: %s", out)
	}
}
