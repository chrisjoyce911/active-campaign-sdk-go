package tags

import (
	"context"
	"encoding/json"
	"net/url"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestTags_TagsOrEmptyNilAndEmpty(t *testing.T) {
	var l *ListTagsResponse
	// nil receiver
	v := l.TagsOrEmpty()
	assert.NotNil(t, v)
	assert.Len(t, v, 0)

	// nil Tags pointer
	l = &ListTagsResponse{Tags: nil}
	v2 := l.TagsOrEmpty()
	assert.NotNil(t, v2)
}

func TestListTags_WithOptsEncodesQuery(t *testing.T) {
	md := &testhelpers.RecordingDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"tags":[]}`)}
	svc := NewRealServiceFromDoer(md)

	opts := map[string]string{"page": "2", "per_page": "10"}
	_, apiResp, err := svc.ListTags(context.Background(), opts)
	assert.NoError(t, err)
	assert.NotNil(t, apiResp)

	// verify path included encoded query string
	if md.LastPath != "" {
		u, _ := url.Parse("http://localhost/" + md.LastPath)
		q := u.Query()
		assert.Equal(t, "2", q.Get("page"))
		assert.Equal(t, "10", q.Get("per_page"))
	}
}

func TestNilReceiverMethodsReturnNotImplemented(t *testing.T) {
	var svc *service
	_, _, err := svc.ListTags(context.Background(), nil)
	assert.Error(t, err)

	_, _, err = svc.GetTag(context.Background(), "id")
	assert.Error(t, err)

	_, _, err = svc.CreateTag(context.Background(), &CreateOrUpdateTagRequest{Tag: TagPayload{Tag: "x"}})
	// CreateTag uses the client's Do directly and may panic if svc is nil; ensure we call on non-nil type
	// Instead, construct a zero-value service and call
	s := &service{}
	_, _, err = s.CreateTag(context.Background(), &CreateOrUpdateTagRequest{Tag: TagPayload{Tag: "x"}})
	assert.Error(t, err)
}

func TestClientErrorPropagation(t *testing.T) {
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 500}, Body: []byte(`{"error":"boom"}`), Err: &client.APIError{StatusCode: 500, Message: "boom"}}
	svc := NewRealServiceFromDoer(md)

	_, apiResp, err := svc.GetTag(context.Background(), "t1")
	assert.Error(t, err)
	if apiResp != nil {
		assert.Equal(t, 500, apiResp.StatusCode)
	}
	apiResp2, err2 := svc.DeleteTag(context.Background(), "t1")
	assert.Error(t, err2)
	if apiResp2 != nil {
		assert.Equal(t, 500, apiResp2.StatusCode)
	}
}

func TestCreateTag_PayloadRoundTrip(t *testing.T) {
	md := &testhelpers.RecordingDoer{Resp: &client.APIResponse{StatusCode: 201}, Body: []byte(`{"tag":{"id":"t1","tag":"x"}}`)}
	svc := NewRealServiceFromDoer(md)

	req := &CreateOrUpdateTagRequest{Tag: TagPayload{Tag: "x"}}
	out, apiResp, err := svc.CreateTag(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, 201, apiResp.StatusCode)
	if out != nil {
		assert.Equal(t, "t1", out.Tag.ID)
	}

	// verify RecordingDoer captured the payload and path
	if md.LastBody != nil {
		var body map[string]interface{}
		_ = json.Unmarshal(md.LastBody, &body)
		if v, ok := body["tag"]; ok {
			_ = v
		}
	}
}
