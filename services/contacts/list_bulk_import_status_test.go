package contacts

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestListBulkImportStatus_NoOpts(t *testing.T) {
	outBody := []byte(`[{"id":"b1"}]`)
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: http.StatusOK, Body: outBody}, Body: outBody}
	svc := NewRealServiceFromDoer(md)

	_, apiResp, err := svc.ListBulkImportStatus(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if apiResp == nil || apiResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %+v", apiResp)
	}
	var parsed []map[string]interface{}
	if err := json.Unmarshal(apiResp.Body, &parsed); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}
}

func TestListBulkImportStatus_WithOpts(t *testing.T) {
	outBody := []byte(`[{"id":"b2"}]`)
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: http.StatusOK, Body: outBody}, Body: outBody}
	svc := NewRealServiceFromDoer(md)

	opts := map[string]string{"limit": "5", "status": "done"}
	_, apiResp, err := svc.ListBulkImportStatus(context.Background(), opts)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if apiResp == nil || apiResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %+v", apiResp)
	}
	var parsed []map[string]interface{}
	if err := json.Unmarshal(apiResp.Body, &parsed); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}
}

func TestRealService_ListBulkImportStatus(t *testing.T) {
	md := &mockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"jobs":[{"id":"j1"}]}`)}
	svc := NewRealServiceFromDoer(md)

	out, apiResp, err := svc.ListBulkImportStatus(context.Background(), map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, 200, apiResp.StatusCode)
	_ = out
}
