# services/ecommerce

E-Commerce Customers
E-commerce customer resources represent a customer in an external e-commerce service such as Shopify. Customer resources primarily hold aggregate e-commerce data associated with a contact including the total revenue, total number of orders, and total number of products ordered (see the table below). This data cannot be saved to a customer object directly, but will be updated when order resources are created or updated for a customer. Note that a customer is related to a contact by the email address.

Field

Description

totalRevenue

The total revenue amount for the customer.

totalOrders

The total number of orders placed by the customer.

totalProducts

The total number of products ordered by the customer.

avgRevenuePerOrder

The average revenue per order for the customer.

avgProductCategory

The most frequent product category ordered by a customer.

🚧
Fields representing a monetary value are formatted in the lowest unit of the currency. For example, $32.80 would be represented as 3280.

The API allows you to create, update, and delete customer resources. You can retrieve individual customers as well as a list of all customers. Before you can create any customers, you must have created a connection resource for the e-commerce service.

## Create a customer

POST https://{youraccountname}.api-us1.com/api/3/ecomCustomers
Create a new e-commerce customer resource.

POST /ecomCustomers (Example REQUEST)

```json
{
  "ecomCustomer": {
    "connectionid": "1",
    "externalid": "56789",
    "email": "alice@example.com",
    "acceptsMarketing": "1"
  }
}
```

POST /ecomCustomers (Example RESPONSE)

```json
{
  "ecomCustomer": {
    "connectionid": "1",
    "externalid": "56789",
    "email": "alice@example.com",
    "links": {
      "connection": "/api/3/ecomCustomers/1/connection",
      "orders": "/api/3/ecomCustomers/1/orders"
    },
    "id": "1",
    "connection": "1"
  }
}
```

```bash
curl --request POST \
 --url https://youraccountname.api-us1.com/api/3/ecomCustomers \
 --header 'accept: application/json' \
 --header 'content-type: application/json' \
 --data '
{
"ecomCustomer": {
"connectionid": "1",
"externalid": "56789",
"email": "alice@example.com",
"acceptsMarketing": "1"
}
}
'
```

201 OK

```json
{
  "ecomCustomer": {
    "connectionid": "1",
    "externalid": "56789",
    "email": "alice@example.com",
    "links": {
      "connection": "/api/3/ecomCustomers/1/connection",
      "orders": "/api/3/ecomCustomers/1/orders"
    },
    "id": "1",
    "connection": "1"
  }
}
```

400 Bad Request

```json
{}
```

## Retrieve a customer

GET https://{youraccountname}.api-us1.com/api/3/ecomCustomers/{id}
Retrieve an existing e-commerce customer resource.

```json
{
  "ecomCustomer": {
    "connectionid": "3",
    "externalid": "1440117293120",
    "email": "john.doe@example.com",
    "totalRevenue": "150761",
    "totalOrders": "1",
    "totalProducts": "3",
    "avgRevenuePerOrder": "150761",
    "avgProductCategory": "Musical Instrument",
    "tstamp": "2019-01-01T10:23:22-06:00",
    "acceptsMarketing": "0",
    "links": {
      "connection": "https://:account.api-us1.com/api/:version/ecomCustomers/15/connection",
      "orders": "https://:account.api-us1.com/api/:version/ecomCustomers/15/orders"
    },
    "id": "15",
    "connection": "3"
  }
}
```

```bash
curl --request GET \
 --url https://youraccountname.api-us1.com/api/3/ecomCustomers/id \
 --header 'accept: application/json'
```

200 OK

```json
{
  "ecomCustomer": {
    "connectionid": "3",
    "externalid": "1440117293120",
    "email": "john.doe@example.com",
    "totalRevenue": "150761",
    "totalOrders": "1",
    "totalProducts": "3",
    "avgRevenuePerOrder": "150761",
    "avgProductCategory": "Musical Instrument",
    "tstamp": "2019-01-01T10:23:22-06:00",
    "acceptsMarketing": "0",
    "links": {
      "connection": "https://:account.api-us1.com/api/:version/ecomCustomers/15/connection",
      "orders": "https://:account.api-us1.com/api/:version/ecomCustomers/15/orders"
    },
    "id": "15",
    "connection": "3"
  }
}
```

400 Bad Request

```json
{}
```

## Update a customer

PUT https://{youraccountname}.api-us1.com/api/3/ecomCustomers/{id}
Update an existing e-commerce customer resource.

```json
{
  "ecomCustomer": {
    "externalid": "98765"
  }
}
```

ecomCustomer object
externalid
string
The id of the customer in the external service.

98765
connectionid
string
The id of the connection object for the service where the customer originates.

connectionid
email
string
The email address of the customer.

email
acceptsMarketing
string
Indication of whether customer has opt-ed in to marketing communications. 0 = not opted-in, 1 = opted-in. A value of 0 means the contact will match the "Has not opted in to marketing" segment condition and a value of 1 means the contact will match the "Has opted in to marketing" segment condition.

1

```bash
curl --request PUT \
 --url https://youraccountname.api-us1.com/api/3/ecomCustomers/id \
 --header 'accept: application/json' \
 --header 'content-type: application/json' \
 --data '
{
"ecomCustomer": {
"externalid": "98765",
"email": "email",
"connectionid": "connectionid",
"acceptsMarketing": "1"
}
}
'
```

200 OK

```json
{
  "ecomCustomer": {
    "connectionid": "1",
    "externalid": "98765",
    "email": "alice@example.com",
    "totalRevenue": "3280",
    "totalOrders": "2",
    "totalProducts": "2",
    "avgRevenuePerOrder": "2285",
    "avgProductCategory": "Electronics",
    "tstamp": "2017-02-06T14:05:31-06:00",
    "links": {
      "connection": "/api/3/ecomCustomers/1/connection",
      "orders": "/api/3/ecomCustomers/1/orders"
    },
    "id": "1",
    "connection": "1"
  }
}
```

