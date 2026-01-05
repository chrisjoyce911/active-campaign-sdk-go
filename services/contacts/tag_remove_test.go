package contacts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRealService_TagRemove(t *testing.T) {
	tests := []struct {
		name       string
		contactID  string
		tag        string
		tagsBody   []byte
		statusCode int
	}{
		{
			name:       "remove tag from contact",
			contactID:  "1",
			tag:        "foo",
			tagsBody:   []byte(`{"contactTags":[{"id":"22","tag":"foo","contact":"1","cdate":"2025-01-01T00:00:00-06:00"}]}`),
			statusCode: 200,
		},
		{
			name:       "tag not found",
			contactID:  "1",
			tag:        "bar",
			tagsBody:   []byte(`{"contactTags":[{"id":"22","tag":"foo","contact":"1","cdate":"2025-01-01T00:00:00-06:00"}]}`),
			statusCode: 404,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &mockDoer{Resp: &client.APIResponse{StatusCode: tt.statusCode}, Body: tt.tagsBody}
			require := require.New(t)
			require.NotNil(md)

			svc := NewRealServiceFromDoer(md)
			require.NotNil(svc)

			apiResp, err := svc.TagRemove(context.Background(), tt.contactID, tt.tag)
			assert.NoError(t, err)
			require.NotNil(apiResp)
			assert.Equal(t, tt.statusCode, apiResp.StatusCode)
		})
	}
}
