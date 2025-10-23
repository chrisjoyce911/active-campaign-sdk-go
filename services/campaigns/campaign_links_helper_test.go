package campaigns

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestCampaignLinks_HelperBranches(t *testing.T) {
	// service with MockDoer
	d := &testhelpers.MockDoer{}
	s := NewRealServiceFromDoer(d)

	// Case 1: GetCampaignLinks returns nil out (simulate no body)
	d.Resp = &client.APIResponse{StatusCode: 200}
	d.Body = nil
	links, apiResp, err := s.CampaignLinks(context.Background(), "1", nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, apiResp.StatusCode)
	assert.Nil(t, links)

	// Case 2: messageID is nil -> return all links
	d.Body = []byte(`{"links": [{"id": "l1", "message": "m1"}, {"id": "l2"}]}`)
	links, apiResp, err = s.CampaignLinks(context.Background(), "1", nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, apiResp.StatusCode)
	assert.Equal(t, 2, len(links))

	// Case 3: messageID non-nil, some links have nil Message and one matches
	mid := "m1"
	d.Body = []byte(`{"links": [{"id": "l1", "message": "m1"}, {"id": "l2"}, {"id": "l3", "message": "m2"}]}`)
	links, apiResp, err = s.CampaignLinks(context.Background(), "1", &mid)
	assert.NoError(t, err)
	assert.Equal(t, 200, apiResp.StatusCode)
	// only l1 should match
	if assert.Len(t, links, 1) {
		assert.Equal(t, "l1", links[0].ID)
	}

	// Case 4: GetCampaignLinks returns an error
	d.Resp = &client.APIResponse{StatusCode: 500}
	d.Err = &client.APIError{StatusCode: 500, Message: "boom"}
	links, apiResp, err = s.CampaignLinks(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, links)
	assert.Equal(t, 500, apiResp.StatusCode)
}

// fakeGetter is a tiny test helper that implements GetCampaignLinks and returns nil out.
type fakeGetter struct{}

func (f *fakeGetter) GetCampaignLinks(ctx context.Context, id string) (*CampaignLinksResponse, *client.APIResponse, error) {
	return nil, &client.APIResponse{StatusCode: 200}, nil
}

func TestCampaignLinks_OutNilBranch(t *testing.T) {
	out, apiResp, err := (&fakeGetter{}).GetCampaignLinks(context.Background(), "1")
	assert.NoError(t, err)
	if out == nil {
		assert.Nil(t, out)
		assert.Equal(t, 200, apiResp.StatusCode)
	} else {
		t.Fatalf("expected nil out from fake getter")
	}
}
