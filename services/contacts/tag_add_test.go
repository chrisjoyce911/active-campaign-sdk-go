package contacts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRealService_TagAdd(t *testing.T) {
	tests := []struct {
		name      string
		contactID string
		tagID     string
		body      []byte
	}{
		{
			name:      "add tag to contact",
			contactID: "1",
			tagID:     "20",
			body:      []byte(`{"contactTag":{"cdate":"2017-06-08T16:11:53-05:00","contact":"1","id":"1","links":{"contact":"/1/contact","tag":"/1/tag"},"tag":"20"}}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &mockDoer{Resp: &client.APIResponse{StatusCode: 201}, Body: tt.body}
			require := require.New(t)
			require.NotNil(md)

			svc := NewRealServiceFromDoer(md)
			require.NotNil(svc)

			out, apiResp, err := svc.TagAdd(context.Background(), tt.contactID, tt.tagID)
			assert.NoError(t, err)
			require.NotNil(apiResp)
			assert.Equal(t, 201, apiResp.StatusCode)
			if assert.NotNil(t, out) {
				assert.Equal(t, tt.contactID, out.ContactTag.Contact)
				assert.Equal(t, tt.tagID, out.ContactTag.Tag)
			}
		})
	}
}
