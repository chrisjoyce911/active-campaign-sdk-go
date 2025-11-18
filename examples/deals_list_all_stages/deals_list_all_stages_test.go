package main

import (
	"bytes"
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	dealsmock "github.com/chrisjoyce911/active-campaign-sdk-go/mocks/deals"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/deals"
)

func TestRun_PrintsStages(t *testing.T) {
	fake := &dealsmock.Service{ListDealStagesFunc: func(ctx context.Context, opts map[string]string) (*deals.ListDealStagesResponse, *client.APIResponse, error) {
		return &deals.ListDealStagesResponse{DealStages: []deals.DealStage{{ID: "15", Title: "Initial Contact", Group: "4"}, {ID: "16", Title: "Qualifications - Low", Group: "4"}}}, &client.APIResponse{StatusCode: 200}, nil
	}}
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
