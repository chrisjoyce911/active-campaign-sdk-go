//go:build examples

package main

import (
	"context"
	"log"
	"os"

	legacy "github.com/chrisjoyce911/active-campaign-sdk-go/legacy"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()
	baseURL := os.Getenv("ACTIVE_URL")
	token := os.Getenv("ACTIVE_TOKEN")
	// avoid unused variable compile errors in examples — don't print token
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

	// Using the legacy adapter surface for example purposes. These adapters
	// are placeholders and return a not-implemented error until fully wired.
	_, _, err := legacy.GetContact(context.Background(), "287199")
	if err != nil {
		log.Printf("legacy.GetContact returned: %v", err)
	} else {
		log.Printf("legacy.GetContact succeeded (placeholder)")
	}

}
