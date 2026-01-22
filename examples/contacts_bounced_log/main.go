package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contacts"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Overload()

	url := os.Getenv("ACTIVE_URL")
	token := os.Getenv("ACTIVE_TOKEN")
	contactID := os.Getenv("CONTACT_ID")
	if contactID == "" {
		contactID = "1" // default
	}

	if url == "" || token == "" {
		fmt.Fprintln(os.Stderr, "ACTIVE_URL and ACTIVE_TOKEN must be set")
		os.Exit(1)
	}

	core, err := client.NewCoreClient(url, token)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create core client: %v\n", err)
		os.Exit(1)
	}

	svc := contacts.NewRealService(core)

	ctx := context.Background()
	resp, apiResp, err := svc.GetContactBounceLogs(ctx, contactID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error getting bounce logs: %v\n", err)
		if apiResp != nil {
			fmt.Fprintf(os.Stderr, "status=%d body=%s\n", apiResp.StatusCode, string(apiResp.Body))
		}
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(out))
}
