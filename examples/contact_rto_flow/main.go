package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contacts"

	"github.com/joho/godotenv"
)

func env(k string) string { return strings.TrimSpace(os.Getenv(k)) }

func Run(ctx context.Context, svc *contacts.RealService, email, first, last, phone, company, rto, listID, tagID, campaignTagID, cfCompany, cfRTO string, apply bool) error {
	// updateCustomField centralizes the example-level dry-run/logging and calls the SDK helper
	updateCustomField := func(ctx context.Context, svc *contacts.RealService, contactID, fieldIdentifier, value string, apply bool, hintTitle, hintPerstag string) error {
		if fieldIdentifier == "" {
			return nil
		}
		if !apply {
			fmt.Printf("dry-run: would update custom field %s -> %s for contact %s\n", fieldIdentifier, value, contactID)
			return nil
		}
		out, apiResp, err := svc.UpdateOrCreateFieldValueForContact(ctx, contactID, fieldIdentifier, value)
		if err != nil {
			return fmt.Errorf("UpdateOrCreateFieldValueForContact error: %w (api resp: %+v)", err, apiResp)
		}
		_ = out
		fmt.Println("updated custom field via UpdateOrCreateFieldValueForContact")
		return nil
	}

	// (Field ID detection helper removed; example uses SDK-level UpdateOrCreateFieldValueForContact)

	// 1) Search by email
	out, apiResp, err := svc.SearchByEmail(ctx, email)
	if err != nil {
		return fmt.Errorf("SearchByEmail error: %w (api resp: %+v)", err, apiResp)
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
		if supplied := os.Getenv("ACTIVE_CONTACT_ID"); supplied != "" {
			contactID = supplied
			fmt.Printf("using contact id from env ACTIVE_CONTACT_ID=%s\n", contactID)
		} else if supplied := os.Getenv("CONTACT_ID"); supplied != "" {
			contactID = supplied
			fmt.Printf("using contact id from env CONTACT_ID=%s\n", contactID)
		}
	}

	// 2) Create contact if not found
	if contactID == "" {
		if !apply {
			fmt.Println("dry-run: would create contact with email", email)
		} else {
			// Use OrgName for company (Contact struct has OrgName field). Include LastName and Phone if present.
			req := &contacts.CreateContactRequest{Contact: &contacts.Contact{Email: email, FirstName: first, LastName: last, OrgName: company, Phone: phone}}
			created, apiResp, err := svc.Create(ctx, req)
			if err != nil {
				return fmt.Errorf("Create contact error: %w (api resp: %+v)", err, apiResp)
			}
			contactID = created.Contact.ID
			fmt.Printf("created contact id=%s\n", contactID)
		}
	}

	// 3) Update custom fields
	if err := updateCustomField(ctx, svc, contactID, cfCompany, company, apply, "Company", "company"); err != nil {
		return err
	}
	if err := updateCustomField(ctx, svc, contactID, cfRTO, rto, apply, "RTO ID", "rto_id"); err != nil {
		return err
	}

	// 4) Add to list
	if listID != "" {
		if !apply {
			fmt.Printf("dry-run: would add contact %s to list %s\n", contactID, listID)
		} else {
			req := &contacts.ContactList{Contact: contactID, List: listID, Status: "1"}
			out, apiResp, err := svc.AddContactToList(ctx, req)
			if err != nil {
				return fmt.Errorf("CreateContactList error: %w (api resp: %+v)", err, apiResp)
			}
			_ = out
			fmt.Println("added contact to list")
		}
	}

	// 5) Add tags
	if tagID != "" {
		if !apply {
			fmt.Printf("dry-run: would add tag %s to contact %s\n", tagID, contactID)
		} else {
			req := &contacts.ContactTagRequest{ContactTag: contacts.ContactTagPayload{Contact: contactID, Tag: tagID}}
			out, apiResp, err := svc.CreateContactTag(ctx, req)
			if err != nil {
				return fmt.Errorf("CreateContactTag error: %w (api resp: %+v)", err, apiResp)
			}
			_ = out
			fmt.Println("added tag to contact")
		}
	}

	if campaignTagID != "" {
		if !apply {
			fmt.Printf("dry-run: would add campaign tag %s to contact %s\n", campaignTagID, contactID)
		} else {
			req := &contacts.ContactTagRequest{ContactTag: contacts.ContactTagPayload{Contact: contactID, Tag: campaignTagID}}
			out, apiResp, err := svc.CreateContactTag(ctx, req)
			if err != nil {
				return fmt.Errorf("CreateContactTag campaign error: %w (api resp: %+v)", err, apiResp)
			}
			_ = out
			fmt.Println("added campaign tag to contact")
		}
	}

	return nil
}

func main() {
	if os.Getenv("ACTIVE_URL") == "" {
		_ = godotenv.Load()
	}

	// Flags
	fs := flag.NewFlagSet("contact_rto_flow", flag.ExitOnError)
	apply := fs.Bool("apply", false, "If set, perform mutating operations (create/update/list/tag). Otherwise runs in dry-run mode")
	fs.Parse(os.Args[1:])

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

	if err := Run(ctx, contactsSvc, email, first, last, phone, company, rto, listID, tagID, campaignTagID, cfCompany, cfRTO, *apply); err != nil {
		if os.Getenv("TEST") != "1" {
			log.Fatal(err)
		}
		return
	}
}
