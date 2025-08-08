package main

import (
	"log"
	"os"

	ac "github.com/chrisjoyce911/active-campaign-sdk-go"
)

func main() {

	baseURL := os.Getenv("ACTIVE_URL")
	token := os.Getenv("ACTIVE_TOKEN")

	campaign, err := ac.NewClient(
		&ac.ClientOpts{
			BaseUrl: baseURL,
			Token:   token,
		},
	)
	if err != nil {
		panic(err)
	}

	contacts, _, _ := campaign.Contacts.SearchEmail("mujaddid2004+work@hotmail.com")

	for i := range contacts.Contact {
		contact := contacts.Contact[i]
		log.Printf("Contact ID: %s", contact.ID)
		log.Printf("Email: %s", contact.Email)
		log.Printf("Created Date: %s", contact.Cdate)
		log.Printf("Updated Date: %s", contact.Udate)
		log.Printf("Organization ID: %v", contact.Orgid)
		log.Printf("Organization: %s", contact.Organization)

		// Display custom field values if any
		if len(contact.FieldValues) > 0 {
			log.Printf("Custom Field Values:")
			for j, fieldValue := range contact.FieldValues {
				log.Printf("  Field %d: %s = %v", j+1, fieldValue.Field, fieldValue.Value)
			}
		}

		log.Printf("--- Contact Links ---")
		log.Printf("Contact Lists: %s", contact.Links.ContactLists)
		log.Printf("Contact Tags: %s", contact.Links.ContactTags)
		log.Printf("Field Values: %s", contact.Links.FieldValues)
		log.Printf("Notes: %s", contact.Links.Notes)
		log.Printf("Contact Deals: %s", contact.Links.ContactDeals)
		log.Printf("Contact Automations: %s", contact.Links.ContactAutomations)
		log.Printf("Bounce Logs: %s", contact.Links.BounceLogs)
		log.Printf("Geo IPs: %s", contact.Links.GeoIps)
		log.Printf("Tracking Logs: %s", contact.Links.TrackingLogs)
		log.Printf("Score Values: %s", contact.Links.ScoreValues)
		log.Printf("Plus Append: %s", contact.Links.PlusAppend)
		log.Printf("Contact Data: %s", contact.Links.ContactData)
		log.Printf("Contact Goals: %s", contact.Links.ContactGoals)
		log.Printf("Contact Logs: %s", contact.Links.ContactLogs)
		log.Printf("Deals: %s", contact.Links.Deals)

		log.Println("=") // Separator between contacts
	}

}
