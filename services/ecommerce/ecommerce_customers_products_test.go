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

func TestCustomersEndpoints(t *testing.T) {
	createBody := []byte(`{"ecomCustomer":{"id":"10","externalid":"c-ext"}}`)
	md := &th.MockDoer{Resp: &client.APIResponse{StatusCode: http.StatusCreated}, Body: createBody}
	svc := NewRealServiceFromDoer(md).(*service)

	cust, apiResp, err := svc.CreateCustomer(context.Background(), CreateCustomerRequest{EcomCustomer: EcomCustomer{ExternalID: "c-ext"}})
	assert.NoError(t, err)
	assert.NotNil(t, apiResp)
	assert.Equal(t, "10", cust.ID)

	// Get
	getBody := []byte(`{"ecomCustomer":{"id":"10","email":"a@b.com"}}`)
	mdGet := &th.MockDoer{Resp: &client.APIResponse{StatusCode: http.StatusOK}, Body: getBody}
	svcGet := NewRealServiceFromDoer(mdGet).(*service)
	cust, apiResp, err = svcGet.GetCustomer(context.Background(), "10")
	assert.NoError(t, err)
	assert.Equal(t, "10", cust.ID)

	// Update
	updBody := []byte(`{"ecomCustomer":{"id":"10","email":"x@y.com"}}`)
	mdUpd := &th.MockDoer{Resp: &client.APIResponse{StatusCode: http.StatusOK}, Body: updBody}
	svcUpd := NewRealServiceFromDoer(mdUpd).(*service)
	cust, apiResp, err = svcUpd.UpdateCustomer(context.Background(), "10", UpdateCustomerRequest{EcomCustomer: EcomCustomer{Email: "x@y.com"}})
	assert.NoError(t, err)
	assert.Equal(t, "10", cust.ID)

	// Delete
	mdDel := &th.MockDoer{Resp: &client.APIResponse{StatusCode: http.StatusOK}, Body: []byte(`{}`)}
	svcDel := NewRealServiceFromDoer(mdDel).(*service)
	apiResp, err = svcDel.DeleteCustomer(context.Background(), "10")
	assert.NoError(t, err)
	assert.NotNil(t, apiResp)

	// List
	listBody := []byte(`{"ecomCustomers":[{"id":"10"}],"meta":{"total":"1"}}`)
	mdList := &th.MockDoer{Resp: &client.APIResponse{StatusCode: http.StatusOK}, Body: listBody}
	svcList := NewRealServiceFromDoer(mdList).(*service)
	out, apiResp, err := svcList.ListCustomers(context.Background(), map[string]string{"email": "a@b.com"})
	assert.NoError(t, err)
	assert.NotNil(t, out)
	assert.Len(t, out.EcomCustomers, 1)

	// error path
	mdErr := &th.MockDoer{Err: errors.New("boom")}
	svcErr := NewRealServiceFromDoer(mdErr).(*service)
	_, _, err = svcErr.CreateCustomer(context.Background(), CreateCustomerRequest{})
	assert.Error(t, err)
}

func TestOrderProductsEndpoints(t *testing.T) {
	// List products
	listBody := []byte(`{"ecomOrderProducts":[{"externalid":"P1"}],"meta":{"total":"1"}}`)
	md := &th.MockDoer{Resp: &client.APIResponse{StatusCode: http.StatusOK}, Body: listBody}
	svc := NewRealServiceFromDoer(md).(*service)
	out, apiResp, err := svc.ListEcomOrderProducts(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, apiResp)
	assert.Len(t, out.EcomOrderProducts, 1)

	// List for order
	mdOrder := &th.MockDoer{Resp: &client.APIResponse{StatusCode: http.StatusOK}, Body: listBody}
	svcOrder := NewRealServiceFromDoer(mdOrder).(*service)
	out2, apiResp, err := svcOrder.ListEcomOrderProductsForOrder(context.Background(), "1", nil)
	assert.NoError(t, err)
	assert.Len(t, out2.EcomOrderProducts, 1)

	// Get product
	getBody := []byte(`{"ecomOrderProduct":{"externalid":"P1","name":"N"}}`)
	mdGet := &th.MockDoer{Resp: &client.APIResponse{StatusCode: http.StatusOK}, Body: getBody}
	svcGet := NewRealServiceFromDoer(mdGet).(*service)
	p, apiResp, err := svcGet.GetEcomOrderProduct(context.Background(), "1")
	assert.NoError(t, err)
	assert.Equal(t, "P1", p.ExternalID)

	// error path
	mdErr := &th.MockDoer{Err: errors.New("boom")}
	svcErr := NewRealServiceFromDoer(mdErr).(*service)
	_, _, err = svcErr.ListEcomOrderProducts(context.Background(), nil)
	assert.Error(t, err)
}
