package custom_objects

import (
	"context"
	"errors"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestCreateObjectType(t *testing.T) {
	tests := []struct {
		name       string
		doer       *testhelpers.MockDoer
		payload    *CreateObjectTypeRequest
		wantStatus int
		wantErr    bool
	}{
		{"ok", &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 201}, Body: []byte(`{"schema": {"id":"ot1"}}`)}, &CreateObjectTypeRequest{Name: "X"}, 201, false},
		{"ok-empty", &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 201}, Body: []byte(`{"schema": {}}`)}, &CreateObjectTypeRequest{Name: "Y"}, 201, false},
		{"err-400", &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 400}, Body: []byte(`{"error":"bad"}`), Err: errors.New("bad")}, &CreateObjectTypeRequest{Name: ""}, 400, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc := NewRealServiceFromDoer(tc.doer)
			out, apiResp, err := svc.CreateObjectType(context.Background(), tc.payload)
			if tc.wantErr {
				assert.Error(t, err)
				if apiResp != nil {
					assert.Equal(t, tc.wantStatus, apiResp.StatusCode)
				}
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.wantStatus, apiResp.StatusCode)
			if assert.NotNil(t, out) {
				if tc.name == "ok" {
					assert.Equal(t, "ot1", out.ID)
				} else if tc.name == "ok-empty" {
					assert.Equal(t, "", out.ID)
				}
			}
		})
	}
}
