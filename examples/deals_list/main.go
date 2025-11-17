package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/deals"
	"github.com/joho/godotenv"
)

// Run lists deals for pipeline=2 and stage=7 and prints id, title, pipeline, and stage.
func Run(ctx context.Context, svc deals.DealsService, out io.Writer) error {
	opts := map[string]string{
		"filters[group]": "2",
		"filters[stage]": "7",
		"orders[title]":  "ASC",
	}
	list, apiResp, err := deals.ListDealsAll(ctx, svc, opts)
	if err != nil {
		return fmt.Errorf("list deals (all pages): %w (api resp: %+v)", err, apiResp)
	}
	for _, d := range list {
		fmt.Fprintf(out, "deal %s: %s (pipeline %s, stage %s)\n", d.ID, d.Title, d.Group, d.Stage)
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
	svc := deals.NewRealService(core)

	ctx := context.Background()
	if err := Run(ctx, svc, os.Stdout); err != nil {
		log.Fatalf("%v", err)
	}
}
