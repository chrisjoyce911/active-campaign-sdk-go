package contacts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRealService_ListTags(t *testing.T) {
	require := require.New(t)

	body := []byte(`{"tags":[{"id":"52","tag":"Booking Confirmed","tagType":"contact","description":"set when a booking is confirmed"},{"id":"97","tag":"EFA expiring & No booking"}]}`)
	md := &pathRecordingDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: body}

	out, apiResp, err := NewRealServiceFromDoer(md).ListTags(context.Background())
	require.NoError(err)
	require.NotNil(apiResp)
	assert.Equal(t, "tags", md.Path)

	tags := out.TagsOrEmpty()
	require.Len(tags, 2)
	assert.Equal(t, "52", tags[0].ID)
	assert.Equal(t, "Booking Confirmed", tags[0].Tag)
	assert.Equal(t, "EFA expiring & No booking", tags[1].Tag)
}

func TestListTagsResponse_TagsOrEmpty_Nil(t *testing.T) {
	var l *ListTagsResponse
	assert.Empty(t, l.TagsOrEmpty())
	assert.Empty(t, (&ListTagsResponse{}).TagsOrEmpty())
}
