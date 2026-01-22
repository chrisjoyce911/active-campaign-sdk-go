# Adaptive Rate Limiting (simple)

A minimal example that respects ActiveCampaign rate limits while staying near the allowed throughput. It reacts to the `RateLimit-Limit`, `RateLimit-Remaining`, and `Retry-After` headers to avoid 429s.

## How it works

- Start gentle at **0.5 req/s** with a burst of 1.
- For each successful response:
    - Read `RateLimit-Limit` and `RateLimit-Remaining`.
    - Compute `ratio = remaining / limit`.
    - If `ratio >= 0.7`, ramp up by 10% but cap at **70% of the advertised limit**.
    - If `ratio <= 0.4`, slow down toward **30% of the limit** (never below 0.2 req/s).
- On `429`:
    - Sleep for the `Retry-After` seconds (fallback 1s if missing).
    - Halve the current rate and cap it to **30% of the limit**, but never below 0.2 req/s.

Key tuning constants (see `main.go`):

```go
minLimit     = 0.2   // floor rate (req/s)
maxScale     = 0.7   // max fraction of RateLimit-Limit
rampUpFactor = 1.1   // 10% growth when we have headroom
```

## Running

```bash
export ACTIVE_URL="https://youraccount.api-us1.com"
export ACTIVE_TOKEN="your-api-token"
export CONTACT_EMAIL="test@example.com" # optional override

go run main.go
```

## Sample output

```
🚀 Adaptive rate limiting (simple)
📊 Initial limiter: 0.50 req/s

Request 1 (0.50 req/s): ✅ Success (contact=none, limit=5, remaining=3)
Request 2 (0.60 req/s): ⏱️  hit rate limit (Retry-After=1, limit=5, remaining=0)
Request 3 (0.30 req/s): ✅ Success (contact=none, limit=5, remaining=1)
...
```

Tweak the three constants if you still see 429s (lower `maxScale` or `rampUpFactor`) or want more throughput (raise them gradually)." 