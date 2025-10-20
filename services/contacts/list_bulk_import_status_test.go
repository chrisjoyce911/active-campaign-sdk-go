package contacts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/stretchr/testify/assert"
)

func TestRealService_ListBulkImportStatus(t *testing.T) {
	md := &mockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"jobs":[{"id":"j1"}]}`)}
	svc := NewRealServiceFromDoer(md)

	out, apiResp, err := svc.ListBulkImportStatus(context.Background(), map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, 200, apiResp.StatusCode)
	_ = out
}
