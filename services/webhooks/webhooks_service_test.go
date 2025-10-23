package webhooks

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestWebhooks_ServiceMethods_HappyPath(t *testing.T) {
	// Create
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 201}, Body: []byte(`{"webhook":{"id":"1","name":"My Hook","url":"http://example.com","events":["subscribe"],"sources":["public"]}}`)}
	svc := NewRealServiceFromDoer(md)

	crReq := &CreateWebhookRequest{Webhook: &Webhook{Name: "My Hook", URL: "http://example.com", Events: []string{"subscribe"}, Sources: []string{"public"}}}
	crOut, resp, err := svc.CreateWebhook(context.Background(), crReq)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "1", crOut.Webhook.ID)

	// Get
	md = &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"webhook":{"id":"1","name":"My Hook"}}`)}
	svc = NewRealServiceFromDoer(md)
	gOut, _, err := svc.GetWebhook(context.Background(), "1")
	assert.NoError(t, err)
	assert.Equal(t, "1", gOut.Webhook.ID)

	// List
	md = &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"webhooks":[{"id":"1","name":"My Hook"}],"meta":{"total":"1"}}`)}
	svc = NewRealServiceFromDoer(md)
	lOut, _, err := svc.ListWebhooks(context.Background(), map[string]string{"filters[name]": "My"})
	assert.NoError(t, err)
	assert.Len(t, lOut.Webhooks, 1)

	// List events
	md = &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"events":["subscribe","unsubscribe"]}`)}
	svc = NewRealServiceFromDoer(md)
	eOut, _, err := svc.ListWebhookEvents(context.Background(), nil)
	assert.NoError(t, err)
	assert.Contains(t, eOut.Events, "subscribe")

	// Update
	md = &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"webhook":{"id":"1","name":"Updated"}}`)}
	svc = NewRealServiceFromDoer(md)
	upReq := &UpdateWebhookRequest{Webhook: &Webhook{Name: "Updated"}}
	uOut, _, err := svc.UpdateWebhook(context.Background(), "1", upReq)
	assert.NoError(t, err)
	assert.Equal(t, "Updated", uOut.Webhook.Name)

	// Delete
	md = &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{}`)}
	svc = NewRealServiceFromDoer(md)
	respDel, err := svc.DeleteWebhook(context.Background(), "1")
	assert.NoError(t, err)
	assert.NotNil(t, respDel)

	// NewRealService wiring via CoreClient: construct minimal CoreClient to ensure constructor compiles
	cc := &client.CoreClient{}
	_ = NewRealService(cc)
}
