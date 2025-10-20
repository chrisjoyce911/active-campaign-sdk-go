package contacts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestRealService_UpdateFieldGroup(t *testing.T) {
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"success":true}`)}
	svc := NewRealServiceFromDoer(md)

	apiResp, err := svc.UpdateFieldGroup(context.Background(), "g1", map[string]interface{}{"title": "G1-Updated"})
	assert.NoError(t, err)
	assert.Equal(t, 200, apiResp.StatusCode)
}
