package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/campaigns"
	"github.com/joho/godotenv"
)

// Example: list campaigns and print their status
// Usage:
//
//	export AC_BASE_URL="https://youraccount.api-us1.com"
//	export AC_TOKEN="your_token"
//	go run ./examples/list_campaigns
func main() {

	_ = godotenv.Load()

	base := os.Getenv("ACTIVE_URL")
	token := os.Getenv("ACTIVE_TOKEN")
	if base == "" || token == "" {
		log.Fatalf("ACTIVE_URL and ACTIVE_TOKEN must be set")
	}

	core, err := client.NewCoreClient(base, token)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create core client: %v\n", err)
		os.Exit(1)
	}
	svc := campaigns.NewRealService(core)

	ctx := context.Background()

	list, apiResp, err := svc.ListCampaigns(ctx, nil)
	if err != nil {
		log.Fatalf("list campaigns: %v (api resp: %+v)", err, apiResp)
	}

	for _, c := range list.Campaigns {
		i, perr := c.StatusInt()
		if perr != nil {
			fmt.Printf("campaign %s (%s): status parse error: %v\n", c.ID, c.Name, perr)
			continue
		}
		st := c.StatusEnum()
		fmt.Printf("campaign %s (%s): status=%d (%s)\n", c.ID, c.Name, i, st.String())
	}
}
