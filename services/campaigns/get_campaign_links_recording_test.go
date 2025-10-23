package campaigns

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	th "github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestService_GetCampaignLinks_Recording_Path(t *testing.T) {
	rd := &th.RecordingDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"links": []}`)}
	svc := NewRealServiceFromDoer(rd)

	_, _, err := svc.GetCampaignLinks(context.Background(), "42")
	assert.NoError(t, err)

	assert.Equal(t, "GET", rd.LastMethod)
	assert.Equal(t, "campaigns/42/links", rd.LastPath)
}

func TestService_CampaignLinks_Filtering(t *testing.T) {
	body := []byte(`{"links":[{"id":"1","message":"3"},{"id":"2","message":"4"},{"id":"3","message":null}]}`)
	rd := &th.RecordingDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: body}
	svc := NewRealServiceFromDoer(rd)

	mid := "3"
	links, _, err := svc.CampaignLinks(context.Background(), "1", &mid)
	assert.NoError(t, err)
	if assert.NotNil(t, links) {
		assert.Len(t, links, 1)
		assert.Equal(t, "1", links[0].ID)
	}
}