## Delete a customer

DELETE https://{youraccountname}.api-us1.com/api/3/ecomCustomers/{id}
Delete an existing e-commerce customer resource.

````bash
curl --request DELETE \
 --url https://youraccountname.api-us1.com/api/3/ecomCustomers/id \
 --header 'accept: application/json'

200 OK
```json
{}
````

400 Bad Request

```json
{}
```

List all customers
GET https://{youraccountname}.api-us1.com/api/3/ecomCustomers
List all e-commerce customer resources.

Query Params
filters[email]
string
Filter by the email address of a customer.

filters[externalid]
string
Filter by the id of the customer in the external service.

filters[connectionid]
string
Filter by the id of the connection object for the service where the customer originates.

```json
{
  "ecomCustomers": [
    {
      "connectionid": "1",
      "externalid": "56789",
      "email": "alice@example.com",
      "totalRevenue": "3280",
      "totalOrders": "2",
      "totalProducts": "2",
      "avgRevenuePerOrder": "2285",
      "avgProductCategory": "Electronics",
      "tstamp": "2017-02-06T14:05:31-06:00",
      "links": {
        "connection": "/api/3/ecomCustomers/1/connection",
        "orders": "/api/3/ecomCustomers/1/orders"
      },
      "id": "1",
      "connection": "1"
    },
    {
      "connectionid": "2",
      "externalid": "44322",
      "email": "alice@example.com",
      "totalRevenue": "7599",
      "totalOrders": "1",
      "totalProducts": "1",
      "avgRevenuePerOrder": "7599",
      "avgProductCategory": "Books",
      "tstamp": "2016-12-13T18:02:07-06:00",
      "links": {
        "connection": "/api/3/ecomCustomers/3/connection",
        "orders": "/api/3/ecomCustomers/3/orders"
      },
      "id": "3",
      "connection": "2"
    },
    {
      "connectionid": "0",
      "externalid": "0",
      "email": "alice@example.com",
      "totalRevenue": "10879",
      "totalOrders": "3",
      "totalProducts": "3",
      "avgRevenuePerOrder": "3626",
      "avgProductCategory": "Electronics",
      "tstamp": "2017-02-06T14:05:31-06:00",
      "links": {
        "connection": "/api/3/ecomCustomers/2/connection",
        "orders": "/api/3/ecomCustomers/2/orders"
      },
      "id": "2",
      "connection": null
    }
  ],
  "meta": {
    "total": "3"
  }
}
```

200 OK

```json
{
  "ecomCustomers": [
    {
      "connectionid": "1",
      "externalid": "56789",
      "email": "alice@example.com",
      "totalRevenue": "3280",
      "totalOrders": "2",
      "totalProducts": "2",
      "avgRevenuePerOrder": "2285",
      "avgProductCategory": "Electronics",
      "tstamp": "2017-02-06T14:05:31-06:00",
      "links": {
        "connection": "/api/3/ecomCustomers/1/connection",
        "orders": "/api/3/ecomCustomers/1/orders"
      },
      "id": "1",
      "connection": "1"
    },
    {
      "connectionid": "2",
      "externalid": "44322",
      "email": "alice@example.com",
      "totalRevenue": "7599",
      "totalOrders": "1",
      "totalProducts": "1",
      "avgRevenuePerOrder": "7599",
      "avgProductCategory": "Books",
      "tstamp": "2016-12-13T18:02:07-06:00",
      "links": {
        "connection": "/api/3/ecomCustomers/3/connection",
        "orders": "/api/3/ecomCustomers/3/orders"
      },
      "id": "3",
      "connection": "2"
    },
    {
      "connectionid": "0",
      "externalid": "0",
      "email": "alice@example.com",
      "totalRevenue": "10879",
      "totalOrders": "3",
      "totalProducts": "3",
      "avgRevenuePerOrder": "3626",
      "avgProductCategory": "Electronics",
      "tstamp": "2017-02-06T14:05:31-06:00",
      "links": {
        "connection": "/api/3/ecomCustomers/2/connection",
        "orders": "/api/3/ecomCustomers/2/orders"
      },
      "id": "2",
      "connection": null
    }
  ],
  "meta": {
    "total": "3"
  }
}
```

E-Commerce Orders
E-Commerce order resources represent orders in an external e-commerce service such as Shopify. The API allows you to create, update, and delete order resources. You can retrieve individual orders as well as a list of all orders. Before you can create any orders, you must have created a connection resource for the e-commerce service and a customer resource for the customer who placed the order.

Orders can be created from two primary sources: real-time webhooks/events and historical syncs. Orders should only be marked as “real-time” if the order data is transmitted at the time of purchase. If the purchase occurred in the past, the order should be marked as “historical”.

Create an order
POST https://{youraccountname}.api-us1.com/api/3/ecomOrders
Create a new e-commerce order resource.

POST /ecomOrders (Example REQUEST)

```json
{
  "ecomOrder": {
    "externalid": "3246315233",
    "source": "1",
    "email": "alice@example.com",
    "orderProducts": [
      {
        "externalid": "PROD12345",
        "name": "Pogo Stick",
        "price": 4900,
        "quantity": 1,
        "category": "Toys",
        "sku": "POGO-12",
        "description": "lorem ipsum...",
        "imageUrl": "https://example.com/product.jpg",
        "productUrl": "https://store.example.com/product12345"
      },
      {
        "externalid": "PROD23456",
        "name": "Skateboard",
        "price": 3000,
        "quantity": 1,
        "category": "Toys",
        "sku": "SK8BOARD145",
        "description": "lorem ipsum...",
        "imageUrl": "https://example.com/product.jpg",
        "productUrl": "https://store.example.com/product45678"
      }
    ],
    "orderDiscounts": [
      {
        "name": "1OFF",
        "type": "order",
        "discountAmount": 100
      }
    ],
    "orderUrl": "https://example.com/orders/3246315233",
    "externalCreatedDate": "2016-09-13T17:41:39-04:00",
    "externalUpdatedDate": "2016-09-14T17:41:39-04:00",
    "shippingMethod": "UPS Ground",
    "totalPrice": 9111,
    "shippingAmount": 200,
    "taxAmount": 500,
    "discountAmount": 100,
    "currency": "USD",
    "orderNumber": "myorder-123",
    "connectionid": "1",
    "customerid": "1"
  }
}
```

POST /ecomOrders (Example RESPONSE)

```json
{
  "connections": [
    {
      "service": "example",
      "externalid": "examplestore",
      "name": "My Example Store",
      "isInternal": "0",
      "connectionType": "ecommerce",
      "status": "1",
      "syncStatus": "0",
      "sync_request_time": null,
      "sync_start_time": null,
      "lastSync": null,
      "logoUrl": "https://myexamplestore.com/images/logo.jpg",
      "linkUrl": "https://myexamplestore.com",
      "cdate": "2018-01-12T13:13:53-06:00",
      "udate": "2018-01-12T13:13:53-06:00",
      "credentialExpiration": null,
      "links": {
        "options": "https://exampleaccount.api-us1.com/api/3/connections/1/options",
        "customers": "https://exampleaccount.api-us1.com.api-us1.com/api/3/connections/1/customers"
      },
      "id": "1",
      "serviceName": "shopify"
    }
  ],
  "ecomOrderProducts": [
    {
      "externalid": "PROD12345",
      "name": "Pogo Stick",
      "price": 4900,
      "quantity": 1,
      "category": "Toys",
      "sku": "POGO-12",
      "description": "lorem ipsum...",
      "imageUrl": "https://example.com/product.jpg",
      "productUrl": "https://store.example.com/product12345"
    },
    {
      "externalid": "PROD23456",
      "name": "Skateboard",
      "price": 3000,
      "quantity": 1,
      "category": "Toys",
      "sku": "SK8BOARD145",
      "description": "lorem ipsum...",
      "imageUrl": "https://example.com/product.jpg",
      "productUrl": "https://store.example.com/product45678"
    }
  ],
  "ecomOrderDiscounts": [
    {
      "name": "1OFF",
      "type": "order",
      "orderid": "5355",
      "discountAmount": "100",
      "id": "1",
      "createdDate": "2019-09-05T12:16:18-05:00",
      "updatedDate": "2019-09-05T12:16:18-05:00"
    }
  ],
  "ecomOrder": {
    "externalid": "3246315234",
    "source": "1",
    "email": "alice@example.com",
    "currency": "USD",
    "connectionid": "1",
    "customerid": "1",
    "orderUrl": "https://example.com/orders/3246315233",
    "shippingMethod": "UPS Ground",
    "totalPrice": 9111,
    "shippingAmount": 200,
    "taxAmount": 500,
    "discountAmount": 100,
    "externalCreatedDate": "2016-09-13T16:41:39-05:00",
    "totalProducts": 2,
    "createdDate": "2019-09-05T12:16:18-05:00",
    "updatedDate": "2019-09-05T12:16:18-05:00",
    "state": 1,
    "connection": "1",
    "orderProducts": ["1", "2"],
    "orderDiscounts": ["1"],
    "customer": "1",
    "orderDate": "2016-09-13T16:41:39-05:00",
    "tstamp": "2019-09-05T12:16:18-05:00",
    "links": {
      "connection": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/connection",
      "customer": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/customer",
      "orderProducts": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/orderProducts",
      "orderDiscounts": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/orderDiscounts",
      "orderActivities": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/orderActivities"
    },
    "id": "1"
  }
}
```

```bash
curl --request POST \
 --url https://youraccountname.api-us1.com/api/3/ecomOrders \
 --header 'accept: application/json' \
 --header 'content-type: application/json' \
 --data '
{
"ecomOrder": {
"externalid": "3246315233",
"source": "1",
"email": "alice@example.com",
"orderProducts": [
{
"externalid": "PROD12345",
"name": "Pogo Stick",
"price": 4900,
"quantity": 1,
"category": "Toys",
"sku": "POGO-12",
"description": "lorem ipsum...",
"imageUrl": "https://example.com/product.jpg",
"productUrl": "https://store.example.com/product12345"
},
{
"externalid": "PROD23456",
"name": "Skateboard",
"price": 3000,
"quantity": 1,
"category": "Toys",
"sku": "SK8BOARD145",
"description": "lorem ipsum...",
"imageUrl": "https://example.com/product.jpg",
"productUrl": "https://store.example.com/product45678"
}
],
"orderDiscounts": [
{
"name": "1OFF",
"type": "order",
"discountAmount": 100
}
],
"orderUrl": "https://example.com/orders/3246315233",
"externalCreatedDate": "2016-09-13T17:41:39-04:00",
"externalUpdatedDate": "2016-09-14T17:41:39-04:00",
"shippingMethod": "UPS Ground",
"totalPrice": 9111,
"shippingAmount": 200,
"taxAmount": 500,
"discountAmount": 100,
"currency": "USD",
"orderNumber": "myorder-123",
"connectionid": "1",
"customerid": "1"
}
}
'
```

201

```json
{
  "connections": [
    {
      "service": "example",
      "externalid": "examplestore",
      "name": "My Example Store",
      "isInternal": "0",
      "connectionType": "ecommerce",
      "status": "1",
      "syncStatus": "0",
      "sync_request_time": null,
      "sync_start_time": null,
      "lastSync": null,
      "logoUrl": "https://myexamplestore.com/images/logo.jpg",
      "linkUrl": "https://myexamplestore.com",
      "cdate": "2018-01-12T13:13:53-06:00",
      "udate": "2018-01-12T13:13:53-06:00",
      "credentialExpiration": null,
      "links": {
        "options": "https://exampleaccount.api-us1.com/api/3/connections/1/options",
        "customers": "https://exampleaccount.api-us1.com.api-us1.com/api/3/connections/1/customers"
      },
      "id": "1",
      "serviceName": "shopify"
    }
  ],
  "ecomOrderProducts": [
    {
      "externalid": "PROD12345",
      "name": "Pogo Stick",
      "price": 4900,
      "quantity": 1,
      "category": "Toys",
      "sku": "POGO-12",
      "description": "lorem ipsum...",
      "imageUrl": "https://example.com/product.jpg",
      "productUrl": "https://store.example.com/product12345"
    },
    {
      "externalid": "PROD23456",
      "name": "Skateboard",
      "price": 3000,
      "quantity": 1,
      "category": "Toys",
      "sku": "SK8BOARD145",
      "description": "lorem ipsum...",
      "imageUrl": "https://example.com/product.jpg",
      "productUrl": "https://store.example.com/product45678"
    }
  ],
  "ecomOrderDiscounts": [
    {
      "name": "1OFF",
      "type": "order",
      "orderid": "5355",
      "discountAmount": "100",
      "id": "1",
      "createdDate": "2019-09-05T12:16:18-05:00",
      "updatedDate": "2019-09-05T12:16:18-05:00"
    }
  ],
  "ecomOrder": {
    "externalid": "3246315234",
    "source": "1",
    "email": "alice@example.com",
    "currency": "USD",
    "connectionid": "1",
    "customerid": "1",
    "orderUrl": "https://example.com/orders/3246315233",
    "shippingMethod": "UPS Ground",
    "totalPrice": 9111,
    "shippingAmount": 200,
    "taxAmount": 500,
    "discountAmount": 100,
    "externalCreatedDate": "2016-09-13T16:41:39-05:00",
    "totalProducts": 2,
    "createdDate": "2019-09-05T12:16:18-05:00",
    "updatedDate": "2019-09-05T12:16:18-05:00",
    "state": 1,
    "connection": "1",
    "orderProducts": ["1", "2"],
    "orderDiscounts": ["1"],
    "customer": "1",
    "orderDate": "2016-09-13T16:41:39-05:00",
    "tstamp": "2019-09-05T12:16:18-05:00",
    "links": {
      "connection": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/connection",
      "customer": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/customer",
      "orderProducts": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/orderProducts",
      "orderDiscounts": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/orderDiscounts",
      "orderActivities": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/orderActivities"
    },
    "id": "1"
  }
}
```

## Retrieve an order

GET https://{youraccountname}.api-us1.com/api/3/ecomOrders/{ecomOrderId}
Retrieve an existing e-commerce order resource.

```json
{
  "ecomOrder": {
    "customerid": "1",
    "connectionid": "1",
    "state": "1",
    "source": "1",
    "externalid": "1233456789",
    "externalcheckoutid": null,
    "orderNumber": "123456789",
    "email": "alice@example.com",
    "totalPrice": "1000",
    "discountAmount": "100",
    "shippingAmount": "120",
    "taxAmount": "120",
    "totalProducts": "1",
    "currency": "usd",
    "shippingMethod": "ground shipping",
    "orderUrl": "https://examplestore.com/orders/123456789",
    "externalCreatedDate": "2019-06-29T14:44:49-05:00",
    "externalUpdatedDate": "2019-06-29T14:44:49-05:00",
    "abandonedDate": null,
    "createdDate": "2019-06-29T14:44:51-05:00",
    "updatedDate": "2019-06-29T14:44:51-05:00",
    "orderDate": "2019-06-29T14:44:49-05:00",
    "tstamp": "2019-06-29T14:44:51-05:00",
    "links": {
      "connection": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/connection",
      "customer": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/customer",
      "orderProducts": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/orderProducts",
      "orderDiscounts": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/orderDiscounts",
      "orderActivities": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/orderActivities"
    },
    "id": "1",
    "connection": "1",
    "customer": "1"
  }
}
```

```bash
curl --request GET \
 --url https://youraccountname.api-us1.com/api/3/ecomOrders/ecomOrderId \
 --header 'accept: application/json'
