package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contacts"
	"github.com/joho/godotenv"
	"golang.org/x/time/rate"
)

func main() {
	_ = godotenv.Load()

	url := os.Getenv("ACTIVE_URL")
	token := os.Getenv("ACTIVE_TOKEN")
	email := os.Getenv("CONTACT_EMAIL")
	if email == "" {
		email = "chris@joyce.au" // default
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
	limiter := rate.NewLimiter(rate.Limit(0.5), 2) // 0.5 requests per second (2s per request), burst of 2
	lastTime := time.Now()

	for i := 1; i <= 50; i++ {
		// Wait for rate limiter
		err := limiter.Wait(ctx)
		if err != nil {
			fmt.Printf("Request %d: Rate limiter error: %v\n", i, err)
			continue
		}

		fmt.Printf("Request %d: ", i)
		start := time.Now()
		resp, apiResp, err := svc.SearchByEmail(ctx, email)
		duration := time.Since(start)
		timeSinceLast := time.Since(lastTime)

		if apiResp != nil && apiResp.StatusCode == 429 {
			fmt.Printf("Rate limited! Retry after: %s\n", apiResp.RetryAfter)
			if apiResp.RetryAfter != "" {
				if seconds, parseErr := time.ParseDuration(apiResp.RetryAfter + "s"); parseErr == nil {
					fmt.Printf("Sleeping for %v...\n", seconds)
					time.Sleep(seconds)
					lastTime = time.Now()
					continue
				}
			}
		}

		if err != nil {
			fmt.Printf("Error: %v (status=%d, duration=%v, time_since_last=%v)\n", err, apiResp.StatusCode, duration, timeSinceLast)
		} else {
			if len(resp.Contacts) > 0 {
				fmt.Printf("Success (contact ID: %s, duration=%v, time_since_last=%v)\n", resp.Contacts[0].ID, duration, timeSinceLast)
			} else {
				fmt.Printf("Success (no contacts found, duration=%v, time_since_last=%v)\n", duration, timeSinceLast)
			}
		}

		lastTime = time.Now()
		// No manual sleep, rate limiter handles pacing
	}
}
