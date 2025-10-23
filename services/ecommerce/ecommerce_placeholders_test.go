package ecommerce

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/stretchr/testify/assert"
)

func TestEcommerce_Placeholders_OtherMethods(t *testing.T) {
	s := &service{}

	// GetOrder placeholder
	_, _, err := s.GetOrder(context.Background(), "1")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ecommerce service not configured")

	// UpdateOrder placeholder
	_, _, err = s.UpdateOrder(context.Background(), "1", UpdateOrderRequest{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ecommerce service not configured")

	// DeleteOrder placeholder
	_, err = s.DeleteOrder(context.Background(), "1")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ecommerce service not configured")

	// ListOrders placeholder
	_, _, err = s.ListOrders(context.Background(), nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ecommerce service not configured")
}

func TestNewRealServiceConstructor(t *testing.T) {
	// NewCoreClient should construct a CoreClient even for a simple base URL.
	c, err := client.NewCoreClient("http://example.local/", "tok")
	assert.NoError(t, err)
	svc := NewRealService(c)
	// The returned value should be non-nil and convertible to *service
	s, ok := svc.(*service)
	assert.True(t, ok)
	assert.NotNil(t, s.client)
}
