//go:build examples

package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contacts"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()
	baseURL := os.Getenv("ACTIVE_URL")
	token := os.Getenv("ACTIVE_TOKEN")

	c, err := client.NewCoreClient(baseURL, token)
	if err != nil {
		log.Fatalf("failed to create core client: %v", err)
	}

	svc := contacts.NewRealService(c)

	if baseURL == "" {
		log.Printf("ACTIVE_URL not set; running example in placeholder mode")
	} else {
		log.Printf("ACTIVE_URL set to %s", baseURL)
	}
	if token == "" {
		log.Printf("ACTIVE_TOKEN not set")
	} else {
		log.Printf("ACTIVE_TOKEN is set (redacted)")
	}

	// Fetch and print contact tags for an example contact id
	contactID := "287199"
	out, apiResp, err := svc.TagsGet(context.Background(), contactID)
	if err != nil {
		// Log error and raw response body when available for debugging
		if apiResp != nil && len(apiResp.Body) > 0 {
			log.Printf("TagsGet error: %v (status: %d)\nraw body:\n%s", err, apiResp.StatusCode, string(apiResp.Body))
		} else {
			log.Printf("TagsGet error: %v (status: %v)", err, apiResp)
		}
		return
	}
	// Print the typed ContactTagsResponse in a friendly way.
	if out == nil {
		log.Printf("Contact %s tags: (nil response)", contactID)
		return
	}

	// out is a *contacts.ContactTagsResponse
	resp := out
	tags := resp.ContactTagsOrEmpty()
	if len(tags) > 0 {
		log.Printf("Contact %s has %d tags:", contactID, len(tags))
		for _, ct := range tags {
			log.Printf(" - tag id=%s (tag=%s) added=%s", ct.ID, ct.Tag, ct.CDate)
		}

		log.Printf("raw body:\n%s", string(apiResp.Body))
		return
	}

	// Fallback: pretty-print the whole response as JSON
	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Printf("Contact %s tags: %s", contactID, string(b))

}
