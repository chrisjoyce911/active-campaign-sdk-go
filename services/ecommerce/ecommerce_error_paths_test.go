package ecommerce

import (
	"context"
	"errors"
	"testing"

	th "github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestCustomersAndProducts_DoerErrors(t *testing.T) {
	mdErr := &th.MockDoer{Err: errors.New("boom")}
	svc := NewRealServiceFromDoer(mdErr).(*service)

	// Customers
	_, _, err := svc.CreateCustomer(context.Background(), CreateCustomerRequest{})
	assert.Error(t, err)
	_, _, err = svc.GetCustomer(context.Background(), "1")
	assert.Error(t, err)
	_, _, err = svc.UpdateCustomer(context.Background(), "1", UpdateCustomerRequest{})
	assert.Error(t, err)
	_, err = svc.DeleteCustomer(context.Background(), "1")
	assert.Error(t, err)
	_, _, err = svc.ListCustomers(context.Background(), nil)
	assert.Error(t, err)

	// Order-products
	_, _, err = svc.ListEcomOrderProducts(context.Background(), nil)
	assert.Error(t, err)
	_, _, err = svc.ListEcomOrderProductsForOrder(context.Background(), "1", nil)
	assert.Error(t, err)
	_, _, err = svc.GetEcomOrderProduct(context.Background(), "1")
	assert.Error(t, err)
}
