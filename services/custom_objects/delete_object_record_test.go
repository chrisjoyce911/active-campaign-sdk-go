package custom_objects

import (
	"context"
	"errors"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestDeleteObjectRecord(t *testing.T) {
	tests := []struct {
		name       string
		doer       *testhelpers.MockDoer
		wantStatus int
		wantErr    bool
	}{
		{"ok", &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"success":true}`)}, 200, false},
		{"err-notfound", &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 404}, Body: []byte(`{"error":"not found"}`), Err: errors.New("not found")}, 404, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc := NewRealServiceFromDoer(tc.doer)
			apiResp, err := svc.DeleteObjectRecord(context.Background(), "ot1", "r1")
			if tc.wantErr {
				assert.Error(t, err)
				if apiResp != nil {
					assert.Equal(t, tc.wantStatus, apiResp.StatusCode)
				}
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.wantStatus, apiResp.StatusCode)
		})
	}
}
