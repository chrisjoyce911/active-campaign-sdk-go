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

	_ = godotenv.Overload()
	baseURL := os.Getenv("ACTIVE_URL")
	token := os.Getenv("ACTIVE_TOKEN")

	if baseURL == "" {
		log.Printf("ACTIVE_URL not set; running example in placeholder mode")
	} else {
		log.Printf("ACTIVE_URL set to %s", baseURL)
	}
	if token == "" {
		log.Printf("ACTIVE_TOKEN not set")
	} else {
		log.Printf("ACTIVE_TOKEN is set %s", token)
	}

	// Build core client and contacts service
	c, err := client.NewCoreClient(baseURL, token)
	if err != nil {
		log.Fatalf("failed to create core client: %v", err)
	}
	svc := contacts.NewRealService(c)

	// Search for a contact by email
	email := "brooke@joyce.id.au"
	out, apiResp, err := svc.SearchByEmail(context.Background(), email)
	if err != nil {
		if apiResp != nil && len(apiResp.Body) > 0 {
			log.Printf("SearchByEmail error: %v (status: %d)\nraw body:\n%s", err, apiResp.StatusCode, string(apiResp.Body))
		} else {
			log.Printf("SearchByEmail error: %v (resp: %v)", err, apiResp)
		}
		return
	}

	if out == nil {
		log.Printf("no contacts returned for %s", email)
		return
	}

	b, _ := json.MarshalIndent(out, "", "  ")
	log.Printf("Search results for %s:\n%s", email, string(b))

	// log.Printf("raw body:\n%s", string(apiResp.Body))

}
