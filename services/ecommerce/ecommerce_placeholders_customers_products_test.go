package ecommerce

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomersAndProducts_Placeholders(t *testing.T) {
	s := &service{}

	// Customers placeholders
	_, _, err := s.CreateCustomer(context.Background(), CreateCustomerRequest{})
	assert.Error(t, err)
	_, _, err = s.GetCustomer(context.Background(), "1")
	assert.Error(t, err)
	_, _, err = s.UpdateCustomer(context.Background(), "1", UpdateCustomerRequest{})
	assert.Error(t, err)
	_, err = s.DeleteCustomer(context.Background(), "1")
	assert.Error(t, err)
	_, _, err = s.ListCustomers(context.Background(), nil)
	assert.Error(t, err)

	// Order-products placeholders
	_, _, err = s.ListEcomOrderProducts(context.Background(), nil)
	assert.Error(t, err)
	_, _, err = s.ListEcomOrderProductsForOrder(context.Background(), "1", nil)
	assert.Error(t, err)
	_, _, err = s.GetEcomOrderProduct(context.Background(), "1")
	assert.Error(t, err)
}
