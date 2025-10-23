//go:build examples

package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"reflect"

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

	// Fetch and print custom fields
	out, apiResp, err := svc.ListCustomFields(context.Background())
	if err != nil {
		// Log error and raw response body when available for debugging
		if apiResp != nil && len(apiResp.Body) > 0 {
			log.Printf("GetContactTags error: %v (status: %d)\nraw body:\n%s", err, apiResp.StatusCode, string(apiResp.Body))
		} else {
			log.Printf("GetContactTags error: %v (status: %v)", err, apiResp)
		}
		return
	}
	// Print the typed ContactTagsResponse in a friendly way.
	if out == nil {
		log.Printf("ListCustomFields: nil response")
		return
	}

	resp := out
	fields := resp.FieldsOrEmpty()
	if len(fields) > 0 {
		log.Printf("Found %d custom fields:", len(fields))
		for _, cf := range fields {
			log.Printf(" - field id=%s title=%s type=%s", cf.ID, cf.Title, cf.Type)
		}

		log.Printf("raw body:\n%s", string(apiResp.Body))

		// Verify raw JSON matches the typed response we received
		var raw contacts.ListFieldsResponse
		if apiResp != nil && len(apiResp.Body) > 0 {
			if err := json.Unmarshal(apiResp.Body, &raw); err != nil {
				log.Printf("failed to unmarshal raw body into ListFieldsResponse: %v", err)
			} else {
				if reflect.DeepEqual(raw.FieldsOrEmpty(), resp.FieldsOrEmpty()) {
					log.Printf("raw JSON matches typed response (fields)")
				} else {
					rb, _ := json.MarshalIndent(raw, "", "  ")
					tb, _ := json.MarshalIndent(resp, "", "  ")
					log.Printf("RAW != TYPED\nraw:\n%s\n\ntyped:\n%s", string(rb), string(tb))
				}
			}
		}
		return
	}

	// Fallback: pretty-print the whole response as JSON
	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Printf("Custom fields: %s", string(b))

}
