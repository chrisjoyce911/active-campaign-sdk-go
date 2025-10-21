package contacts

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestRealService_UpdateContact_RequestShape(t *testing.T) {
	tests := []struct {
		name             string
		req              *CreateContactRequest
		respBody         []byte
		respStatus       int
		wantPath         string
		wantMethod       string
		wantBodyContains string
	}{
		{name: "basic", req: &CreateContactRequest{Contact: &Contact{Email: "a@b.com", FirstName: "A"}}, respBody: []byte(`{"contact": {"id":"1","email":"a@b.com"}}`), respStatus: 200, wantPath: "contacts/1", wantMethod: "PUT", wantBodyContains: `"email":"a@b.com"`},
		{name: "with fieldValues", req: &CreateContactRequest{Contact: &Contact{Email: "b@c.com", FieldValues: &[]FieldValue{{Field: "1", Value: "V"}}}}, respBody: []byte(`{"contact": {"id":"2","email":"b@c.com","fieldValues":[{"field":"1","value":"V"}]}}`), respStatus: 200, wantPath: "contacts/2", wantMethod: "PUT", wantBodyContains: `"fieldValues"`},
		{name: "validation error 422", req: &CreateContactRequest{Contact: &Contact{Email: "bad"}}, respBody: []byte(`{"errors":[{"title":"invalid","detail":"email invalid"}]}`), respStatus: 422, wantPath: "contacts/999", wantMethod: "PUT", wantBodyContains: `"email":"bad"`},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			hd := &testhelpers.HTTPDoer{BaseURL: "https://example.api-us1.com/api/3/", Token: "tok", RespStatus: tc.respStatus, RespBody: tc.respBody}
			svc := NewRealServiceFromDoer(hd)

			// choose id based on resp body id if present, else use 999
			id := "999"
			var parsed struct {
				Contact struct {
					ID string `json:"id"`
				} `json:"contact"`
			}
			_ = json.Unmarshal(tc.respBody, &parsed)
			if parsed.Contact.ID != "" {
				id = parsed.Contact.ID
			}

			out, apiResp, err := svc.UpdateContact(context.Background(), id, tc.req)

			// Even for 4xx our Mock doesn't return error; the client will set APIResponse
			if tc.respStatus >= 200 && tc.respStatus < 300 {
				assert.NoError(t, err)
				assert.NotNil(t, out)
			}
			if apiResp != nil {
				assert.Equal(t, tc.respStatus, apiResp.StatusCode)
			}

			// verify the underlying HTTP request
			assert.Equal(t, "PUT", hd.LastRequest.Method)
			assert.Contains(t, hd.LastRequest.URL.String(), "/api/3/contacts/")
			assert.Equal(t, "application/json", hd.LastRequest.Header.Get("Content-Type"))
			assert.Equal(t, "tok", hd.LastRequest.Header.Get("Api-Token"))
			if hd.LastRequestBody == nil {
				t.Fatalf("no recorded request body")
			}
			expB, _ := json.Marshal(tc.req)
			assert.JSONEq(t, string(expB), string(hd.LastRequestBody))
		})
	}
}
