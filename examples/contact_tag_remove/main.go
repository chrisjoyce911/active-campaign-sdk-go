package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contacts"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	fs := flag.NewFlagSet("contact_tag_remove", flag.ExitOnError)
	url := fs.String("url", os.Getenv("ACTIVE_URL"), "ActiveCampaign base URL")
	token := fs.String("token", os.Getenv("ACTIVE_TOKEN"), "ActiveCampaign API token")
	contactID := fs.String("contact-id", os.Getenv("ACTIVE_CONTACTID"), "Contact ID")
	tagID := fs.String("tag-id", os.Getenv("ACTIVE_CONTACT_TAG"), "Tag ID to remove")
	fs.Parse(os.Args[1:])

	if *url == "" || *token == "" || *contactID == "" || *tagID == "" {
		fmt.Fprintln(os.Stderr, "Please provide ActiveCampaign credentials, contact-id, and tag-id via flags or environment variables")
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
	apiResp, err := svc.TagRemove(ctx, *contactID, *tagID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error removing tag: %v\n", err)
		if apiResp != nil {
			fmt.Fprintf(os.Stderr, "status=%d body=%s\n", apiResp.StatusCode, string(apiResp.Body))
		}
		os.Exit(1)
	}

	fmt.Printf("Tag %s removed successfully from contact %s\n", *tagID, *contactID)
}
