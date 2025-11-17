package main

import (
	"bytes"
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/deals"
)

type fakeDealsSvc struct {
	stages *deals.ListDealStagesResponse
	api    *client.APIResponse
	err    error
}

func (f *fakeDealsSvc) ListDealStages(ctx context.Context, opts map[string]string) (*deals.ListDealStagesResponse, *client.APIResponse, error) {
	return f.stages, f.api, f.err
}

// Stub the rest of the interface methods (not used by this example).
func (f *fakeDealsSvc) CreateDeal(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) {
	return nil, nil, nil
}
func (f *fakeDealsSvc) CreateDealNote(ctx context.Context, dealID string, req interface{}) (interface{}, *client.APIResponse, error) {
	return nil, nil, nil
}
func (f *fakeDealsSvc) GetDeal(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	return nil, nil, nil
}
func (f *fakeDealsSvc) GetDealActivities(ctx context.Context, dealID string) (interface{}, *client.APIResponse, error) {
	return nil, nil, nil
}
func (f *fakeDealsSvc) ListDeals(ctx context.Context, opts map[string]string) (*deals.ListDealsResponse, *client.APIResponse, error) {
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
func (f *fakeDealsSvc) DeleteDeal(ctx context.Context, id string) (*client.APIResponse, error) {
	return nil, nil
}

func TestRun_PrintsStages(t *testing.T) {
	fake := &fakeDealsSvc{stages: &deals.ListDealStagesResponse{DealStages: []deals.DealStage{
		{ID: "15", Title: "Initial Contact", Group: "4"},
		{ID: "16", Title: "Qualifications - Low", Group: "4"},
	}}}
	var buf bytes.Buffer
	if err := Run(context.Background(), fake, &buf); err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	out := buf.String()
	if out == "" {
		t.Fatalf("expected output, got empty string")
	}
	if !bytes.Contains(buf.Bytes(), []byte("stage 15")) || !bytes.Contains(buf.Bytes(), []byte("stage 16")) {
		t.Fatalf("unexpected output: %s", out)
	}
}
