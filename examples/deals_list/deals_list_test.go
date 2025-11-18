package main

import (
	"bytes"
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	dealsmock "github.com/chrisjoyce911/active-campaign-sdk-go/mocks/deals"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/deals"
)

func TestRun_PrintsDealsAcrossPages(t *testing.T) {
	var buf bytes.Buffer
	svc := &dealsmock.Service{
		ListDealsFunc: func(ctx context.Context, opts map[string]string) (*deals.ListDealsResponse, *client.APIResponse, error) {
			offset := opts["offset"]
			if offset == "0" {
				return &deals.ListDealsResponse{Deals: []deals.Deal{{ID: "1", Title: "A", Group: "2", Stage: "7"}, {ID: "2", Title: "B", Group: "2", Stage: "7"}}, Meta: &deals.DealsListMeta{Total: 2}}, &client.APIResponse{StatusCode: 200}, nil
			}
			return &deals.ListDealsResponse{Deals: []deals.Deal{}, Meta: &deals.DealsListMeta{Total: 2}}, &client.APIResponse{StatusCode: 200}, nil
		},
	}
	if err := Run(context.Background(), svc, &buf); err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	out := buf.String()
	if out == "" || !bytes.Contains(buf.Bytes(), []byte("deal 1")) || !bytes.Contains(buf.Bytes(), []byte("deal 2")) {
		t.Fatalf("unexpected output: %s", out)
	}
}
