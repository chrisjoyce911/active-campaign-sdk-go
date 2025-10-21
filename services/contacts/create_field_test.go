package contacts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestRealService_CreateField(t *testing.T) {
	body := []byte(`{"field": {"id":"f1","title":"X"}}`)
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 201}, Body: body}
	svc := NewRealServiceFromDoer(md)

	req := &FieldPayload{Title: "X", Type: "text"}
	out, apiResp, err := svc.CreateCustomField(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, 201, apiResp.StatusCode)
	assert.Equal(t, "f1", out.Field.ID)
}
