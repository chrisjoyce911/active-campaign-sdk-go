package main

import (
	"context"
	"fmt"
	"math"
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
	limiter := rate.NewLimiter(rate.Limit(0.5), 1) // Start gentle: 0.5 req/s
	minLimit := rate.Limit(0.2)                    // Never go below 0.2 req/s unless Retry-After says otherwise
	maxScale := 0.7                                // Never exceed 70% of the advertised limit
	rampUpFactor := 1.1                            // Grow 10% when we have headroom

	fmt.Println("🚀 Adaptive rate limiting (simple)")
	fmt.Printf("📊 Initial limiter: %.2f req/s\n\n", float64(limiter.Limit()))

	for i := 1; i <= 30; i++ {
		if err := limiter.Wait(ctx); err != nil {
			fmt.Printf("Request %d: limiter wait error: %v\n", i, err)
			continue
		}

		fmt.Printf("Request %d (%.2f req/s): ", i, float64(limiter.Limit()))
		resp, apiResp, err := svc.SearchByEmail(ctx, email)

		// Handle rate limit errors first
		if apiResp != nil && apiResp.StatusCode == 429 {
			fmt.Printf("⏱️  hit rate limit (Retry-After=%s, limit=%s, remaining=%s)\n",
				apiResp.RetryAfter, apiResp.RateLimitLimit, apiResp.RateLimitRemaining)
			// Respect Retry-After if present, otherwise short backoff
			sleepDuration := 1 * time.Second
			if apiResp.RetryAfter != "" {
				if seconds, parseErr := time.ParseDuration(apiResp.RetryAfter + "s"); parseErr == nil {
					sleepDuration = seconds
				} else if secs, atoiErr := strconv.Atoi(apiResp.RetryAfter); atoiErr == nil {
					sleepDuration = time.Duration(secs) * time.Second
				}
			}
			time.Sleep(sleepDuration)
			// Drop to a cautious rate after a 429 (halve current, but not below min and not above 30% of limit)
			newLimit := limiter.Limit() / 2
			if apiResp.RateLimitLimit != "" {
				if limit, parseErr := strconv.Atoi(apiResp.RateLimitLimit); parseErr == nil && limit > 0 {
					maxAfterHit := rate.Limit(float64(limit) * 0.3)
					if newLimit > maxAfterHit {
						newLimit = maxAfterHit
					}
				}
			}
			if newLimit < minLimit {
				newLimit = minLimit
			}
			limiter.SetLimit(newLimit)
			continue
		}

		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			continue
		}

		// Successful response: print minimal info
		found := "none"
		if len(resp.Contacts) > 0 {
			found = resp.Contacts[0].ID
		}
		fmt.Printf("✅ Success (contact=%s, limit=%s, remaining=%s)\n",
			found, apiResp.RateLimitLimit, apiResp.RateLimitRemaining)

		// Adjust limiter based on headers
		if apiResp != nil && apiResp.RateLimitLimit != "" {
			if limit, parseErr := strconv.Atoi(apiResp.RateLimitLimit); parseErr == nil && limit > 0 {
				remaining := limit
				if apiResp.RateLimitRemaining != "" {
					if r, remErr := strconv.Atoi(apiResp.RateLimitRemaining); remErr == nil {
						remaining = r
					}
				}
				ratio := float64(remaining) / float64(limit)

				target := limiter.Limit()
				maxAllowed := rate.Limit(float64(limit) * maxScale)
				minAllowed := rate.Limit(float64(limit) * 0.2)
				if minAllowed < minLimit {
					minAllowed = minLimit
				}

				// Speed up cautiously when plenty of budget remains, capped
				if ratio >= 0.7 {
					proposed := rate.Limit(float64(limiter.Limit()) * rampUpFactor)
					if proposed > maxAllowed {
						proposed = maxAllowed
					}
					if proposed > target {
						target = proposed
					}
				}
				// Slow down early when budget gets tight; clamp to minAllowed
				if ratio <= 0.4 {
					target = rate.Limit(math.Max(float64(minAllowed), float64(limit)*0.3))
				}

				if target != limiter.Limit() {
					limiter.SetLimit(target)
				}
			}
		}
	}
}
