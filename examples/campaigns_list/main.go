package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/campaigns"
	"github.com/joho/godotenv"
)

// Run prints campaign statuses using the provided campaigns service and writer.
func Run(ctx context.Context, svc campaigns.CampaignsService, out io.Writer) error {
	list, apiResp, err := svc.ListCampaigns(ctx, nil)
	if err != nil {
		return fmt.Errorf("list campaigns: %w (api resp: %+v)", err, apiResp)
	}

	for _, c := range list.Campaigns {
		i, perr := c.StatusInt()
		if perr != nil {
			fmt.Fprintf(out, "campaign %s (%s): status parse error: %v\n", c.ID, c.Name, perr)
			continue
		}
		st := c.StatusEnum()
		fmt.Fprintf(out, "campaign %s (%s): status=%d (%s)\n", c.ID, c.Name, i, st.String())
	}
	return nil
}

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
	if err := Run(ctx, svc, os.Stdout); err != nil {
		log.Fatalf("%v", err)
	}
}
