package deals

import (
    "context"
    "strconv"

    "github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// ListDealsAll retrieves all deals matching the provided filters by paginating
// through GET /api/3/deals using limit/offset under the hood.
//
// Behavior:
// - Starts from offset=0 (or the caller-provided offset) and continues
//   requesting pages until all items are retrieved.
// - Uses the provided limit if present; otherwise defaults to 100.
// - Stops when the accumulated count reaches meta.total (when available) or
//   when a page returns fewer items than the limit.
// - Returns the concatenated slice of Deal and the last APIResponse.
//
// Parameters:
// - ctx  : context.Context
// - svc  : DealsService used to issue requests
// - opts : base query parameters (e.g., filters[group], filters[stage]).
//          limit/offset will be managed by this helper but any caller-provided
//          values are honored as starting values.
//
// Returns: ([]Deal, *client.APIResponse, error)
func ListDealsAll(ctx context.Context, svc DealsService, opts map[string]string) ([]Deal, *client.APIResponse, error) {
    // Copy opts so we don't mutate the caller's map.
    base := map[string]string{}
    if opts != nil {
        for k, v := range opts {
            base[k] = v
        }
    }

    // Determine limit (default 100) and starting offset (default 0).
    limit := 100
    if s, ok := base["limit"]; ok && s != "" {
        if n, err := strconv.Atoi(s); err == nil && n > 0 {
            limit = n
        }
    }
    offset := 0
    if s, ok := base["offset"]; ok && s != "" {
        if n, err := strconv.Atoi(s); err == nil && n >= 0 {
            offset = n
        }
    }

    var (
        all   []Deal
        last  *client.APIResponse
    )

    for {
        pageOpts := map[string]string{}
        for k, v := range base {
            pageOpts[k] = v
        }
        pageOpts["limit"] = strconv.Itoa(limit)
        pageOpts["offset"] = strconv.Itoa(offset)

        resp, apiResp, err := svc.ListDeals(ctx, pageOpts)
        if err != nil {
            return nil, apiResp, err
        }
        last = apiResp

        // Append page results
        if resp != nil {
            all = append(all, resp.Deals...)
        }

        // Stop conditions: reached reported total (when available) or
        // fewer results than the limit indicating the last page.
        total := -1
        if resp != nil && resp.Meta != nil {
            total = resp.Meta.Total
        }

        if total >= 0 && len(all) >= total {
            break
        }
        if resp == nil || len(resp.Deals) < limit {
            break
        }
        offset += limit
    }

    return all, last, nil
}
