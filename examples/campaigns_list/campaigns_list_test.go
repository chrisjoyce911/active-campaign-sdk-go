package main

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	campaignsmock "github.com/chrisjoyce911/active-campaign-sdk-go/mocks/campaigns"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/campaigns"
)

func TestRun_PrintsStatuses(t *testing.T) {
	fake := &campaignsmock.Service{ListCampaignsFunc: func(ctx context.Context, opts interface{}) (*campaigns.ListCampaignsResponse, *client.APIResponse, error) {
		return &campaigns.ListCampaignsResponse{Campaigns: []campaigns.Campaign{{ID: "1", Name: "One", Status: "1"}, {ID: "2", Name: "Two", Status: "0"}}}, &client.APIResponse{StatusCode: 200}, nil
	}}

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
	fake := &campaignsmock.Service{ListCampaignsFunc: func(ctx context.Context, opts interface{}) (*campaigns.ListCampaignsResponse, *client.APIResponse, error) {
		return nil, &client.APIResponse{StatusCode: 500}, fmt.Errorf("boom")
	}}

	var buf bytes.Buffer
	if err := Run(context.Background(), fake, &buf); err == nil {
		t.Fatalf("expected error from Run")
	}
}

func TestRun_StatusParseError(t *testing.T) {
	fake := &campaignsmock.Service{ListCampaignsFunc: func(ctx context.Context, opts interface{}) (*campaigns.ListCampaignsResponse, *client.APIResponse, error) {
		return &campaigns.ListCampaignsResponse{Campaigns: []campaigns.Campaign{{ID: "x", Name: "Bad", Status: "not-a-number"}}}, &client.APIResponse{StatusCode: 200}, nil
	}}
	var buf bytes.Buffer
	if err := Run(context.Background(), fake, &buf); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(buf.String(), "status parse error") {
		t.Fatalf("expected status parse error in output, got: %q", buf.String())
	}
}
