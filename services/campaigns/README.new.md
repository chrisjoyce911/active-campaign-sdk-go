````markdown
# services/campaigns

Campaigns are broadcast emails sent out to a list of contacts.

See the ActiveCampaign API docs for Campaigns: https://developers.activecampaign.com/reference/campaign

This package implements the core campaign endpoints (create, get, list, update, delete, send) and includes unit tests and examples that exercise request wiring and responses.

Campaign status available values

For sending status use the following values:

| Status Value | Meaning   |
| -----------: | :-------- |
|            0 | Draft     |
|            1 | Scheduled |
|            2 | Sending   |
|            3 | Paused    |
|            4 | Stopped   |
|            5 | Completed |

## Create Shareable Campaign Template Link

Create a shareable link to a campaign template. This endpoint is a V2-style
legacy endpoint; the SDK does not currently provide a typed helper for it but
the HTTP contract is shown below for reference.

GET https://{youraccountname}.api-us1.com/api/2/template/share

Query Parameters

- api_key (string) — ActiveCampaign Account API Key
- id (int) — The id of the campaign template to be shared

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

## Usage: listing campaigns and checking status

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

Notes:

- `StatusInt()` returns an integer parsed from the API's string-encoded status.
- `StatusEnum()` returns a typed `CampaignStatus` and defaults to Draft on parse error.
````
