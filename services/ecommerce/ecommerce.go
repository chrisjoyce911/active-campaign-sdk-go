package ecommerce

import (
	"context"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

type service struct {
	client client.Doer
}

// EcommerceService is the public interface for ecommerce actions.
type EcommerceService interface {
	CreateOrder(ctx context.Context, req CreateOrderRequest) (*EcomOrder, *client.APIResponse, error)
	GetOrder(ctx context.Context, id string) (*EcomOrder, *client.APIResponse, error)
	UpdateOrder(ctx context.Context, id string, req UpdateOrderRequest) (*EcomOrder, *client.APIResponse, error)
	DeleteOrder(ctx context.Context, id string) (*client.APIResponse, error)
	ListOrders(ctx context.Context, opts map[string]string) (*EcomOrderListResponse, *client.APIResponse, error)
	// Customers
	CreateCustomer(ctx context.Context, req CreateCustomerRequest) (*EcomCustomer, *client.APIResponse, error)
	GetCustomer(ctx context.Context, id string) (*EcomCustomer, *client.APIResponse, error)
	UpdateCustomer(ctx context.Context, id string, req UpdateCustomerRequest) (*EcomCustomer, *client.APIResponse, error)
	DeleteCustomer(ctx context.Context, id string) (*client.APIResponse, error)
	ListCustomers(ctx context.Context, opts map[string]string) (*EcomCustomerListResponse, *client.APIResponse, error)
	// Order products
	ListEcomOrderProducts(ctx context.Context, opts map[string]string) (*EcomOrderProductsListResponse, *client.APIResponse, error)
	ListEcomOrderProductsForOrder(ctx context.Context, orderID string, opts map[string]string) (*EcomOrderProductsListResponse, *client.APIResponse, error)
	GetEcomOrderProduct(ctx context.Context, id string) (*EcomOrderProduct, *client.APIResponse, error)
}
