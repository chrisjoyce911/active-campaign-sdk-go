package custom_objects

import (
	"context"
	"errors"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestUpdateObjectRecord(t *testing.T) {
	tests := []struct {
		name    string
		doer    *testhelpers.MockDoer
		payload *UpdateRecordRequest
		wantErr bool
	}{
		{"ok", &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"record":{"id":"r1"}}`)}, &UpdateRecordRequest{Fields: map[string]interface{}{"name": "y"}}, false},
		{"ok-empty", &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"record":{}}`)}, &UpdateRecordRequest{Fields: map[string]interface{}{"name": "y"}}, false},
		{"err-422", &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 422}, Body: []byte(`{"error":"invalid"}`), Err: errors.New("invalid")}, &UpdateRecordRequest{Fields: map[string]interface{}{}}, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc := NewRealServiceFromDoer(tc.doer)
			out, apiResp, err := svc.UpdateObjectRecord(context.Background(), "ot1", "r1", tc.payload)
			if tc.wantErr {
				assert.Error(t, err)
				if apiResp != nil {
					assert.Equal(t, 422, apiResp.StatusCode)
				}
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, apiResp)
			switch tc.name {
			case "ok":
				assert.Equal(t, "r1", out.Record.ID)
			case "ok-empty":
				assert.Equal(t, "", out.Record.ID)
			}
		})
	}
}
