package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/deals"
	"github.com/joho/godotenv"
)

var exitFn = os.Exit

// Run lists all deal stages in pipeline 2 and prints their id and title.
func Run(ctx context.Context, svc deals.DealsService, out io.Writer) error {
	opts := map[string]string{
		"filters[d_groupid]": "2",
		"filters[title]":     "To Contact",
		"orders[title]":      "ASC",
	}
	res, apiResp, err := svc.ListDealStages(ctx, opts)
	if err != nil {
		return fmt.Errorf("list deal stages: %w (api resp: %+v)", err, apiResp)
	}
	for _, st := range res.DealStages {
		fmt.Fprintf(out, "stage %s: %s (pipeline %s)\n", st.ID, st.Title, st.Group)
	}
	return nil
}

func main() {
	if os.Getenv("ACTIVE_URL") == "" {
		_ = godotenv.Load()
	}

	base := os.Getenv("ACTIVE_URL")
	token := os.Getenv("ACTIVE_TOKEN")
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
	svc := deals.NewRealService(core)

	ctx := context.Background()
	if err := Run(ctx, svc, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		exitFn(1)
		return
	}
}