```

200 OK

```json
{
  "ecomOrder": {
    "customerid": "1",
    "connectionid": "1",
    "state": "1",
    "source": "1",
    "externalid": "1233456789",
    "externalcheckoutid": null,
    "orderNumber": "123456789",
    "email": "alice@example.com",
    "totalPrice": "1000",
    "discountAmount": "100",
    "shippingAmount": "120",
    "taxAmount": "120",
    "totalProducts": "1",
    "currency": "usd",
    "shippingMethod": "ground shipping",
    "orderUrl": "https://examplestore.com/orders/123456789",
    "externalCreatedDate": "2019-06-29T14:44:49-05:00",
    "externalUpdatedDate": "2019-06-29T14:44:49-05:00",
    "abandonedDate": null,
    "createdDate": "2019-06-29T14:44:51-05:00",
    "updatedDate": "2019-06-29T14:44:51-05:00",
    "orderDate": "2019-06-29T14:44:49-05:00",
    "tstamp": "2019-06-29T14:44:51-05:00",
    "links": {
      "connection": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/connection",
      "customer": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/customer",
      "orderProducts": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/orderProducts",
      "orderDiscounts": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/orderDiscounts",
      "orderActivities": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/orderActivities"
    },
    "id": "1",
    "connection": "1",
    "customer": "1"
  }
}
```

Update an order
PUT https://{youraccountname}.api-us1.com/api/3/ecomOrders/{id}
Update an existing ecommerce order/cart resource.

PUT /ecomOrders/:id (Example REQUEST)

```json
{
  "ecomOrder": {
    "externalid": "3246315237",
    "email": "alice@example.com",
    "orderUrl": "https://example.com/orders/3246315233",
    "orderProducts": [
      {
        "externalid": "PROD12345",
        "name": "Pogo Stick",
        "price": 4900,
        "quantity": 1,
        "category": "Toys",
        "sku": "POGO-12",
        "description": "lorem ipsum...",
        "imageUrl": "https://example.com/product.jpg",
        "productUrl": "https://store.example.com/product12345"
      },
      {
        "externalid": "PROD23456",
        "name": "Skateboard",
        "price": 3000,
        "quantity": 1,
        "category": "Toys",
        "sku": "SK8BOARD145",
        "description": "lorem ipsum...",
        "imageUrl": "https://example.com/product.jpg",
        "productUrl": "https://store.example.com/product45678"
      }
    ],
    "orderDiscounts": [
      {
        "name": "1OFF",
        "type": "order",
        "discountAmount": 100
      }
    ],
    "externalUpdatedDate": "2016-09-15T17:41:39-04:00",
    "shippingMethod": "UPS Ground",
    "totalPrice": 9111,
    "shippingAmount": 200,
    "taxAmount": 500,
    "discountAmount": 100,
    "currency": "USD",
    "orderNumber": "12345-1"
  }
}
```

PUT /ecomOrders/:id (Example RESPONSE)

```json
{
  "ecomOrderProducts": [
    {
      "orderid": "1",
      "connectionid": "1",
      "externalid": "PROD12345",
      "sku": "POGO-12",
      "name": "Pogo Stick",
      "description": "lorem ipsum...",
      "price": "4900",
      "quantity": "1",
      "category": "Toys",
      "imageUrl": "https://example.com/product.jpg",
      "productUrl": "https://store.example.com/product12345",
      "createdDate": "2019-09-05T13:55:37-05:00",
      "updatedDate": "2019-09-05T13:55:37-05:00",
      "tstamp": "2019-09-05T13:55:37-05:00",
      "links": {
        "ecomOrder": "https://youraccounthere.api-us1.com/api/3/ecomOrderProducts/1/ecomOrder"
      },
      "id": "3",
      "ecomOrder": "1"
    },
    {
      "orderid": "1",
      "connectionid": "1",
      "externalid": "PROD23456",
      "sku": "SK8BOARD145",
      "name": "Skateboard",
      "description": "lorem ipsum...",
      "price": "3000",
      "quantity": "1",
      "category": "Toys",
      "imageUrl": "https://example.com/product.jpg",
      "productUrl": "https://store.example.com/product45678",
      "createdDate": "2019-09-05T13:55:37-05:00",
      "updatedDate": "2019-09-05T13:55:37-05:00",
      "tstamp": "2019-09-05T13:55:37-05:00",
      "links": {
        "ecomOrder": "https://youraccounthere.api-us1.com/api/3/ecomOrderProducts/1/ecomOrder"
      },
      "id": "4",
      "ecomOrder": "1"
    }
  ],
  "ecomOrderDiscounts": [
    {
      "name": "1OFF",
      "type": "order",
      "orderid": "5355",
      "discountAmount": "100",
      "id": "1",
      "createdDate": "2019-09-05T12:16:18-05:00",
      "updatedDate": "2019-09-05T12:16:18-05:00"
    }
  ],
  "ecomOrder": {
    "customerid": "1",
    "connectionid": "1",
    "state": "1",
    "source": "1",
    "externalid": "3246315237",
    "orderNumber": "",
    "email": "alice@example.com",
    "totalPrice": 9111,
    "discountAmount": 100,
    "shippingAmount": 200,
    "taxAmount": 500,
    "totalProducts": 2,
    "currency": "USD",
    "shippingMethod": "UPS Ground",
    "orderUrl": "https://example.com/orders/3246315233",
    "externalCreatedDate": "2016-09-13T16:41:39-05:00",
    "externalUpdatedDate": "2016-09-15T16:41:39-05:00",
    "createdDate": "2019-09-05T12:52:13-05:00",
    "updatedDate": "2019-09-05T13:55:37-05:00",
    "orderProducts": ["3", "4"],
    "orderDiscounts": ["1"],
    "customer": "1",
    "orderDate": "2016-09-13T16:41:39-05:00",
    "tstamp": "2019-09-05T13:55:37-05:00",
    "links": {
      "connection": "https://youraccounthere.api-us1.com/api/3/ecomOrders/1/connection",
      "customer": "https://youraccounthere.api-us1.com/api/3/ecomOrders/1/customer",
      "orderProducts": "https://youraccounthere.api-us1.com/api/3/ecomOrders/1/orderProducts",
      "orderDiscounts": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/orderDiscounts",
      "orderActivities": "https://youraccounthere.api-us1.com/api/3/ecomOrders/1/orderActivities"
    },
    "id": "1",
    "connection": "1"
  }
}
```

```bash
curl --request PUT \
 --url https://youraccountname.api-us1.com/api/3/ecomOrders/id \
 --header 'accept: application/json' \
 --header 'content-type: application/json' \
 --data '
{
"ecomOrder": {
"externalid": "3246315237",
"email": "alice@example.com",
"orderProducts": [
{
"externalid": "PROD12345",
"name": "Pogo Stick",
"price": 4900,
"quantity": 1,
"category": "Toys",
"sku": "POGO-12",
"description": "lorem ipsum...",
"imageUrl": "https://example.com/product.jpg",
"productUrl": "https://store.example.com/product12345"
},
{
"externalid": "PROD23456",
"name": "Skateboard",
"price": 3000,
"quantity": 1,
"category": "Toys",
"sku": "SK8BOARD145",
"description": "lorem ipsum...",
"imageUrl": "https://example.com/product.jpg",
"productUrl": "https://store.example.com/product45678"
}
],
"orderDiscounts": [
{
"name": "1OFF",
"type": "order",
"discountAmount": 100
}
],
"externalUpdatedDate": "2016-09-15T17:41:39-04:00",
"shippingMethod": "UPS Ground",
"totalPrice": 9111,
"shippingAmount": 200,
"taxAmount": 500,
"discountAmount": 100,
"currency": "USD",
"orderNumber": "12345-1"
}
}
'
```

200 OK

```json
{
  "ecomOrderProducts": [
    {
      "orderid": "1",
      "connectionid": "1",
      "externalid": "PROD12345",
      "sku": "POGO-12",
      "name": "Pogo Stick",
      "description": "lorem ipsum...",
      "price": "4900",
      "quantity": "1",
      "category": "Toys",
      "imageUrl": "https://example.com/product.jpg",
      "productUrl": "https://store.example.com/product12345",
      "createdDate": "2019-09-05T13:55:37-05:00",
      "updatedDate": "2019-09-05T13:55:37-05:00",
      "tstamp": "2019-09-05T13:55:37-05:00",
      "links": {
        "ecomOrder": "https://youraccounthere.api-us1.com/api/3/ecomOrderProducts/1/ecomOrder"
      },
      "id": "3",
      "ecomOrder": "1"
    },
    {
      "orderid": "1",
      "connectionid": "1",
      "externalid": "PROD23456",
      "sku": "SK8BOARD145",
      "name": "Skateboard",
      "description": "lorem ipsum...",
      "price": "3000",
      "quantity": "1",
      "category": "Toys",
      "imageUrl": "https://example.com/product.jpg",
      "productUrl": "https://store.example.com/product45678",
      "createdDate": "2019-09-05T13:55:37-05:00",
      "updatedDate": "2019-09-05T13:55:37-05:00",
      "tstamp": "2019-09-05T13:55:37-05:00",
      "links": {
        "ecomOrder": "https://youraccounthere.api-us1.com/api/3/ecomOrderProducts/1/ecomOrder"
      },
      "id": "4",
      "ecomOrder": "1"
    }
  ],
  "ecomOrderDiscounts": [
    {
      "name": "1OFF",
      "type": "order",
      "orderid": "5355",
      "discountAmount": "100",
      "id": "1",
      "createdDate": "2019-09-05T12:16:18-05:00",
      "updatedDate": "2019-09-05T12:16:18-05:00"
    }
  ],
  "ecomOrder": {
    "customerid": "1",
    "connectionid": "1",
    "state": "1",
    "source": "1",
    "externalid": "3246315237",
    "orderNumber": "",
    "email": "alice@example.com",
    "totalPrice": 9111,
    "discountAmount": 100,
    "shippingAmount": 200,
    "taxAmount": 500,
    "totalProducts": 2,
    "currency": "USD",
    "shippingMethod": "UPS Ground",
    "orderUrl": "https://example.com/orders/3246315233",
    "externalCreatedDate": "2016-09-13T16:41:39-05:00",
    "externalUpdatedDate": "2016-09-15T16:41:39-05:00",
    "createdDate": "2019-09-05T12:52:13-05:00",
    "updatedDate": "2019-09-05T13:55:37-05:00",
    "orderProducts": ["3", "4"],
    "orderDiscounts": ["1"],
    "customer": "1",
    "orderDate": "2016-09-13T16:41:39-05:00",
    "tstamp": "2019-09-05T13:55:37-05:00",
    "links": {
      "connection": "https://youraccounthere.api-us1.com/api/3/ecomOrders/1/connection",
      "customer": "https://youraccounthere.api-us1.com/api/3/ecomOrders/1/customer",
      "orderProducts": "https://youraccounthere.api-us1.com/api/3/ecomOrders/1/orderProducts",
      "orderDiscounts": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/orderDiscounts",
      "orderActivities": "https://youraccounthere.api-us1.com/api/3/ecomOrders/1/orderActivities"
    },
    "id": "1",
    "connection": "1"
  }
}
```

## Delete an order

DELETE https://{youraccountname}.api-us1.com/api/3/ecomOrders/{ecomOrderId}
Delete an existing e-commerce order resource.

```bash
curl --request DELETE \
 --url https://youraccountname.api-us1.com/api/3/ecomOrders/ecomOrderId \
 --header 'accept: application/json'
