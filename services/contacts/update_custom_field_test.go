package contacts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestRealService_UpdateCustomField(t *testing.T) {
	body := []byte(`{"field": {"id":"f1","title":"X-Updated"}}`)
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: body}
	svc := NewRealServiceFromDoer(md)

	req := &FieldPayload{Title: "X-Updated"}
	out, apiResp, err := svc.UpdateCustomField(context.Background(), "f1", req)
	assert.NoError(t, err)
	assert.Equal(t, 200, apiResp.StatusCode)
	assert.Equal(t, "f1", out.Field.ID)
}
