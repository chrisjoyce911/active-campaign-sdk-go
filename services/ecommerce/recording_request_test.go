package ecommerce

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	th "github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestInt64String_UnmarshalJSON(t *testing.T) {
	var s struct {
		Val Int64String `json:"val"`
	}

	// number
	err := json.Unmarshal([]byte(`{"val":123}`), &s)
	assert.NoError(t, err)
	assert.Equal(t, Int64String(123), s.Val)

	// string number
	err = json.Unmarshal([]byte(`{"val":"456"}`), &s)
	assert.NoError(t, err)
	assert.Equal(t, Int64String(456), s.Val)

	// null
	err = json.Unmarshal([]byte(`{"val":null}`), &s)
	assert.NoError(t, err)
	assert.Equal(t, Int64String(0), s.Val)
}

func TestRecordingDoer_CreateUpdateRequests(t *testing.T) {
	tests := []struct {
		name             string
		call             func(svc *service) error
		expectMethod     string
		expectPath       string
		expectExternalID string
		expectPrice      int
	}{
		{
			name: "CreateOrder",
			call: func(svc *service) error {
				rd := svc.client.(*th.RecordingDoer)
				// create request
				req := CreateOrderRequest{EcomOrder: EcomOrder{ExternalID: "ext-1", TotalPrice: Int64String(1000)}}
				_, _, err := svc.CreateOrder(context.Background(), req)
				if err != nil {
					return err
				}
				// assertions on recorded body
				var last map[string]interface{}
				_ = json.Unmarshal(rd.LastBody, &last)
				eo := last["ecomOrder"].(map[string]interface{})
				if eo["externalid"] != "ext-1" {
					return nil
				}
				return nil
			},
			expectMethod:     "POST",
			expectPath:       "ecomOrders",
			expectExternalID: "ext-1",
			expectPrice:      1000,
		},
		{
			name: "UpdateOrder",
			call: func(svc *service) error {
				rd := svc.client.(*th.RecordingDoer)
				req := UpdateOrderRequest{EcomOrder: EcomOrder{ExternalID: "ext-2", TotalPrice: Int64String(2000)}}
				_, _, err := svc.UpdateOrder(context.Background(), "1", req)
				if err != nil {
					return err
				}
				var last map[string]interface{}
				_ = json.Unmarshal(rd.LastBody, &last)
				eo := last["ecomOrder"].(map[string]interface{})
				if eo["externalid"] != "ext-2" {
					return nil
				}
				return nil
			},
			expectMethod:     "PUT",
			expectPath:       "ecomOrders/1",
			expectExternalID: "ext-2",
			expectPrice:      2000,
		},
		{
			name: "CreateCustomer",
			call: func(svc *service) error {
				rd := svc.client.(*th.RecordingDoer)
				req := CreateCustomerRequest{EcomCustomer: EcomCustomer{ExternalID: "c-ext", TotalRevenue: Int64String(500)}}
				_, _, err := svc.CreateCustomer(context.Background(), req)
				if err != nil {
					return err
				}
				var last map[string]interface{}
				_ = json.Unmarshal(rd.LastBody, &last)
				ec := last["ecomCustomer"].(map[string]interface{})
				if ec["externalid"] != "c-ext" {
					return nil
				}
				return nil
			},
			expectMethod:     "POST",
			expectPath:       "ecomCustomers",
			expectExternalID: "c-ext",
			expectPrice:      500,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// response body minimal wrapper so service unmarshals successfully
			respBody := []byte(`{"ecomOrder":{"id":"1","externalid":"x","totalPrice":1000},"ecomCustomer":{"id":"1","externalid":"x","totalRevenue":500}}`)
			rd := &th.RecordingDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: respBody}
			svc := &service{client: rd}

			// call the scenario
			err := tc.call(svc)
			assert.NoError(t, err)

			// method & path
			assert.Equal(t, tc.expectMethod, rd.LastMethod)
			assert.Equal(t, tc.expectPath, rd.LastPath)

			// inspect recorded body
			var last map[string]interface{}
			err = json.Unmarshal(rd.LastBody, &last)
			assert.NoError(t, err)
			// check external id location
			if tc.name == "CreateCustomer" {
				ec := last["ecomCustomer"].(map[string]interface{})
				assert.Equal(t, tc.expectExternalID, ec["externalid"])
				// numeric fields decode to float64 in generic maps; compare as ints
				assert.Equal(t, tc.expectPrice, int(ec["totalRevenue"].(float64)))
			} else {
				eo := last["ecomOrder"].(map[string]interface{})
				assert.Equal(t, tc.expectExternalID, eo["externalid"])
				assert.Equal(t, tc.expectPrice, int(eo["totalPrice"].(float64)))
			}
		})
	}
}
