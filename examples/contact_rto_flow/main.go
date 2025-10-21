package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contacts"

	"github.com/joho/godotenv"
)

func env(k string) string { return strings.TrimSpace(os.Getenv(k)) }

func mustIntEnv(k string) string { return env(k) }

func main() {
	_ = godotenv.Load()

	// Flags
	apply := flag.Bool("apply", false, "If set, perform mutating operations (create/update/list/tag). Otherwise runs in dry-run mode")
	flag.Parse()

	// Load env vars
	baseURL := env("ACTIVE_URL")
	token := env("ACTIVE_TOKEN")
	email := env("CONTACT_EMAIL")
	first := env("CONTACT_FIRST_NAME")
	last := env("CONTACT_LAST_NAME")
	phone := env("CONTACT_PHONE")
	company := env("CONTACT_COMPANY_NAME")
	rto := env("CONTACT_RTO_ID")

	listID := env("ACTIVE_CONTACT_LIST")
	tagID := env("ACTIVE_CONTACT_TAG")
	campaignTagID := env("ACTIVE_CONTACT_CAMPAIGN_TAG")

	cfCompany := env("ACTIVE_CONTACT_CF_COMPANY_NAME")
	cfRTO := env("ACTIVE_CONTACT_CF_RTO_ID")

	if baseURL == "" || token == "" || email == "" {
		fmt.Fprintln(os.Stderr, "ACTIVE_URL, ACTIVE_TOKEN and CONTACT_EMAIL must be set in environment")
		os.Exit(2)
	}

	core, err := client.NewCoreClient(baseURL, token)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create client: %v\n", err)
		os.Exit(1)
	}

	// Optional debug: set CLIENT_DEBUG=1 to enable outgoing request debug
	if env("CLIENT_DEBUG") == "1" {
		core.SetDebug(true, os.Stdout)
	}

	contactsSvc := contacts.NewRealService(core)

	ctx := context.Background()

	// small helper: detect a field id by title or perstag (used only for diagnostics)
	detectFieldID := func(ctx context.Context, svc *contacts.RealService, wantTitle string, wantPerstag string) (string, error) {
		resp, _, err := svc.ListCustomFieldsWithOpts(ctx, map[string]string{"limit": "100"})
		if err != nil || resp == nil {
			return "", err
		}
		wantTitleLower := strings.ToLower(strings.TrimSpace(wantTitle))
		wantPerstagLower := strings.ToLower(strings.TrimSpace(wantPerstag))
		for _, f := range resp.FieldsOrEmpty() {
			if wantTitleLower != "" && strings.ToLower(strings.TrimSpace(f.Title)) == wantTitleLower {
				return f.ID, nil
			}
			if wantPerstagLower != "" && strings.ToLower(strings.TrimSpace(f.Perstag)) == wantPerstagLower {
				return f.ID, nil
			}
		}
		return "", nil
	}

	// updateCustomField centralizes the example-level dry-run/logging and calls the SDK helper
	updateCustomField := func(ctx context.Context, svc *contacts.RealService, contactID, fieldIdentifier, value string, apply bool, hintTitle, hintPerstag string) {
		if fieldIdentifier == "" {
			return
		}
		if !apply {
			fmt.Printf("dry-run: would update custom field %s -> %s for contact %s\n", fieldIdentifier, value, contactID)
			return
		}
		out, apiResp, err := svc.UpdateOrCreateFieldValueForContact(ctx, contactID, fieldIdentifier, value)
		if err != nil {
			fmt.Fprintf(os.Stderr, "UpdateOrCreateFieldValueForContact error: %v\n", err)
			if apiResp != nil {
				fmt.Fprintf(os.Stderr, "status=%d body=%s\n", apiResp.StatusCode, string(apiResp.Body))
			}
			// helpful diagnostic: try to find candidate field ids
			if candidate, derr := detectFieldID(ctx, svc, hintTitle, hintPerstag); derr == nil && candidate != "" {
				fmt.Fprintf(os.Stderr, "detected candidate field id=%s (by title/perstag)\n", candidate)
			}
			return
		}
		_ = out
		fmt.Println("updated custom field via UpdateOrCreateFieldValueForContact")
	}

	// (Field ID detection helper removed; example uses SDK-level UpdateOrCreateFieldValueForContact)

	// 1) Search by email
	out, apiResp, err := contactsSvc.SearchByEmail(ctx, email)
	if err != nil {
		fmt.Fprintf(os.Stderr, "SearchByEmail error: %v\n", err)
		if apiResp != nil {
			fmt.Fprintf(os.Stderr, "status=%d body=%s\n", apiResp.StatusCode, string(apiResp.Body))
		}
		os.Exit(1)
	}

	var contactID string
	if out == nil || len(out.Contacts) == 0 {
		fmt.Printf("no contact found for %s\n", email)
	} else {
		// Use typed response to extract ID
		fmt.Printf("search result: found %d contact(s)\n", len(out.Contacts))
		contactID = out.Contacts[0].ID
	}

	// Allow overriding/short-circuit with an env-provided contact id
	if contactID == "" {
		if supplied := env("ACTIVE_CONTACT_ID"); supplied != "" {
			contactID = supplied
			fmt.Printf("using contact id from env ACTIVE_CONTACT_ID=%s\n", contactID)
		} else if supplied := env("CONTACT_ID"); supplied != "" {
			contactID = supplied
			fmt.Printf("using contact id from env CONTACT_ID=%s\n", contactID)
		}
	}

	// 2) Create contact if not found
	if contactID == "" {
		if !*apply {
			fmt.Println("dry-run: would create contact with email", email)
		} else {
			// Use OrgName for company (Contact struct has OrgName field). Include LastName and Phone if present.
			req := &contacts.CreateContactRequest{Contact: &contacts.Contact{Email: email, FirstName: first, LastName: last, OrgName: company, Phone: phone}}
			created, apiResp, err := contactsSvc.Create(ctx, req)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Create contact error: %v\n", err)
				if apiResp != nil {
					fmt.Fprintf(os.Stderr, "status=%d body=%s\n", apiResp.StatusCode, string(apiResp.Body))
				}
				os.Exit(1)
			}
			// Try to pull id
			if created != nil && created.Contact.ID != "" {
				contactID = created.Contact.ID
			} else {
				// best-effort: marshal & inspect
				b, _ := json.Marshal(created)
				fmt.Printf("created response: %s\n", string(b))
			}
			fmt.Println("created contact id:", contactID)
		}
	}

	if contactID == "" {
		fmt.Fprintln(os.Stderr, "unable to determine contact id; aborting further mutating actions")
		os.Exit(1)
	}

	// 3) Update custom fields via UpdateFieldValueForContact
	// Use ACTIVE_CONTACT_CF_* env vars mapping to field ids
	if cfCompany != "" {
		updateCustomField(ctx, contactsSvc, contactID, cfCompany, company, *apply, "Company Name", "COMPANY_NAME")
	}
	if cfRTO != "" {
		updateCustomField(ctx, contactsSvc, contactID, cfRTO, rto, *apply, "RTO ID", "RTO_ID")
	}

	// 4) Add contact to list
	if listID != "" {
		if !*apply {
			fmt.Printf("dry-run: would add contact %s to list %s\n", contactID, listID)
		} else {
			req := &contacts.UpdateListStatusForContactRequest{ContactList: &contacts.ContactList{Contact: contactID, List: listID, Status: "1"}}
			_, apiResp, err := contactsSvc.UpdateListStatus(ctx, req)
			if err != nil {
				fmt.Fprintf(os.Stderr, "UpdateListStatus error: %v\n", err)
				if apiResp != nil {
					fmt.Fprintf(os.Stderr, "status=%d body=%s\n", apiResp.StatusCode, string(apiResp.Body))
				}
			} else {
				fmt.Println("added contact to list")
			}
		}
	}

	// 5) Add tags using the contacts service (POST /contactTags)
	tryTag := func(tid string) {
		if tid == "" {
			return
		}
		if !*apply {
			fmt.Printf("dry-run: would add tag %s to contact %s\n", tid, contactID)
			return
		}
		ct := &contacts.ContactTagRequest{ContactTag: contacts.ContactTagPayload{Contact: contactID, Tag: tid}}
		_, apiResp, err := contactsSvc.CreateContactTag(ctx, ct)
		if err != nil {
			fmt.Fprintf(os.Stderr, "CreateContactTag error: %v\n", err)
			if apiResp != nil {
				fmt.Fprintf(os.Stderr, "status=%d body=%s\n", apiResp.StatusCode, string(apiResp.Body))
			}
		} else {
			fmt.Printf("added tag %s\n", tid)
		}
	}

	tryTag(tagID)
	tryTag(campaignTagID)

	fmt.Println("done")
}
