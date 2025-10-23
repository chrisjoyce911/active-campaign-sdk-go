package contacts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
)

func TestUpdateField_DelegatesToUpdateCustomField(t *testing.T) {
	// Verify UpdateField delegates to UpdateCustomField and decodes response
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"field":{"id":"f1"}}`)}
	svc := NewRealServiceFromDoer(md)

	req := &FieldPayload{Title: "T"}
	out, apiResp, err := svc.UpdateField(context.Background(), "f1", req)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if apiResp == nil || apiResp.StatusCode != 200 {
		t.Fatalf("expected 200, got %+v", apiResp)
	}
	if out == nil || out.Field.ID != "f1" {
		t.Fatalf("unexpected out: %+v", out)
	}
}
