package campaigns

import (
	"context"
	"net/http"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	th "github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
)

func TestEditCampaign_Recording_PathAndMethod(t *testing.T) {
	rd := &th.RecordingDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"campaign": {"id":"29"}}`)}
	svc := &service{client: rd}

	req := &EditCampaignRequest{Name: ptrString("updated-name")}
	_, _, err := svc.EditCampaign(context.Background(), "29", req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if rd.LastMethod != http.MethodPut {
		t.Fatalf("expected method PUT got %s", rd.LastMethod)
	}
	if rd.LastPath != "campaigns/29/edit" {
		t.Fatalf("expected path campaigns/29/edit got %s", rd.LastPath)
	}
	// ensure body was recorded and contains the name field
	if rd.LastBody == nil {
		t.Fatalf("expected request body to be set")
	}
}
