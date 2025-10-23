package ecommerce

import (
	"encoding/json"
	"strconv"
)

// EcomOrder represents an e-commerce order resource. Fields intentionally
// keep string types for many numeric-like values because the API often
// returns numbers as strings in examples. Adjust types later if stricter
// parsing is desired.
type EcomOrder struct {
	ID                  string                 `json:"id,omitempty"`
	ExternalID          string                 `json:"externalid,omitempty"`
	Source              string                 `json:"source,omitempty"`
	Email               string                 `json:"email,omitempty"`
	OrderProducts       []EcomOrderProduct     `json:"orderProducts,omitempty"`
	OrderDiscounts      []EcomOrderDiscount    `json:"orderDiscounts,omitempty"`
	OrderURL            string                 `json:"orderUrl,omitempty"`
	ExternalCreatedDate string                 `json:"externalCreatedDate,omitempty"`
	ExternalUpdatedDate string                 `json:"externalUpdatedDate,omitempty"`
	ShippingMethod      string                 `json:"shippingMethod,omitempty"`
	TotalPrice          Int64String            `json:"totalPrice,omitempty"`
	ShippingAmount      Int64String            `json:"shippingAmount,omitempty"`
	TaxAmount           Int64String            `json:"taxAmount,omitempty"`
	DiscountAmount      Int64String            `json:"discountAmount,omitempty"`
	Currency            string                 `json:"currency,omitempty"`
	OrderNumber         string                 `json:"orderNumber,omitempty"`
	ConnectionID        string                 `json:"connectionid,omitempty"`
	CustomerID          string                 `json:"customerid,omitempty"`
	CreatedDate         string                 `json:"createdDate,omitempty"`
	UpdatedDate         string                 `json:"updatedDate,omitempty"`
	State               interface{}            `json:"state,omitempty"`
	TotalProducts       Int64String            `json:"totalProducts,omitempty"`
	Tstamp              string                 `json:"tstamp,omitempty"`
	Links               map[string]interface{} `json:"links,omitempty"`
	Connection          interface{}            `json:"connection,omitempty"`
}

type EcomOrderProduct struct {
	ExternalID  string      `json:"externalid,omitempty"`
	Name        string      `json:"name,omitempty"`
	Price       Int64String `json:"price,omitempty"`
	Quantity    Int64String `json:"quantity,omitempty"`
	Category    string      `json:"category,omitempty"`
	SKU         string      `json:"sku,omitempty"`
	Description string      `json:"description,omitempty"`
	ImageURL    string      `json:"imageUrl,omitempty"`
	ProductURL  string      `json:"productUrl,omitempty"`
}

type EcomOrderDiscount struct {
	Name           string      `json:"name,omitempty"`
	Type           string      `json:"type,omitempty"`
	DiscountAmount Int64String `json:"discountAmount,omitempty"`
}

type EcomOrderListResponse struct {
	EcomOrders []EcomOrder       `json:"ecomOrders"`
	Meta       map[string]string `json:"meta,omitempty"`
}

// Customers
type EcomCustomer struct {
	ID                 string                 `json:"id,omitempty"`
	ConnectionID       string                 `json:"connectionid,omitempty"`
	ExternalID         string                 `json:"externalid,omitempty"`
	Email              string                 `json:"email,omitempty"`
	TotalRevenue       Int64String            `json:"totalRevenue,omitempty"`
	TotalOrders        Int64String            `json:"totalOrders,omitempty"`
	TotalProducts      Int64String            `json:"totalProducts,omitempty"`
	AvgRevenuePerOrder Int64String            `json:"avgRevenuePerOrder,omitempty"`
	AvgProductCategory string                 `json:"avgProductCategory,omitempty"`
	Tstamp             string                 `json:"tstamp,omitempty"`
	AcceptsMarketing   string                 `json:"acceptsMarketing,omitempty"`
	Links              map[string]interface{} `json:"links,omitempty"`
}

type EcomCustomerListResponse struct {
	EcomCustomers []EcomCustomer    `json:"ecomCustomers"`
	Meta          map[string]string `json:"meta,omitempty"`
}

// Order products list wrapper
type EcomOrderProductsListResponse struct {
	EcomOrderProducts []EcomOrderProduct `json:"ecomOrderProducts"`
	Meta              map[string]string  `json:"meta,omitempty"`
}

// Int64String unmarshals a JSON number or string into an int64.
type Int64String int64

func (i *Int64String) UnmarshalJSON(b []byte) error {
	if len(b) == 0 || string(b) == "null" {
		*i = 0
		return nil
	}
	if b[0] == '"' {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		if s == "" {
			*i = 0
			return nil
		}
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		*i = Int64String(v)
		return nil
	}
	var v int64
	if err := json.Unmarshal(b, &v); err == nil {
		*i = Int64String(v)
		return nil
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if s == "" {
		*i = 0
		return nil
	}
	vi, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*i = Int64String(vi)
	return nil
}

// Request wrapper types
type CreateOrderRequest struct {
	EcomOrder EcomOrder `json:"ecomOrder"`
}
type UpdateOrderRequest = CreateOrderRequest

type CreateCustomerRequest struct {
	EcomCustomer EcomCustomer `json:"ecomCustomer"`
}
type UpdateCustomerRequest = CreateCustomerRequest
