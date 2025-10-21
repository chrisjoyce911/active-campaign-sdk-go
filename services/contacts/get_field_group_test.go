package contacts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestRealService_GetFieldGroup(t *testing.T) {
	body := []byte(`{"fieldGroup":{"id":"g1","title":"G1"}}`)
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: body}
	svc := NewRealServiceFromDoer(md)

	out, apiResp, err := svc.GetFieldGroup(context.Background(), "g1")
	assert.NoError(t, err)
	assert.Equal(t, 200, apiResp.StatusCode)
	if assert.NotNil(t, out) {
		assert.Equal(t, "g1", out.FieldGroup.ID)
	}
}
