# services/deals

Typed client for ActiveCampaign Deals (create, get, list, update, delete), deal notes,
deal activities, and deal stages.

## Overview

This package wraps common Deals endpoints:

- Create, update, delete deals and notes
- Get a single deal, deal activities
- List all deals (with rich filtering, ordering, and pagination)
- List all deal stages (with filtering/ordering)

It also provides a helper to fetch all pages of deals via limit/offset.

## Permissions

- Deal permission: user must be allowed to manage deals
- Pipeline-specific permission: user must be allowed to manage the target
  pipeline(s). When missing, ListDeals returns limited fields for those deals
  (id, title, isDisabled=1). ListDealStages does not return stages from
  pipelines the user cannot manage.

## List deals — GET /api/3/deals

Method: `ListDeals(ctx context.Context, opts map[string]string) (*ListDealsResponse, *client.APIResponse, error)`

Filtering (pass exact keys in `opts`):

- filters[search] — search text used with filters[search_field]
- filters[search_field] — field to search
- filters[title] — deal title
- filters[stage] — stage ID (int)
- filters[group] — pipeline ID (int)
- filters[status] — deal status (int)
- filters[owner] — owner ID (int)
- filters[nextdate_range] — next task due date range
- filters[tag] — tag names on deal’s primary contact
- filters[tasktype] — deals having tasks of given type
- filters[created_before] / filters[created_after] — YYYY-MM-DD
- filters[updated_before] / filters[updated_after] — YYYY-MM-DD
- filters[organization] — primary contact’s organization ID (int)
- filters[minimum_value] / filters[maximum_value] — USD dollar portion (int)
- filters[score_greater_than] / filters[score_less_than] / filters[score] — "<score_id>:<score_value>"

Ordering (defaults to ASC):

- orders[title], orders[value], orders[cdate], orders[contact_name],
  orders[contact_orgname], orders[next-action]

Pagination:

- limit — max records per page (string)
- offset — starting offset (string)

Example:

```go
opts := map[string]string{
	"filters[group]": "2",   // pipeline id
	"filters[stage]": "7",   // stage id
	"orders[title]":  "ASC",
}
res, apiResp, err := svc.ListDeals(ctx, opts)
```

### Fetch all pages

Helper: `ListDealsAll(ctx, svc, opts) ([]Deal, *client.APIResponse, error)`

- Paginates using limit/offset until all pages are fetched
- Honors caller-provided `limit`/`offset` (defaults to limit=100, offset=0)
- Stops using `meta.total` when available or on a short page

```go
all, apiResp, err := deals.ListDealsAll(ctx, svc, map[string]string{
	"filters[group]": "2",
	"filters[stage]": "7",
	"orders[title]":  "ASC",
})
```

See runnable example: `../../examples/deals_list`.

## List deal stages — GET /api/3/dealStages

Method: `ListDealStages(ctx context.Context, opts map[string]string) (*ListDealStagesResponse, *client.APIResponse, error)`

Filtering/ordering:

- filters[title] — partial match by stage title
- filters[d_groupid] — pipeline ID (string)
- orders[title] — ASC|DESC

See runnable example: `../../examples/deals_list_all_stages`.

## Types

- `Deal` — deal record; some numeric-looking fields are strings per API
- `DealStage` — stage record (id, title, group, card regions, etc.)
- `ListDealsResponse` — `{ deals: [...], meta: { total, currencies... } }`
- `ListDealStagesResponse` — `{ dealStages: [...], meta: { total } }`

## Notes

- The field `isDisabled` on deals may be returned as 0/1 or boolean; the SDK
  provides a `Boolish` helper to parse it.
- `meta.total` can be a number or string; the SDK parses it permissively where
  needed.
