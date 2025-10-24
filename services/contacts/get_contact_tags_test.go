package contacts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRealService_GetContactTags(t *testing.T) {
	tests := []struct {
		name string
		id   string
		body []byte
	}{
		{name: "tags present", id: "1", body: []byte(`{"contactTags":[{"id":"10","tag":"x","contact":"1","cdate":"2025-01-01T00:00:00-06:00"}]}`)},
		{name: "no tags", id: "2", body: []byte(`{"contactTags":[]}`)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &mockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: tt.body}
			require := require.New(t)
			require.NotNil(md)

			svc := NewRealServiceFromDoer(md)
			require.NotNil(svc)

			out, apiResp, err := svc.GetContactTags(context.Background(), tt.id)
			assert.NoError(t, err)
			require.NotNil(apiResp)
			assert.Equal(t, 200, apiResp.StatusCode)
			if tt.name == "tags present" {
				if assert.NotNil(t, out.ContactTags) {
					assert.Len(t, *out.ContactTags, 1)
					assert.Equal(t, "10", (*out.ContactTags)[0].ID)
					assert.Equal(t, "x", (*out.ContactTags)[0].Tag)
				}
			} else {
				// when API returns an empty array we expect a non-nil pointer to an empty slice
				if assert.NotNil(t, out.ContactTags) {
					assert.Len(t, *out.ContactTags, 0)
				}
			}
		})
	}
}