```

200 OK

```json
{}
```

List all orders
GET https://{youraccountname}.api-us1.com/api/3/ecomOrders
List all existing e-commerce order resources.

```json
{
  "ecomOrders": [
    {
      "externalid": "3246315234",
      "source": "1",
      "email": "alice@example.com",
      "currency": "USD",
      "connectionid": "1",
      "customerid": "1",
      "orderUrl": "https://example.com/orders/3246315233",
      "shippingMethod": "UPS Ground",
      "totalPrice": 9111,
      "shippingAmount": 200,
      "taxAmount": 500,
      "discountAmount": 100,
      "externalCreatedDate": "2016-09-13T16:41:39-05:00",
      "totalProducts": 2,
      "createdDate": "2019-09-05T12:16:18-05:00",
      "updatedDate": "2019-09-05T12:16:18-05:00",
      "state": 1,
      "connection": "1",
      "orderProducts": ["1", "2"],
      "customer": "1",
      "orderDate": "2016-09-13T16:41:39-05:00",
      "tstamp": "2019-09-05T12:16:18-05:00",
      "links": {
        "connection": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/connection",
        "customer": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/customer",
        "orderProducts": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/orderProducts",
        "orderDiscounts": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/orderDiscounts",
        "orderActivities": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/orderActivities"
      },
      "id": "1"
    },
    {
      "externalid": "47856739866",
      "source": "1",
      "email": "example@example.com",
      "currency": "USD",
      "connectionid": "1",
      "customerid": "2",
      "orderUrl": "https://example.com/orders/47856739866",
      "shippingMethod": "UPS Ground",
      "totalPrice": 3450,
      "shippingAmount": 100,
      "taxAmount": 0,
      "discountAmount": 0,
      "externalCreatedDate": "2019-09-06T2:10:00-05:00",
      "totalProducts": 2,
      "createdDate": "2019-09-06T2:10:00-05:00",
      "updatedDate": "2019-09-06T2:10:00-05:00",
      "state": 1,
      "connection": "1",
      "orderProducts": ["3"],
      "customer": "2",
      "orderDate": "2019-09-06T2:10:00-05:00",
      "tstamp": "2019-09-06T2:10:00-05:00",
      "links": {
        "connection": "https://exampleaccount.api-us1.com/api/3/ecomOrders/2/connection",
        "customer": "https://exampleaccount.api-us1.com/api/3/ecomOrders/2/customer",
        "orderProducts": "https://exampleaccount.api-us1.com/api/3/ecomOrders/2/orderProducts",
        "orderDiscounts": "https://exampleaccount.api-us1.com/api/3/ecomOrders/2/orderDiscounts",
        "orderActivities": "https://exampleaccount.api-us1.com/api/3/ecomOrders/2/orderActivities"
      },
      "id": "2"
    }
  ],
  "meta": {
    "total": "2"
  }
}
```

```bash
curl --request GET \
 --url https://youraccountname.api-us1.com/api/3/ecomOrders \
 --header 'accept: application/json'
