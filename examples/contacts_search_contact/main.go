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

	fs := flag.NewFlagSet("contacts_search_contact", flag.ExitOnError)
	url := fs.String("url", envURL, "ActiveCampaign base URL")
	token := fs.String("token", envToken, "ActiveCampaign API token")
	email := fs.String("email", envEmail, "Contact email to search for")
	fs.Parse(os.Args[1:])

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

	out, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(out))
}
