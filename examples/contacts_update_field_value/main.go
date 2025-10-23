package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contacts"

	"github.com/joho/godotenv"
)

func env(k string) string { return strings.TrimSpace(os.Getenv(k)) }

func main() {
	_ = godotenv.Load()
	apply := flag.Bool("apply", false, "perform mutating calls against API")
	flag.Parse()

	base := env("ACTIVE_URL")
	token := env("ACTIVE_TOKEN")
	contact := env("ACTIVE_CONTACTID")
	company := env("CONTACT_COMPANY_NAME")
	field := env("ACTIVE_CONTACT_CF_COMPANY_NAME")

	if base == "" || token == "" || contact == "" || field == "" {
		fmt.Fprintln(os.Stderr, "set ACTIVE_URL, ACTIVE_TOKEN, ACTIVE_CONTACTID and ACTIVE_CONTACT_CF_COMPANY_NAME in .env")
		os.Exit(2)
	}

	cc, err := client.NewCoreClient(base, token)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create client: %v\n", err)
		os.Exit(1)
	}

	cc.SetDebug(true, os.Stdout)

	svc := contacts.NewRealService(cc)
	ctx := context.Background()

	if !*apply {
		fmt.Printf("dry-run: would set field %s -> %s for contact %s\n", field, company, contact)
		return
	}

	req := &contacts.FieldValuePayload{Contact: contact, Field: field, Value: company}
	out, apiResp, err := svc.UpdateFieldValueForContact(ctx, req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		if apiResp != nil {
			fmt.Fprintf(os.Stderr, "status=%d body=%s\n", apiResp.StatusCode, string(apiResp.Body))
		}
		os.Exit(1)
	}

	fmt.Printf("ok: created/updated fieldValue id=%s\n", out.FieldValue.ID)
}
