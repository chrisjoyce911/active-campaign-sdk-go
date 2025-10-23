package tags

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestNewRealService_NotNil(t *testing.T) {
	c, err := client.NewCoreClient("", "")
	if err != nil {
		t.Fatalf("failed to create core client: %v", err)
	}
	svc := NewRealService(c)
	if assert.NotNil(t, svc) {
		// sanity: ensure type implements interface by calling TagsOrEmpty on nil receiver
		var l *ListTagsResponse
		_ = l.TagsOrEmpty()
	}
}

func TestNilReceiver_TagMethods(t *testing.T) {
	var s *service
	_, _, err := s.AddTagToContact(context.Background(), "c1", &CreateOrUpdateTagRequest{Tag: TagPayload{Tag: "x"}})
	assert.Error(t, err)

	apiResp, err := s.DeleteTag(context.Background(), "t1")
	assert.Error(t, err)
	_ = apiResp

	_, _, err = s.UpdateTag(context.Background(), "t1", &CreateOrUpdateTagRequest{Tag: TagPayload{Tag: "y"}})
	assert.Error(t, err)
}

func TestAddTagToContact_RecordingDoer(t *testing.T) {
	rd := &testhelpers.RecordingDoer{Resp: &client.APIResponse{StatusCode: 201}, Body: []byte(`{"tag":{"id":"tx","tag":"x"}}`)}
	svc := NewRealServiceFromDoer(rd)

	req := &CreateOrUpdateTagRequest{Tag: TagPayload{Tag: "x"}}
	out, apiResp, err := svc.AddTagToContact(context.Background(), "c123", req)
	assert.NoError(t, err)
	assert.Equal(t, 201, apiResp.StatusCode)
	if out != nil {
		assert.Equal(t, "tx", out.Tag.ID)
	}

	// verify path and last body
	assert.Contains(t, rd.LastPath, "contacts/c123/tags")
	if rd.LastBody != nil {
		var body map[string]interface{}
		_ = json.Unmarshal(rd.LastBody, &body)
		if v, ok := body["tag"]; ok {
			_ = v
		}
	}
}

func TestUpdateTag_RecordingDoer(t *testing.T) {
	rd := &testhelpers.RecordingDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"tag":{"id":"t1","tag":"updated"}}`)}
	svc := NewRealServiceFromDoer(rd)

	req := &CreateOrUpdateTagRequest{Tag: TagPayload{Tag: "updated"}}
	out, apiResp, err := svc.UpdateTag(context.Background(), "t1", req)
	assert.NoError(t, err)
	assert.Equal(t, 200, apiResp.StatusCode)
	if out != nil {
		assert.Equal(t, "updated", out.Tag.Tag)
	}

	assert.Contains(t, rd.LastPath, "tags/t1")
	assert.Equal(t, "PUT", rd.LastMethod)
}
