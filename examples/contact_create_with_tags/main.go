//go:build examples

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contacts"
)

// Example showing how to map a simple struct into Contact + tags and create
// a contact with custom field values and attach tags.
func main() {
	_ = godotenv.Load()
	baseURL := os.Getenv("ACTIVE_URL")
	token := os.Getenv("ACTIVE_TOKEN")

	if baseURL == "" || token == "" {
		log.Printf("ACTIVE_URL/ACTIVE_TOKEN not set — example will run in placeholder mode and will not call the API")
	}

	core, err := client.NewCoreClient(baseURL, token)
	if err != nil {
		log.Fatalf("failed to create core client: %v", err)
	}
	svc := contacts.NewRealService(core)

	// In a real project you'd use generated constants from genconstants (Fields, Tags)
	// Here we provide a minimal mapping for the example.
	fieldIDByName := map[string]string{"Star": "21", "Speed": "22"}
	tagNameToID := map[string]string{"VIP": "100", "Trial": "101"}

	type MyContact struct {
		Email     string `contact:"Email"`
		FirstName string `contact:"FirstName,omitempty"`
		LastName  string `contact:"LastName,omitempty"`
		Star      string `fieldValues:"Star,omitempty"`
		Speed     string `fieldValues:"Speed"`
		Tags      string `tags:"Tags,omitempty"` // comma-separated tag names or ids
	}

	src := MyContact{
		Email:     "jdoe@example.com",
		FirstName: "John",
		LastName:  "Doe",
		Star:      "Gold",
		Speed:     "2008-01-20",
		Tags:      "VIP,101", // mix of name and explicit id
	}

	c, tagIDs, err := contacts.MapToContact(src, fieldIDByName, tagNameToID)
	if err != nil {
		log.Fatalf("map error: %v", err)
	}

	// Prepare request and call Create
	req := &contacts.CreateContactRequest{Contact: &c}
	created, apiResp, err := svc.Create(context.Background(), req)
	if err != nil {
		if apiResp != nil {
			log.Fatalf("create contact failed: %v (status=%d) body=%s", err, apiResp.StatusCode, string(apiResp.Body))
		}
		log.Fatalf("create contact failed: %v", err)
	}
	if created == nil || created.Contact == nil {
		log.Fatalf("create returned nil contact")
	}

	fmt.Printf("Created contact id=%s email=%s\n", created.Contact.ID, created.Contact.Email)

	// Attach tags (uses CreateContactTag which expects tag id)
	for _, tid := range tagIDs {
		ctReq := &contacts.ContactTagRequest{ContactTag: contacts.ContactTagPayload{Contact: created.Contact.ID, Tag: tid}}
		_, apiResp, err := svc.CreateContactTag(context.Background(), ctReq)
		if err != nil {
			if apiResp != nil {
				log.Printf("attach tag %s failed: %v (status=%d) body=%s", tid, err, apiResp.StatusCode, string(apiResp.Body))
			} else {
				log.Printf("attach tag %s failed: %v", tid, err)
			}
			continue
		}
		fmt.Printf("Attached tag id=%s to contact %s\n", tid, created.Contact.ID)
	}
}