```

200 OK

```json
{
  "ecomOrders": [
    {
      "externalid": "3246315234",
      "source": "1",
      "email": "alice@example.com",
      "currency": "USD",
      "connectionid": "1",
      "customerid": "1",
      "orderUrl": "https://example.com/orders/3246315233",
      "shippingMethod": "UPS Ground",
      "totalPrice": 9111,
      "shippingAmount": 200,
      "taxAmount": 500,
      "discountAmount": 100,
      "externalCreatedDate": "2016-09-13T16:41:39-05:00",
      "totalProducts": 2,
      "createdDate": "2019-09-05T12:16:18-05:00",
      "updatedDate": "2019-09-05T12:16:18-05:00",
      "state": 1,
      "connection": "1",
      "orderProducts": ["1", "2"],
      "customer": "1",
      "orderDate": "2016-09-13T16:41:39-05:00",
      "tstamp": "2019-09-05T12:16:18-05:00",
      "links": {
        "connection": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/connection",
        "customer": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/customer",
        "orderProducts": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/orderProducts",
        "orderDiscounts": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/orderDiscounts",
        "orderActivities": "https://exampleaccount.api-us1.com/api/3/ecomOrders/1/orderActivities"
      },
      "id": "1"
    },
    {
      "externalid": "47856739866",
      "source": "1",
      "email": "example@example.com",
      "currency": "USD",
      "connectionid": "1",
      "customerid": "2",
      "orderUrl": "https://example.com/orders/47856739866",
      "shippingMethod": "UPS Ground",
      "totalPrice": 3450,
      "shippingAmount": 100,
      "taxAmount": 0,
      "discountAmount": 0,
      "externalCreatedDate": "2019-09-06T2:10:00-05:00",
      "totalProducts": 2,
      "createdDate": "2019-09-06T2:10:00-05:00",
      "updatedDate": "2019-09-06T2:10:00-05:00",
      "state": 1,
      "connection": "1",
      "orderProducts": ["3"],
      "customer": "2",
      "orderDate": "2019-09-06T2:10:00-05:00",
      "tstamp": "2019-09-06T2:10:00-05:00",
      "links": {
        "connection": "https://exampleaccount.api-us1.com/api/3/ecomOrders/2/connection",
        "customer": "https://exampleaccount.api-us1.com/api/3/ecomOrders/2/customer",
        "orderProducts": "https://exampleaccount.api-us1.com/api/3/ecomOrders/2/orderProducts",
        "orderDiscounts": "https://exampleaccount.api-us1.com/api/3/ecomOrders/2/orderDiscounts",
        "orderActivities": "https://exampleaccount.api-us1.com/api/3/ecomOrders/2/orderActivities"
      },
      "id": "2"
    }
  ],
  "meta": {
    "total": "2"
  }
}
```

422 Unprocessable Entity

```json
{}
```

# E-Commerce Order Products

E-Commerce Order Product resources represent a combination of line items and products in an external e-commerce service such as Shopify. The API allows you to query these resources, but to create them you must create them as part of an order.

## List EcomOrderProducts

GET https://{youraccountname}.api-us1.com/api/3/ecomOrderProducts

```json
{
  "ecomOrderProducts": [
    {
      "orderid": "1",
      "connectionid": "1",
      "externalid": "1269385775",
      "sku": "SKATE-8",
      "name": "My Cool Skateboard",
      "description": "This skateboard is so cool, it rides itself",
      "price": "5000",
      "quantity": "1",
      "category": "skateboards",
      "imageUrl": "https://example.com/images/skateboard.jpg",
      "productUrl": "http://example.com/products/skate-8",
      "createdDate": "2019-06-28T18:04:19-05:00",
      "updatedDate": "2019-06-28T18:04:19-05:00",
      "tstamp": "2019-06-28T18:04:19-05:00",
      "links": {
        "ecomOrder": "https://commissiontheatre.api-us1.com/api/3/ecomOrderProducts/1/ecomOrder"
      },
      "id": "1",
      "ecomOrder": "2"
    }
  ]
}
```

```bash
curl --request GET \
 --url https://youraccountname.api-us1.com/api/3/ecomOrderProducts \
 --header 'accept: application/json'
