package contacts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestRealService_AddFieldToGroup(t *testing.T) {
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 201}, Body: []byte(`{"success":true}`)}
	svc := NewRealServiceFromDoer(md)

	apiResp, err := svc.AddFieldToGroup(context.Background(), map[string]interface{}{"field": "f1", "group": "g1"})
	assert.NoError(t, err)
	assert.Equal(t, 201, apiResp.StatusCode)
}
