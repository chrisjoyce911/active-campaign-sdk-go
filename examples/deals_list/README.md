# Deals: List Deals Example

This example lists deals filtered by pipeline and stage using the Deals service and paginates to retrieve all results.

- Endpoint: GET /api/3/deals
- Example filters: pipeline id 2 and stage id 7

## Prerequisites

- Set environment variables:
  - `ACTIVE_URL` (e.g. https://youraccount.api-us1.com)
  - `ACTIVE_TOKEN`

## Run (fetch all pages)

```
go run ./examples/deals_list
```

Expected output (truncated, similar):

```
deal 46: Able Hyena (pipeline 2, stage 7)
deal 1: Test Deal (pipeline 2, stage 7)
```

## How it works

The example calls:

```
deals.ListDealsAll(ctx, svc, map[string]string{
  "filters[group]": "2",   // pipeline id
  "filters[stage]": "7",   // stage id within the pipeline
  "orders[title]":  "ASC", // optional ordering
  // Optional: "limit": "200", "offset": "0"
})
```

This helper automatically paginates using limit/offset until all records are fetched (based on meta.total or a short page). It returns a []deals.Deal. The example prints each deal's ID, title, pipeline (group), and stage.
