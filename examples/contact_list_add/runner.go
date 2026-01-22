package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contacts"
)

type runConfig struct {
	ContactID string
	ListID    string
	Status    string
}

func run(ctx context.Context, svc *contacts.RealService, cfg runConfig, apply bool, w io.Writer) error {
	if w == nil {
		w = io.Discard
	}
	if cfg.ContactID == "" || cfg.ListID == "" {
		return fmt.Errorf("contact id and list id are required")
	}

	status := cfg.Status
	if status == "" {
		status = "1"
	}

	if !apply {
		fmt.Fprintf(w, "dry-run: would subscribe contact %s to list %s with status %s\n", cfg.ContactID, cfg.ListID, status)
		return nil
	}

	payload := &contacts.AddContactToListPayload{Contact: contacts.ContactID(cfg.ContactID), List: contacts.ListID(cfg.ListID), Status: status}
	resp, apiResp, err := svc.AddContactToList(ctx, payload)
	if err != nil {
		if apiResp != nil {
			return fmt.Errorf("AddContactToList failed with status %d: %w", apiResp.StatusCode, err)
		}
		return fmt.Errorf("AddContactToList failed: %w", err)
	}

	membershipID := ""
	if resp != nil && resp.ContactList != nil {
		membershipID = resp.ContactList.ID
	}
	fmt.Fprintf(w, "contact %s subscribed to list %s (membership id %s)\n", cfg.ContactID, cfg.ListID, membershipID)
	return nil
}

func envOrFatal(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s not set", key)
	}
	return value
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

func filterTestFlags(args []string) []string {
	if len(args) == 0 {
		return args
	}
	filtered := make([]string, 0, len(args))
	for _, a := range args {
		if strings.HasPrefix(a, "-test.") {
			continue
		}
		filtered = append(filtered, a)
	}
	return filtered
}
