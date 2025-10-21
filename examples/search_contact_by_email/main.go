package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contacts"
)

func main() {
	activeURL := os.Getenv("ACTIVE_URL")
	token := os.Getenv("ACTIVE_TOKEN")
	email := os.Getenv("CONTACT_EMAIL")

	if activeURL == "" || token == "" || email == "" {
		fmt.Fprintln(os.Stderr, "Please set ACTIVE_URL, ACTIVE_TOKEN and CONTACT_EMAIL environment variables")
		os.Exit(2)
	}

	core, err := client.NewCoreClient(activeURL, token)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create core client: %v\n", err)
		os.Exit(1)
	}
	contactsSvc := contacts.NewRealService(core)

	ctx := context.Background()
	resp, apiResp, err := contactsSvc.SearchByEmail(ctx, email)
	if err != nil {
		// If API returned an error response, print some metadata to help debugging
		fmt.Fprintf(os.Stderr, "error searching contact by email: %v\n", err)
		if apiResp != nil {
			fmt.Fprintf(os.Stderr, "status: %d, body: %s\n", apiResp.StatusCode, string(apiResp.Body))
		}
		os.Exit(1)
	}

	// Pretty-print the response
	out, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(out))
}
