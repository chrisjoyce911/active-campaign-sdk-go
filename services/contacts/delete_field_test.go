package contacts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestRealService_DeleteField(t *testing.T) {
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"success":true}`)}
	svc := NewRealServiceFromDoer(md)

	apiResp, err := svc.DeleteCustomField(context.Background(), "f1")
	assert.NoError(t, err)
	assert.Equal(t, 200, apiResp.StatusCode)
}
