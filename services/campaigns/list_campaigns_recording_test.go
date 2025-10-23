package campaigns

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	th "github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestService_ListCampaigns_records_opts(t *testing.T) {
	rd := &th.RecordingDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"campaigns": [{"id":"c1"}]}`)}
	svc := NewRealServiceFromDoer(rd)

	opts := map[string]interface{}{"limit": 5, "status": "active"}
	out, apiResp, err := svc.ListCampaigns(context.Background(), opts)

	assert.NoError(t, err)
	if assert.NotNil(t, apiResp) {
		assert.Equal(t, 200, apiResp.StatusCode)
	}
	_ = out

	// Recording assertions: method/path and that the opts were marshalled
	assert.Equal(t, "GET", rd.LastMethod)
	assert.Equal(t, "campaigns", rd.LastPath)
	if assert.NotNil(t, rd.LastBody) {
		assert.Contains(t, string(rd.LastBody), "limit")
		assert.Contains(t, string(rd.LastBody), "active")
	}
}
