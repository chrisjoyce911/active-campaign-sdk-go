package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contacts"
	"github.com/joho/godotenv"
	"golang.org/x/time/rate"
)

func main() {
	_ = godotenv.Overload()

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
	limiter := rate.NewLimiter(rate.Limit(0.5), 2) // Start conservative: 0.5 requests per second (2s per request), burst of 2
	lastTime := time.Now()
	apiRateLimit := rate.Limit(5.0) // API limit: 5 requests per second according to docs

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

		// Check if we got rate limit headers and adjust our limiter accordingly
		if apiResp != nil && apiResp.RateLimitLimit != "" {
			if limit, parseErr := strconv.Atoi(apiResp.RateLimitLimit); parseErr == nil && limit > 0 {
				newLimit := rate.Limit(float64(limit) * 0.8) // Use 80% of the limit to be safe
				if newLimit != limiter.Limit() {
					fmt.Printf("Adjusting rate limiter from %.1f to %.1f req/s based on API limit %d\n",
						float64(limiter.Limit()), float64(newLimit), limit)
					limiter.SetLimit(newLimit)
					apiRateLimit = newLimit
				}
			}
		}

		// If RateLimit-Remaining is low, be even more conservative
		if apiResp != nil && apiResp.RateLimitRemaining != "" {
			if remaining, parseErr := strconv.Atoi(apiResp.RateLimitRemaining); parseErr == nil {
				if remaining <= 2 && remaining > 0 { // Only 2 or fewer requests left
					conservativeLimit := apiRateLimit * 0.5 // Half the rate when running low
					if conservativeLimit < limiter.Limit() {
						fmt.Printf("Low on requests (%d remaining), slowing down to %.1f req/s\n",
							remaining, float64(conservativeLimit))
						limiter.SetLimit(conservativeLimit)
					}
				} else if remaining == 0 {
					fmt.Printf("No requests remaining in current window\n")
				}
			}
		}

		if apiResp != nil && apiResp.StatusCode == 429 {
			fmt.Printf("Rate limited! Retry-After: %s, RateLimit-Limit: %s, RateLimit-Remaining: %s\n",
				apiResp.RetryAfter, apiResp.RateLimitLimit, apiResp.RateLimitRemaining)
			if apiResp.RetryAfter != "" {
				// Parse Retry-After as seconds (API docs say it's duration in seconds)
				var sleepDuration time.Duration
				if seconds, parseErr := time.ParseDuration(apiResp.RetryAfter + "s"); parseErr == nil {
					sleepDuration = seconds
				} else if secs, atoiErr := strconv.Atoi(apiResp.RetryAfter); atoiErr == nil {
					sleepDuration = time.Duration(secs) * time.Second
				} else {
					fmt.Printf("Could not parse Retry-After header: %s\n", apiResp.RetryAfter)
					continue
				}
				fmt.Printf("Sleeping for %v...\n", sleepDuration)
				time.Sleep(sleepDuration)
				lastTime = time.Now()
				continue
			}
		}

		if err != nil {
			fmt.Printf("Error: %v (status=%d, duration=%v, time_since_last=%v)\n", err, apiResp.StatusCode, duration, timeSinceLast)
		} else {
			// Show current rate limit status
			fmt.Printf("Success - RateLimit-Limit: %s, RateLimit-Remaining: %s",
				apiResp.RateLimitLimit, apiResp.RateLimitRemaining)
			if len(resp.Contacts) > 0 {
				fmt.Printf(" (contact ID: %s, duration=%v, time_since_last=%v)\n", resp.Contacts[0].ID, duration, timeSinceLast)
			} else {
				fmt.Printf(" (no contacts found, duration=%v, time_since_last=%v)\n", duration, timeSinceLast)
			}
		}

		lastTime = time.Now()
		// No manual sleep, rate limiter handles pacing
	}
}
