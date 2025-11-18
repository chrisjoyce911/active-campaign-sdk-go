# mocks/campaigns

Function-field mock for `campaigns.CampaignsService`.

```go
import (
  campaignsmock "github.com/chrisjoyce911/active-campaign-sdk-go/mocks/campaigns"
  "github.com/chrisjoyce911/active-campaign-sdk-go/services/campaigns"
)

svc := &campaignsmock.Service{
  ListCampaignsFunc: func(ctx context.Context, _ interface{}) (*campaigns.ListCampaignsResponse, *client.APIResponse, error) {
    return &campaigns.ListCampaignsResponse{Campaigns: []campaigns.Campaign{{ID: "c1", Status: "1"}}}, nil, nil
  },
}
```
