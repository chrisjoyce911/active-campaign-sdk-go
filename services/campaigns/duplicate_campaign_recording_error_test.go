package campaigns

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	th "github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestDuplicateCampaign_Recording_Error400(t *testing.T) {
	body := []byte(`{"message":"Bad request. You are not allowed to copy this campaign."}`)
	rd := &th.RecordingDoer{Resp: &client.APIResponse{StatusCode: 400, Body: body}, Body: body, Err: &client.APIError{StatusCode: 400, Message: "Bad request. You are not allowed to copy this campaign.", Body: body}}
	svc := &service{client: rd}

	_, apiResp, err := svc.DuplicateCampaign(context.Background(), "2")

	// client.CoreClient.Do returns *client.APIError for non-2xx; our mock returns apiResp and no error
	if assert.Error(t, err) {
		if apiErr, ok := err.(*client.APIError); ok {
			assert.Contains(t, apiErr.Message, "not allowed to copy")
			assert.Equal(t, 400, apiErr.StatusCode)
		}
	}
	// ensure recorded request path and method
	assert.Equal(t, "POST", rd.LastMethod)
	assert.Equal(t, "campaigns/2/copy", rd.LastPath)
	_ = apiResp
}
