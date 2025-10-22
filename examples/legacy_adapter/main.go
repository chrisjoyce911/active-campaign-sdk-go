package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	ac "github.com/chrisjoyce911/active-campaign-sdk-go"
	"github.com/chrisjoyce911/active-campaign-sdk-go/legacy"
	"github.com/joho/godotenv"
)

func env(k string) string { return strings.TrimSpace(os.Getenv(k)) }

func main() {
	_ = godotenv.Load()
	flag.Parse()

	// Legacy client creation (examples historically used this API)
	client, err := ac.NewClient(&ac.ClientOpts{BaseUrl: env("ACTIVE_URL"), Token: env("ACTIVE_TOKEN")})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create legacy client: %v\n", err)
		os.Exit(1)
	}

	// This stubbed client doesn't perform real requests in this migration, but
	// illustrates the old construction pattern used by older examples.
	_ = client

	// Old legacy client call pattern (examples used this):
	email := env("ACTIVE_EMAIL")
	legacyContact, legacyResp, legacyErr := client.Contacts.SearchEmail(email)
	if legacyErr != nil {
		fmt.Fprintf(os.Stderr, "legacy client search error: %v (resp=%v)\n\n", legacyErr, legacyResp)
	} else {
		fmt.Printf("client.Contacts.SearchEmail(email) legacy client search result:\n%#v (resp=%#v)\n\n", legacyContact, legacyResp)
	}

	// Prefer using the adapter which uses the typed services under the hood.
	// The legacy adapter reads ACTIVE_URL and ACTIVE_TOKEN from the environment.
	if env("ACTIVE_URL") == "" || env("ACTIVE_TOKEN") == "" {
		fmt.Fprintln(os.Stderr, "set ACTIVE_URL and ACTIVE_TOKEN in your environment to call the API")
		os.Exit(2)
	}

	// Example: search for a contact by email using the legacy adapter.
	resp, apiResp, err := legacy.SearchContacts(context.Background(), email)
	if err != nil {
		fmt.Fprintf(os.Stderr, "search error: %v (apiResp=%v)\n\n", err, apiResp)
		os.Exit(1)
	}
	fmt.Printf("legacy.SearchContacts search response:\n%#v\n\n", resp)

	// Example: get automation counts via the adapter
	counts, apiResp, err := legacy.GetAutomationCounts(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "automation counts error: %v (apiResp=%v)\n\n", err, apiResp)
		os.Exit(1)
	}
	fmt.Printf("automation counts: %#v\n\n", counts)
}
