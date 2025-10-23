# services/campaigns

Campaigns are broadcast emails sent to lists of contacts. This package
implements the core campaign endpoints (create, get, list, edit, delete,
send, copy) and includes unit tests and examples that exercise request wiring
and response decoding.

See the ActiveCampaign API docs for Campaigns: https://developers.activecampaign.com/reference/campaign

Contents

- Overview
- Status values
- Usage examples (list / status)
- Create, Edit and Duplicate examples
- Legacy: template share endpoint (V2)

Examples

- A runnable example that lists campaigns is available at `examples/campaigns_list/main.go`.

Quick example: run the sample

```bash
# set your account base URL and API token
export AC_BASE_URL="https://youraccount.api-us1.com"
export AC_TOKEN="your_token"

# run the example
go run ./examples/campaigns_list
```

Overview

This package follows the project's typed service patterns: request payloads
are represented by Go structs (e.g. `CreateCampaignRequest`,
`EditCampaignRequest`) and many endpoints return typed response structs
(e.g. `Campaign`, `ListCampaignsResponse`, `DuplicateCampaignResponse`).

Status values

For sending status use the following values:

| Status Value | Meaning   |
| -----------: | :-------- |
|            0 | Draft     |
|            1 | Scheduled |
|            2 | Sending   |
|            3 | Paused    |
|            4 | Stopped   |
|            5 | Completed |

Usage: listing campaigns and checking status

This short example shows how to call `ListCampaigns`, iterate the returned
`Campaigns` slice, and use the `StatusInt()` / `StatusEnum()` helpers.

```go
package main

import (
  "context"
  "fmt"
  "log"

  "github.com/chrisjoyce911/active-campaign-sdk-go/services/campaigns"
)

func main() {
  ctx := context.Background()
  // svc should be your campaigns service instance created via client.NewRealService(...)
  var svc campaigns.CampaignsService

  list, apiResp, err := svc.ListCampaigns(ctx, nil)
  if err != nil {
    log.Fatalf("list campaigns: %v", err)
  }
  _ = apiResp

  for _, c := range list.Campaigns {
    i, err := c.StatusInt()
    if err != nil {
      fmt.Printf("campaign %s: status parse error: %v\n", c.ID, err)
      continue
    }
    st := c.StatusEnum()
    fmt.Printf("campaign %s: status=%d (%s)\n", c.ID, i, st.String())
  }
}
```

Notes

- `StatusInt()` returns an integer parsed from the API's string-encoded
  status. It now returns an error when the underlying `Status` string is
  empty; callers should handle that case explicitly.
- `StatusEnum()` returns a typed `CampaignStatus` and defaults to Draft on
  parse error; use it when you prefer a safe fallback.

Create, Edit and Duplicate campaigns

CreateCampaign accepts a typed `CreateCampaignRequest`:

```go
req := &campaigns.CreateCampaignRequest{Name: "Campaign Name", Type: "single"}
created, apiResp, err := svc.CreateCampaign(ctx, req)
if err != nil {
  // handle error
}
_ = created // *campaigns.Campaign
```

EditCampaign accepts a typed `EditCampaignRequest`. Use the builder helpers for concise construction:

```go
req := campaigns.NewEditCampaignRequest("My Campaign").WithAddressID(2).WithListIDs(10,20).WithLinkTrackingEnabled(true)
updated, apiResp, err := svc.EditCampaign(ctx, "29", req)
if err != nil {
  // handle error
}
_ = updated
```

Duplicate a campaign (copy):

```go
dup, apiResp, err := svc.DuplicateCampaign(ctx, "2")
if err != nil {
  // handle error (client.APIError on non-2xx)
}
if dup != nil {
  fmt.Printf("new campaign id: %d", dup.NewCampaignID)
}
```

Legacy: Create Shareable Campaign Template Link (V2)

Create a shareable link to a campaign template. This endpoint is a V2-style
legacy endpoint; the SDK does not currently provide a typed helper for it but
the HTTP contract is shown below for reference.

GET https://{youraccountname}.api-us1.com/api/2/template/share

Query Parameters

- `api_key` (string) — ActiveCampaign Account API Key
- `id` (int) — The id of the campaign template to be shared

Success response example

```json
{
  "url": "http://tplshare.com/gXKi4s2$",
  "success": 1,
  "message": "Template shared successfully",
  "result_code": 1,
  "result_message": "Template shared successfully",
  "result_output": "json"
}
```

Example curl

```bash
curl --request GET \
	 --url "https://youraccountname.api-us1.com/api/2/template/share?api_key=YOUR_KEY&id=123" \
	 --header 'accept: application/json'
```

## Usage: Create, Edit and Duplicate campaigns

CreateCampaign now accepts a typed `CreateCampaignRequest`:

```go
req := &campaigns.CreateCampaignRequest{Name: "Campaign Name", Type: "single"}
created, apiResp, err := svc.CreateCampaign(ctx, req)
if err != nil {
  // handle error
}
_ = created // *campaigns.Campaign
```

EditCampaign accepts a typed `EditCampaignRequest`. Use the builder helpers for concise construction:

```go
req := campaigns.NewEditCampaignRequest("My Campaign").WithAddressID(2).WithListIDs(10,20).WithLinkTrackingEnabled(true)
updated, apiResp, err := svc.EditCampaign(ctx, "29", req)
if err != nil {
  // handle error
}
_ = updated
```

Duplicate a campaign (copy):

```go
dup, apiResp, err := svc.DuplicateCampaign(ctx, "2")
if err != nil {
  // handle error (client.APIError on non-2xx)
}
if dup != nil {
  fmt.Printf("new campaign id: %d", dup.NewCampaignID)
}
```
