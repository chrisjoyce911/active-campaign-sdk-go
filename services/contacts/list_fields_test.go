package contacts

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
)

func TestListFields_GetPathAndDecode(t *testing.T) {
	lf := &ListFieldsResponse{Fields: &[]FieldPayload{{ID: "f1", Title: "T"}}}
	b, _ := json.Marshal(lf)
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: http.StatusOK, Body: b}, Body: b}
	svc := NewRealServiceFromDoer(md)

	out, apiResp, err := svc.ListFields(context.Background())
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if apiResp == nil || apiResp.StatusCode != 200 {
		t.Fatalf("expected 200, got %+v", apiResp)
	}
	if out == nil || len(out.FieldsOrEmpty()) != 1 || out.FieldsOrEmpty()[0].ID != "f1" {
		t.Fatalf("unexpected out: %+v", out)
	}
}
