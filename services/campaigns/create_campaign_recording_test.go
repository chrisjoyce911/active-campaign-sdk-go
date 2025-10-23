package campaigns

import (
	"context"
	"errors"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	th "github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

// Test that CreateCampaign marshals the request payload and calls Do with POST /campaigns
func TestService_CreateCampaign_records_payload(t *testing.T) {
	rd := &th.RecordingDoer{Resp: &client.APIResponse{StatusCode: 201}, Body: []byte(`{"campaign":{"id":"c1"}}`)}
	svc := NewRealServiceFromDoer(rd)

	payload := &CreateCampaignRequest{Name: "My Campaign", Type: "regular"}
	out, apiResp, err := svc.CreateCampaign(context.Background(), payload)

	assert.NoError(t, err)
	if assert.NotNil(t, apiResp) {
		assert.Equal(t, 201, apiResp.StatusCode)
	}
	_ = out

	// recording assertions
	assert.Equal(t, "POST", rd.LastMethod)
	assert.Equal(t, "campaigns", rd.LastPath)
	// LastV should be the original payload (or a close equivalent). We verify the recorded JSON body contains the title.
	if assert.NotNil(t, rd.LastBody) {
		assert.Contains(t, string(rd.LastBody), "My Campaign")
		assert.Contains(t, string(rd.LastBody), "regular")
	}
}

// Test that when the underlying Doer returns an error (e.g. unmarshal/parsing), CreateCampaign surfaces it.
func TestService_CreateCampaign_propagates_doer_error(t *testing.T) {
	// create a Doer that returns an error
	bad := &th.MockDoer{Err: errors.New("unmarshal failed")}
	svc := NewRealServiceFromDoer(bad)

	_, _, err := svc.CreateCampaign(context.Background(), &CreateCampaignRequest{Name: "t", Type: "single"})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unmarshal failed")
}
