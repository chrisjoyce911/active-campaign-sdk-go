# mocks/deals

Lightweight function-field mock for `deals.DealsService`.

Usage:

```go
import (
  dealsmock "github.com/chrisjoyce911/active-campaign-sdk-go/mocks/deals"
  "github.com/chrisjoyce911/active-campaign-sdk-go/services/deals"
)

svc := &dealsmock.Service{
  ListDealsFunc: func(ctx context.Context, opts map[string]string) (*deals.ListDealsResponse, *client.APIResponse, error) {
    return &deals.ListDealsResponse{Deals: []deals.Deal{{ID: "1"}}, Meta: &deals.DealsListMeta{Total: 1}}, nil, nil
  },
}
```

Set only the functions you need; others return harmless defaults.
