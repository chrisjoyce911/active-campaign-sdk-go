package contacts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestRealService_ListFieldValues(t *testing.T) {
	body := []byte(`{"fieldValues":[{"id":"fv1","value":"v"}]}`)
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: body}
	svc := NewRealServiceFromDoer(md)

	out, apiResp, err := svc.ListFieldValues(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 200, apiResp.StatusCode)
	if assert.NotNil(t, out) {
		if assert.NotNil(t, out.FieldValues) {
			assert.Equal(t, "fv1", (*out.FieldValues)[0].ID)
		}
	}
}