```

200 OK

```json
{
  "ecomOrderProducts": [
    {
      "orderid": "1",
      "connectionid": "1",
      "externalid": "1269385775",
      "sku": "SKATE-8",
      "name": "My Cool Skateboard",
      "description": "This skateboard is so cool, it rides itself",
      "price": "5000",
      "quantity": "1",
      "category": "skateboards",
      "imageUrl": "https://example.com/images/skateboard.jpg",
      "productUrl": "http://example.com/products/skate-8",
      "createdDate": "2019-06-28T18:04:19-05:00",
      "updatedDate": "2019-06-28T18:04:19-05:00",
      "tstamp": "2019-06-28T18:04:19-05:00",
      "links": {
        "ecomOrder": "https://commissiontheatre.api-us1.com/api/3/ecomOrderProducts/1/ecomOrder"
      },
      "id": "1",
      "ecomOrder": "2"
    }
  ]
}
```

400

```json
{}
```

## List EcomOrderProducts for a Specific EcomOrder

GET https://{youraccountname}.api-us1.com/api/3/ecomOrders/{id}/orderProducts

```json
{
  "ecomOrderProducts": [
    {
      "orderid": "1",
      "connectionid": "1",
      "externalid": "1269385775",
      "sku": "SKATE-8",
      "name": "My Cool Skateboard",
      "description": "This skateboard is so cool, it rides itself",
      "price": "5000",
      "quantity": "1",
      "category": "skateboards",
      "imageUrl": "https://example.com/images/skateboard.jpg",
      "productUrl": "http://example.com/products/skate-8",
      "createdDate": "2019-06-28T18:04:19-05:00",
      "updatedDate": "2019-06-28T18:04:19-05:00",
      "tstamp": "2019-06-28T18:04:19-05:00",
      "links": {
        "ecomOrder": "https://commissiontheatre.api-us1.com/api/3/ecomOrderProducts/1/ecomOrder"
      },
      "id": "1",
      "ecomOrder": "2"
    }
  ]
}
```

```bash
curl --request GET \
 --url https://youraccountname.api-us1.com/api/3/ecomOrders/id/orderProducts \
 --header 'accept: application/json'
