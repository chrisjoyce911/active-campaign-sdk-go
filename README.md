# active-campaign-sdk-go

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/benkrig/active-campaign-sdk-go?tab=overview)
[![Build Status](https://travis-ci.com/benkrig/active-campaign-sdk-go.svg?token=zD75aqrV8gE1Q1ghw6yU&branch=master)](https://travis-ci.com/benkrig/active-campaign-sdk-go)
[![codecov](https://codecov.io/gh/benkrig/active-campaign-sdk-go/branch/master/graph/badge.svg?token=PR8PBM0NGX)](https://codecov.io/gh/benkrig/active-campaign-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/benkrig/active-campaign-sdk-go)](https://goreportcard.com/report/github.com/benkrig/active-campaign-sdk-go)

**active-campaign-sdk-go** provides access to the [Active Campaign API V3](https://developers.activecampaign.com/reference) for Go. Currently, it's heavily under development.

## Usage

```go
package main

import "github.com/benkrig/active-campaign-sdk-go"
```

Construct a new client, then use the services available within the client to access the Active Campaign API.

```go
package main

import (
    ac "github.com/benkrig/active-campaign-sdk-go"
    "os"
)

func main() {
    baseURL := os.Getenv("YOUR_BASE_URL_KEY")
    token := os.Getenv("YOUR_TOKEN_KEY")

    a, err := ac.NewClient(
        &ac.ClientOpts{
            BaseUrl: baseURL,
            Token: token,
        },
    )
    if err != nil {
        panic(err)
    }

    c := ac.CreateContactRequest{
        &ac.Contact{
            Email: "test@email.com",
            FirstName: "testf",
            LastName: "testl",
            Phone: "1234567890",
        },
    }

    contact, _, err := a.Contacts.Create(&c)
    if err != nil {
        panic(err)
    }
}
```

## Code structure

The code structure of this package was inspired by [google/go-github](https://github.com/google/go-github) and [andygrunwald/go-jira](https://github.com/andygrunwald/go-jira).

Everything is based around the Client. The Client contains various services for resources found in the Active Campaign API, like Contacts, or Automations. Each service implements actions for its respective resource(s).

## Contribution

PR's are always welcome! The SDK is still being heavily developed and is missing many entities.

It doesn't matter if you are not able to write code.
Creating issues or holding talks and helping other people use the SDK is contribution as well!
A few examples:

- Correct typos in the README / documentation
- Reporting bugs
- Implement a new feature or endpoint

If you are new to pull requests, checkout [Collaborating on projects using issues and pull requests / Creating a pull request](https://help.github.com/articles/creating-a-pull-request/).

## Code coverage

This repository includes CI that runs the test suite and produces a coverage
report. The GitHub Actions workflow (`.github/workflows/ci.yml`) will run
`go test ./... -coverprofile=coverage.out` and upload `coverage.out` as a
workflow artifact. If you configure the secret `CODECOV_TOKEN` the job will
also upload coverage to Codecov.

Run coverage locally:

```bash
# run tests with coverage and open an HTML report
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
open coverage.html
```

If you'd like Codecov badges added to the README, set up the repository on
Codecov and provide `CODECOV_TOKEN` as a repository secret (private repos).

## License

This project is released under the terms of the [MIT license](http://en.wikipedia.org/wiki/MIT_License).

## Active Campaign API Reference (status)

Below is a checklist-style summary of what this repository implements today (v3 branch) and which areas are known/interface-only. This is a high-level status map based on the code in the `services/` directory (presence of per-endpoint files, `impl_real.go`, and package tests). If an area is marked "Implemented (partial)" it means core endpoints exist but additional endpoints or fields remain TODO.

Legend:

- ✅ Implemented — package has RealService implementations and tests for core endpoints
- ⚠️ Implemented (partial) — package has RealService but not every endpoint/field is converted to typed responses
- ⬜ Interface-known / Planned — listed in README but not yet migrated to typed RealService endpoints

_Status snapshot: 20 Oct 2025_

- Accounts: ✅ Implemented (create/get/list/update/delete, notes) — see `services/accounts` (has `impl_real.go` and tests)
- Contacts: ⚠️ Implemented (partial) — `GetContactTags` and several core endpoints have been migrated to typed RealService; other contact endpoints remain TODO in `services/contacts`
- Tags: ✅ Implemented — `services/tags` provides typed endpoints and tests (list/create/update/delete/add-to-contact)
- Contact Automations: ✅ Implemented — `services/contactautomation` contains typed endpoints and tests (get counts, get automation entry, add/remove contact, list)
- Deals: ✅ Implemented (core) — `services/deals` has RealService implementations and tests for main deal endpoints
- Campaigns: ✅ Implemented (core) — `services/campaigns` includes implementations and tests for listing, get, create, send, and update flows
- Users: ✅ Implemented (core) — `services/users` has implementations and tests for user CRUD flows
- Webhooks: ✅ Implemented (core) — `services/webhooks` has per-endpoint files and tests
- Groups: ✅ Implemented (core) — `services/groups` contains implementations and tests
- Lists: ✅ Implemented (core) — `services/lists` has implementations for list endpoints
- Messages: ✅ Implemented (core) — `services/messages` contains implementations
- Custom Objects: ⚠️ Implemented (partial) — `services/custom_objects` contains schema + record helpers; some areas remain incomplete
- Tracking: ✅ Implemented (core) — `services/tracking` contains site/event tracking helpers and tests
- E-Commerce: ✅ Implemented (core) — `services/ecommerce` has implementations for store, customers, orders, etc.

Notes & assumptions

- This checklist was generated by scanning `services/` for `impl_real.go` files and existing package tests as a signal of implemented endpoints. It does not guarantee full parity with every endpoint in the ActiveCampaign API documentation.
- "Partial" means at least one endpoint in the package was migrated to the v3 typed RealService shape; more endpoints may still return interface{} or be missing typed models.
- If you'd like, I can continue converting specific endpoints to typed request/response shapes (one package at a time) and update this README checklist automatically as I complete each migration.

If you'd like an exportable checklist (Markdown) filtered to a subset of services or a pull request ready to update the README on a different branch, tell me which services to prioritize and I'll update the file accordingly.

# Active Campaign API Reference

Accounts
https://developers.activecampaign.com/reference/create-an-account-new
GET Account
POST Create Account
GET Accounts list
DEL Account
DEL Accounts (bulk)
PUT Update Account
POST Create Account Note
PUT Update Account Note

Contacts
https://developers.activecampaign.com/reference/contact
GET Get Contacts
GET Get Contact by ID
GET Get Contact Data
GET Get Contact Tags
GET Get Contact by Email with Tags
GET Get Contact Bounce Logs
GET Get Contact Data Goals
GET Get Contact Lists
GET Get Contact Logs
GET Get Contact Deal List
GET Get Contact Deals
GET Get Contact Field Values
GET Get Contact Geo Ips List
GET Get Contact Geo Ip Address
GET Get Contact Notes
GET Get Contact Organization
GET Get Contact Account Contacts
GET Get Contact Plus Append
GET Get Contact Tracking Logs
POST Create Contact
POST Sync Contact Data
POST Add contact to list
DEL Delete Contact
POST Bulk Import
GET Bulk Import Status Info
GET Bulk Import Status List

    Automations
    https://developers.activecampaign.com/reference/list-all-automations
    GET Get Contact Automation Entry Counts
    GET Get an Automation a Contact is in
    POST Add a Contact to an Automation
    DEL Remove a Contact from an Automation
    GET Get Contact Automations

    Custom Fields and Values
    POST Add Custom Field
    POST Add Custom Field to Field Group
    GET Get Custom Field Field Group
    PUT Update Custom Field Field Group
    POST Delete Custom Field Field Group
    DEL Delete Custom Field
    PUT Update Custom Field
    GET List All Custom Fields
    POST Add Custom Field Options

        Field Values
        POST Update Custom Field Value For Contact
        GET List All Custom Field Values

        Field Relationships
        POST Add Custom Field
        POST Add Custom Field to Field Group
        GET Get Custom Field Field Group
        PUT Update Custom Field Field Group
        POST Delete Custom Field Field Group
        DEL Delete Custom Field
        PUT Update Custom Field
        GET List All Custom Fields
        POST Add Custom Field Options

Campaigns
https://developers.activecampaign.com/reference/campaign
GET Get Campaigns
GET Get Campaign by ID
GET Get Campaign User
GET Get Campaign Automations
GET Get Campaign Message
GET Get Campaign Messages
GET Get Campaign Links
GET Get Campaign Aggregate Revenues
GET Get Campaign Automation Segments
GET Get Campaign Automation Campaign Lists
PUT Edit Campaign
POST Duplicate Campaign
POST Create Campaign
GET Get Links Associated to Campaign
GET Create Shareable Campaign Template Link

    Variables
    GET Get Variable(s)
    GET Get Variable
    POST Create Variable
    PUT Edit Variable
    DEL Delete Variable
    DEL Bulk Delete Variables
    GET Get Campaigns
    GET Get Campaign by ID
    GET Get Campaign User
    GET Get Campaign Automations
    GET Get Campaign Message
    GET Get Campaign Messages
    GET Get Campaign Links
    GET Get Campaign Aggregate Revenues
    GET Get Campaign Automation Segments
    GET Get Campaign Automation Campaign Lists
    PUT Edit Campaign
    POST Duplicate Campaign
    POST Create Campaign
    GET Get Links Associated to Campaign
    GET Create Shareable Campaign Template Link

Messages
GET Get List of Messages
GET Retrieve a Message
POST Create a Message
DEL Delete a Message
PUT Update a Message

Tags
https://developers.activecampaign.com/reference/create-a-new-tag
GET List all tags
POST Create a tag
PUT Update a tag
DEL Delete a tag
POST Add a tag to a contact

Lists
GET Get Lists
POST Create List

Custom Objects

    Schemas
    https://developers.activecampaign.com/reference/custom-object-schemas
    GET List all schemas
    GET Listing records for a Schema
    PUT Update Schema
    POST Create Schema
    GET Get Schema by ID
    DEL Delete Schema

        Parent and Child Schemas
        https://developers.activecampaign.com/reference/parent-and-child-schemas
        POST Create a public schema
        POST Create a child schema

    Managing Records
    https://developers.activecampaign.com/reference/custom-object-records
    GET Get a list of records
    POST Create or Update a record
    GET Get record by id
    DEL Delete record by id
    GET Get record by external id
    DEL Delete record by external id

Groups
GET Get Groups
GET Get Group Limits
POST Create Group
GET Get Group By ID
GET Get Users By Group

Users
POST Create User
GET Get Users
GET Get User
GET Get User By Email
GET Get User By Username
GET Get Logged-In User
PUT Update a User
DEL Delete a User

Deals
https://developers.activecampaign.com/reference/deal
POST Create a deal
POST Create a deal note
GET Retrieve a deal
GET Retrieve A Deal Activities
GET Retrieve all deals
DEL Delete a deal
PUT Update a deal
PUT Update a deal note
PATCH Bulk update deal owners

    Deal Roles
    https://developers.activecampaign.com/reference/list-all-deal-roles
    POST Create a deal role

Site & Event Tracking

    Event Tracking
    POST Create a new Event (Name Only)
    PUT Enable/Disable Event Tracking
    POST Track Event
    GET List All Events (Name Only)
    GET Get Event Tracking Status
    DEL Delete Event
    GET Get Contact By Event ID
    GET Get Contact Events and Activities
    GET Get Contact's Tracking Logs/Events

    Site Tracking
    POST Add Domain To WhiteList
    DEL Remove Domain From WhiteList
    GET Get Site Tracking Code
    GET Get Site Tracking Status
    GET List All WhiteListed Domains
    PUT Enable/Disable Site Tracking

E-Commerce (REST)

    Connections
    POST Create a Connection
    PUT Update a Connection
    DEL Delete a Connection
    GET Retrieve a Connection
    GET Retrieve all Connections

Customers
POST Create a Customer
PUT Update a Customer
GET Retrieve a Customer
GET Retrieve all Customers
DEL Delete a Customer

    Orders
    POST Create an Order
    PUT Update an Order
    POST Create Abandoned Cart
    GET Retrieve an Order
    GET Retrieve all Orders
    DEL Delete an Order

    Products
    GET Retrieve Order Products
    GET Retrieve all Products for a Specific Order
    GET Retrieve an Order Product

Webhooks
https://developers.activecampaign.com/reference/webhooks
POST Create a Webhook
PUT Update a Webhook
GET Retrieve a Webhook
GET Retrieve All Webhooks
GET Retrieve All Webhook Events
DEL Delete a Webhook

## Debugging outgoing requests

The SDK `CoreClient` includes a simple, opt-in debug facility to dump outgoing
request bodies for easier troubleshooting. This is helpful when developing
against the ActiveCampaign API and wanting to inspect the exact JSON the
client sends.

Usage examples:

- Print debug output to stdout:

```go
cc.SetDebug(true, os.Stdout)
```

- Capture debug output in tests (inspect later):

```go
var buf bytes.Buffer
cc.SetDebug(true, &buf)
// perform calls
// assert strings.Contains(buf.String(), "DEBUG OUTGOING")
```

If you pass a nil writer the SDK falls back to the standard logger
(`log.Printf`). Debug output is gated by a boolean flag so it must be
explicitly enabled. In the current SDK behavior, when Debug is enabled the
client will emit debug output for all outgoing requests by default. If you
prefer to only log a subset of requests, provide a filter using
`SetDebugFilter` (see example below).

The debug header looks like:

```
DEBUG OUTGOING <METHOD> <URL> body:
{...json body...}
```

If you prefer different formatting or structured logs (JSON), you can pass a
custom writer that formats the bytes any way you like.

Example: restrict debug to only POSTs to custom objects endpoints.

```go
cc.SetDebug(true, os.Stdout)
cc.SetDebugFilter(func(method, path string) bool {
    if method != "POST" {
        return false
    }
    return strings.Contains(path, "/customObjects/")
})
```
