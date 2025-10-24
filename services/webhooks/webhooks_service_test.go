package webhooks

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWebhooks_ServiceMethods_HappyPath_TableDriven(t *testing.T) {
	// shared mock payloads used for the various happy-path calls
	createBody := []byte(`{"webhook":{"id":"1","name":"My Hook","url":"http://example.com","events":["subscribe"],"sources":["public"]}}`)
	basicWebhookBody := []byte(`{"webhook":{"id":"1","name":"My Hook"}}`)
	listBody := []byte(`{"webhooks":[{"id":"1","name":"My Hook"}],"meta":{"total":"1"}}`)
	eventsBody := []byte(`{"events":["subscribe","unsubscribe"]}`)
	updateBody := []byte(`{"webhook":{"id":"1","name":"Updated"}}`)

	cases := []struct {
		name string
		call string
		resp *client.APIResponse
		body []byte
	}{
		{name: "create", call: "create", resp: &client.APIResponse{StatusCode: 201}, body: createBody},
		{name: "get", call: "get", resp: &client.APIResponse{StatusCode: 200}, body: basicWebhookBody},
		{name: "list", call: "list", resp: &client.APIResponse{StatusCode: 200}, body: listBody},
		{name: "list-events", call: "list-events", resp: &client.APIResponse{StatusCode: 200}, body: eventsBody},
		{name: "update", call: "update", resp: &client.APIResponse{StatusCode: 200}, body: updateBody},
		{name: "delete", call: "delete", resp: &client.APIResponse{StatusCode: 200}, body: []byte(`{}`)},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			assert := assert.New(t)

			md := &testhelpers.MockDoer{Resp: tt.resp, Body: tt.body}
			require.NotNil(md)

			svc := NewRealServiceFromDoer(md)
			require.NotNil(svc)

			switch tt.call {
			case "create":
				crReq := &CreateWebhookRequest{Webhook: &Webhook{Name: "My Hook", URL: "http://example.com", Events: []string{"subscribe"}, Sources: []string{"public"}}}
				crOut, resp, err := svc.CreateWebhook(context.Background(), crReq)
				assert.NoError(err)
				require.NotNil(resp)
				assert.Equal("1", crOut.Webhook.ID)

			case "get":
				gOut, resp, err := svc.GetWebhook(context.Background(), "1")
				assert.NoError(err)
				require.NotNil(resp)
				assert.Equal("1", gOut.Webhook.ID)

			case "list":
				lOut, resp, err := svc.ListWebhooks(context.Background(), map[string]string{"filters[name]": "My"})
				assert.NoError(err)
				require.NotNil(resp)
				assert.Len(lOut.Webhooks, 1)

			case "list-events":
				eOut, resp, err := svc.ListWebhookEvents(context.Background(), nil)
				assert.NoError(err)
				require.NotNil(resp)
				assert.Contains(eOut.Events, "subscribe")

			case "update":
				upReq := &UpdateWebhookRequest{Webhook: &Webhook{Name: "Updated"}}
				uOut, resp, err := svc.UpdateWebhook(context.Background(), "1", upReq)
				assert.NoError(err)
				require.NotNil(resp)
				assert.Equal("Updated", uOut.Webhook.Name)

			case "delete":
				respDel, err := svc.DeleteWebhook(context.Background(), "1")
				assert.NoError(err)
				require.NotNil(respDel)
			}
		})
	}

	// constructor wiring check (kept separate)
	t.Run("constructor wiring", func(t *testing.T) {
		cc := &client.CoreClient{}
		_ = NewRealService(cc)
	})
}
