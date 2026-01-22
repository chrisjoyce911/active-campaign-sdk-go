package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/campaigns"
	"github.com/joho/godotenv"
)

// allow tests to intercept process exit
var exitFn = os.Exit

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
	err := godotenv.Load()
	if err != nil {
		fmt.Println("godotenv error:", err)
	}

	base := os.Getenv("ACTIVE_URL")
	token := os.Getenv("ACTIVE_TOKEN")
	fmt.Printf("Loaded token: %s\n", token) // Debug print
	if base == "" || token == "" {
		fmt.Fprintln(os.Stderr, "ACTIVE_URL and ACTIVE_TOKEN must be set")
		exitFn(1)
		return
	}

	core, err := client.NewCoreClient(base, token)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create core client: %v\n", err)
		exitFn(1)
		return
	}

	// Enable debug logging for outgoing requests
	core.SetDebug(true, os.Stderr)
	svc := campaigns.NewRealService(core)

	ctx := context.Background()
	if err := Run(ctx, svc, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		exitFn(1)
		return
	}
}