```

200 OK

```json
{
  "ecomOrderProducts": [
    {
      "orderid": "1",
      "connectionid": "1",
      "externalid": "1269385775",
      "sku": "SKATE-8",
      "name": "My Cool Skateboard",
      "description": "This skateboard is so cool, it rides itself",
      "price": "5000",
      "quantity": "1",
      "category": "skateboards",
      "imageUrl": "https://example.com/images/skateboard.jpg",
      "productUrl": "http://example.com/products/skate-8",
      "createdDate": "2019-06-28T18:04:19-05:00",
      "updatedDate": "2019-06-28T18:04:19-05:00",
      "tstamp": "2019-06-28T18:04:19-05:00",
      "links": {
        "ecomOrder": "https://commissiontheatre.api-us1.com/api/3/ecomOrderProducts/1/ecomOrder"
      },
      "id": "1",
      "ecomOrder": "2"
    }
  ]
}
```

400

```json
{}
```

Retrieve an EcomOrderProduct
GET https://{youraccountname}.api-us1.com/api/3/ecomOrderProducts/{id}

```json
{
  "ecomOrderProduct": {
    "orderid": "1",
    "connectionid": "1",
    "externalid": "1269385775",
    "sku": "SKATE-8",
    "name": "My Cool Skateboard",
    "description": "This skateboard is so cool, it rides itself",
    "price": "5000",
    "quantity": "1",
    "category": "skateboards",
    "imageUrl": "https://example.com/images/skateboard.jpg",
    "productUrl": "http://example.com/products/skate-8",
    "createdDate": "2019-06-28T18:04:19-05:00",
    "updatedDate": "2019-06-28T18:04:19-05:00",
    "tstamp": "2019-06-28T18:04:19-05:00",
    "links": {
      "ecomOrder": "https://commissiontheatre.api-us1.com/api/3/ecomOrderProducts/1/ecomOrder"
    },
    "id": "1",
    "ecomOrder": "2"
  }
}
```

```bash
curl --request GET \
 --url https://youraccountname.api-us1.com/api/3/ecomOrderProducts/id \
 --header 'accept: application/json'
```

200 OK

```json
{
  "ecomOrderProduct": {
    "orderid": "1",
    "connectionid": "1",
    "externalid": "1269385775",
    "sku": "SKATE-8",
    "name": "My Cool Skateboard",
    "description": "This skateboard is so cool, it rides itself",
    "price": "5000",
    "quantity": "1",
    "category": "skateboards",
    "imageUrl": "https://example.com/images/skateboard.jpg",
    "productUrl": "http://example.com/products/skate-8",
    "createdDate": "2019-06-28T18:04:19-05:00",
    "updatedDate": "2019-06-28T18:04:19-05:00",
    "tstamp": "2019-06-28T18:04:19-05:00",
    "links": {
      "ecomOrder": "https://commissiontheatre.api-us1.com/api/3/ecomOrderProducts/1/ecomOrder"
    },
    "id": "1",
    "ecomOrder": "2"
  }
}
```

400

```json
{}
```

404

```json
{
  "message": "No Result found for EcomOrderProduct with id {the id requested}"
}
```
