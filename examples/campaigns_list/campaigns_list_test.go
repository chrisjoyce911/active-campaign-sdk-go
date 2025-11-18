package main

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/campaigns"
)

type fakeCampaignsSvc struct {
	resp *campaigns.ListCampaignsResponse
	api  *client.APIResponse
	err  error
}

func (f *fakeCampaignsSvc) ListCampaigns(ctx context.Context, opts interface{}) (*campaigns.ListCampaignsResponse, *client.APIResponse, error) {
	return f.resp, f.api, f.err
}

// The rest of the campaigns.CampaignsService interface methods are not used in this example.
// Provide no-op implementations to satisfy the interface.
func (f *fakeCampaignsSvc) CreateCampaign(ctx context.Context, req *campaigns.CreateCampaignRequest) (*campaigns.Campaign, *client.APIResponse, error) {
	return nil, nil, nil
}
func (f *fakeCampaignsSvc) GetCampaign(ctx context.Context, id string) (*campaigns.Campaign, *client.APIResponse, error) {
	return nil, nil, nil
}
func (f *fakeCampaignsSvc) GetCampaignLinks(ctx context.Context, id string) (*campaigns.CampaignLinksResponse, *client.APIResponse, error) {
	return nil, nil, nil
}
func (f *fakeCampaignsSvc) CampaignLinks(ctx context.Context, id string, messageID *string) ([]campaigns.CampaignLink, *client.APIResponse, error) {
	return nil, nil, nil
}
func (f *fakeCampaignsSvc) EditCampaign(ctx context.Context, id string, req *campaigns.EditCampaignRequest) (*campaigns.Campaign, *client.APIResponse, error) {
	return nil, nil, nil
}
func (f *fakeCampaignsSvc) DuplicateCampaign(ctx context.Context, id string) (*campaigns.DuplicateCampaignResponse, *client.APIResponse, error) {
	return nil, nil, nil
}

func TestRun_PrintsStatuses(t *testing.T) {
	fake := &fakeCampaignsSvc{
		resp: &campaigns.ListCampaignsResponse{Campaigns: []campaigns.Campaign{
			{ID: "1", Name: "One", Status: "1"},
			{ID: "2", Name: "Two", Status: "0"},
		}},
	}

	var buf bytes.Buffer
	if err := Run(context.Background(), fake, &buf); err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	out := buf.String()
	if out == "" {
		t.Fatalf("expected output, got empty string")
	}
	if !bytes.Contains([]byte(out), []byte("campaign 1 (One):")) {
		t.Fatalf("output missing campaign 1: %q", out)
	}
	if !bytes.Contains([]byte(out), []byte("campaign 2 (Two):")) {
		t.Fatalf("output missing campaign 2: %q", out)
	}
}

func TestRun_ErrorFromService(t *testing.T) {
	fake := &fakeCampaignsSvc{
		resp: nil,
		api:  &client.APIResponse{StatusCode: 500},
		err:  fmt.Errorf("boom"),
	}
	var buf bytes.Buffer
	if err := Run(context.Background(), fake, &buf); err == nil {
		t.Fatalf("expected error from Run")
	}
}

func TestRun_StatusParseError(t *testing.T) {
	fake := &fakeCampaignsSvc{
		resp: &campaigns.ListCampaignsResponse{Campaigns: []campaigns.Campaign{
			{ID: "x", Name: "Bad", Status: "not-a-number"},
		}},
	}
	var buf bytes.Buffer
	if err := Run(context.Background(), fake, &buf); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(buf.String(), "status parse error") {
		t.Fatalf("expected status parse error in output, got: %q", buf.String())
	}
}
