# Deals: List All Stages Example

This example lists all deal stages for a specific pipeline using the Deals service.

- Endpoint: GET /api/3/dealStages
- Example filters: pipeline id 2 via `filters[d_groupid]=2`

## Prerequisites

- Set environment variables:
  - `ACTIVE_URL` (e.g. https://youraccount.api-us1.com)
  - `ACTIVE_TOKEN`

## Run

```
go run ./examples/deals_list_all_stages
```

Expected output (similar):

```
stage 15: Initial Contact (pipeline 4)
stage 16: Qualifications - Low (pipeline 4)
```

## How it works

The example calls:

```
svc.ListDealStages(ctx, map[string]string{
  "filters[d_groupid]": "2",
  "orders[title]":      "ASC",
})
```

and prints each stage's ID, title, and pipeline (group).
