package ecommerce

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	th "github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestCreateGetUpdateDeleteListOrders(t *testing.T) {
	t.Run("CreateOrder success and doer error", func(t *testing.T) {
		// success case: MockDoer returns a wrapper body
		createBody := []byte(`{"ecomOrder":{"id":"1","externalid":"ext-1"}}`)
		md := &th.MockDoer{Resp: &client.APIResponse{StatusCode: http.StatusOK}, Body: createBody}
		svc := NewRealServiceFromDoer(md).(*service)

		out, apiResp, err := svc.CreateOrder(context.Background(), CreateOrderRequest{EcomOrder: EcomOrder{ExternalID: "ext-1"}})
		assert.NoError(t, err)
		assert.NotNil(t, apiResp)
		assert.NotNil(t, out)
		assert.Equal(t, "ext-1", out.ExternalID)

		// doer error case
		mdErr := &th.MockDoer{Err: errors.New("boom")}
		svcErr := NewRealServiceFromDoer(mdErr).(*service)
		// nil isn't a valid typed request; call with zero value
		_, _, err = svcErr.CreateOrder(context.Background(), CreateOrderRequest{})
		assert.Error(t, err)
	})

	t.Run("Get/Update/Delete/List success and error paths", func(t *testing.T) {
		// GetOrder success
		getBody := []byte(`{"ecomOrder":{"id":"1","externalid":"ext-1"}}`)
		mdGet := &th.MockDoer{Resp: &client.APIResponse{StatusCode: http.StatusOK}, Body: getBody}
		svcGet := NewRealServiceFromDoer(mdGet).(*service)
		out, apiResp, err := svcGet.GetOrder(context.Background(), "1")
		assert.NoError(t, err)
		assert.NotNil(t, apiResp)
		assert.NotNil(t, out)
		assert.Equal(t, "1", out.ID)

		// UpdateOrder success
		updateBody := []byte(`{"ecomOrder":{"id":"1","externalid":"updated"}}`)
		mdUpd := &th.MockDoer{Resp: &client.APIResponse{StatusCode: http.StatusOK}, Body: updateBody}
		svcUpd := NewRealServiceFromDoer(mdUpd).(*service)
		out, apiResp, err = svcUpd.UpdateOrder(context.Background(), "1", UpdateOrderRequest{EcomOrder: EcomOrder{ExternalID: "updated"}})
		assert.NoError(t, err)
		assert.NotNil(t, apiResp)
		assert.NotNil(t, out)
		assert.Equal(t, "updated", out.ExternalID)

		// DeleteOrder success
		mdDel := &th.MockDoer{Resp: &client.APIResponse{StatusCode: http.StatusOK}, Body: []byte(`{}`)}
		svcDel := NewRealServiceFromDoer(mdDel).(*service)
		apiResp, err = svcDel.DeleteOrder(context.Background(), "1")
		assert.NoError(t, err)
		assert.NotNil(t, apiResp)

		// ListOrders success
		listBody := []byte(`{"ecomOrders":[{"id":"1"}],"meta":{"total":"1"}}`)
		mdList := &th.MockDoer{Resp: &client.APIResponse{StatusCode: http.StatusOK}, Body: listBody}
		svcList := NewRealServiceFromDoer(mdList).(*service)
		listOut, apiResp, err := svcList.ListOrders(context.Background(), map[string]string{"externalid": "1"})
		assert.NoError(t, err)
		assert.NotNil(t, apiResp)
		// wrapper should be an EcomOrderListResponse
		assert.NotNil(t, listOut)
		assert.Len(t, listOut.EcomOrders, 1)

		// Doer error paths for Get
		mdErr := &th.MockDoer{Err: errors.New("boom")}
		svcErr := NewRealServiceFromDoer(mdErr).(*service)
		_, _, err = svcErr.GetOrder(context.Background(), "1")
		assert.Error(t, err)
		_, _, err = svcErr.UpdateOrder(context.Background(), "1", UpdateOrderRequest{})
		assert.Error(t, err)
		_, err = svcErr.DeleteOrder(context.Background(), "1")
		assert.Error(t, err)
		_, _, err = svcErr.ListOrders(context.Background(), nil)
		assert.Error(t, err)
	})
}
