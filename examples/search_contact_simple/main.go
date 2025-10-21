package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contacts"
)

func main() {
	var (
		envURL   = os.Getenv("ACTIVE_URL")
		envToken = os.Getenv("ACTIVE_TOKEN")
		envEmail = os.Getenv("CONTACT_EMAIL")
	)

	// Allow flags to override env vars for convenience in CI/manual runs
	url := flag.String("url", envURL, "ActiveCampaign base URL (e.g. https://youraccount.api-us1.com)")
	token := flag.String("token", envToken, "ActiveCampaign API token")
	email := flag.String("email", envEmail, "Contact email to search for")
	flag.Parse()

	if *url == "" || *token == "" || *email == "" {
		fmt.Fprintln(os.Stderr, "Please provide ActiveCampaign credentials and CONTACT_EMAIL via flags or environment variables")
		flag.Usage()
		os.Exit(2)
	}

	core, err := client.NewCoreClient(*url, *token)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create core client: %v\n", err)
		os.Exit(1)
	}
	svc := contacts.NewRealService(core)

	ctx := context.Background()
	resp, apiResp, err := svc.SearchByEmail(ctx, *email)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error searching contact by email: %v\n", err)
		if apiResp != nil {
			fmt.Fprintf(os.Stderr, "status=%d body=%s\n", apiResp.StatusCode, string(apiResp.Body))
		}
		os.Exit(1)
	}

	// Print a compact JSON representation of the response
	out, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(out))
}
