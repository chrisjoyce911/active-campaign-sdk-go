//go:build examples

package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contacts"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("ACTIVE_URL") == "" {
		_ = godotenv.Load()
	}

	fs := flag.NewFlagSet("contact_add_to_list", flag.ExitOnError)
	apply := fs.Bool("apply", false, "Execute the subscription call; default prints the request that would be sent.")
	contactIDFlag := fs.String("contact-id", "", "Contact ID to subscribe (overrides CONTACT_ID env var).")
	listIDFlag := fs.String("list-id", "", "List ID to subscribe (overrides LIST_ID env var).")
	statusFlag := fs.String("status", "", "Subscription status code (defaults to LIST_STATUS env or 1 to subscribe).")
	args := filterTestFlags(os.Args[1:])
	_ = fs.Parse(args)

	cfg := runConfig{
		ContactID: firstNonEmpty(*contactIDFlag, os.Getenv("CONTACT_ID")),
		ListID:    firstNonEmpty(*listIDFlag, os.Getenv("LIST_ID")),
		Status:    firstNonEmpty(*statusFlag, os.Getenv("LIST_STATUS")),
	}

	baseURL := envOrFatal("ACTIVE_URL")
	token := envOrFatal("ACTIVE_TOKEN")

	coreClient, err := client.NewCoreClient(baseURL, token)
	if err != nil {
		log.Fatalf("failed to create core client: %v", err)
	}

	svc := contacts.NewRealService(coreClient)
	if err := run(context.Background(), svc, cfg, *apply, os.Stdout); err != nil {
		log.Fatalf("contact_add_to_list: %v", err)
	}
}
