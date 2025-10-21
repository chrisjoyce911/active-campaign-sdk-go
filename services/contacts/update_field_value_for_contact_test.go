package contacts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestRealService_UpdateFieldValueForContact(t *testing.T) {
	body := []byte(`{"fieldValue": {"id":"fv1","value":"v"}}`)
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 201}, Body: body}
	svc := NewRealServiceFromDoer(md)

	req := &FieldValuePayload{Contact: "c1", Field: "f1", Value: "v"}
	out, apiResp, err := svc.UpdateFieldValueForContact(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, 201, apiResp.StatusCode)
	assert.Equal(t, "fv1", out.FieldValue.ID)
}
