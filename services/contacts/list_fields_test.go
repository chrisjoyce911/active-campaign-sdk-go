package contacts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestRealService_ListFields(t *testing.T) {
	body := []byte(`{"fields":[{"id":"f1","title":"X"}]}`)
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: body}
	svc := NewRealServiceFromDoer(md)

	out, apiResp, err := svc.ListCustomFields(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 200, apiResp.StatusCode)
	if assert.NotNil(t, out) {
		if assert.NotNil(t, out.Fields) {
			assert.Equal(t, "f1", (*out.Fields)[0].ID)
		}
	}
}
