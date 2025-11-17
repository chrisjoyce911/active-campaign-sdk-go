package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contacts"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/lists"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	activeURL := os.Getenv("ACTIVE_URL")
	token := os.Getenv("ACTIVE_TOKEN")
	if activeURL == "" || token == "" {
		fmt.Println("Set ACTIVE_URL and ACTIVE_TOKEN to run this example")
		return
	}

	// safety guard: if LISTS_SAFE=false we default to deleting created resources
	safe := os.Getenv("LISTS_SAFE")
	defaultDelete := false
	if safe == "false" {
		defaultDelete = true
		fmt.Println("WARNING: LISTS_SAFE=false, example will create & potentially modify data")
	}

	fs := flag.NewFlagSet("lists_create_list", flag.ExitOnError)
	deleteAfter := fs.Bool("delete", defaultDelete, "delete the created list after running the example")
	fs.Parse(os.Args[1:])

	c, err := client.NewCoreClient(activeURL, token)
	if err != nil {
		fmt.Println("failed to create client:", err)
		return
	}
	listsSvc := lists.NewRealService(c)
	contactsSvc := contacts.NewRealService(c)

	// determine a contact to use as the list owner and subscriber (prefer ID, fallback to email search)
	contactID := os.Getenv("ACTIVE_CONTACTID")
	contactEmail := os.Getenv("ACTIVE_EMAIL")
	contact := contactID
	if contact == "" && contactEmail != "" {
		if sr, _, err := contactsSvc.SearchByEmail(context.Background(), contactEmail); err == nil {
			if len(sr.Contacts) > 0 {
				contact = sr.Contacts[0].ID
			}
		} else {
			fmt.Println("SearchByEmail error:", err)
		}
	}

	req := lists.CreateListRequest{List: lists.List{
		Name:      "Example List",
		StringID:  fmt.Sprintf("example-list-%d", time.Now().Unix()),
		Channel:   "email",
		SenderURL: "https://example.com",
		User:      lists.StringToUser(contact),
	}}

	out, apiResp, err := listsSvc.CreateList(context.Background(), req)
	if err != nil {
		fmt.Println("CreateList error:", err)
		if apiResp != nil {
			fmt.Println("status:", apiResp.StatusCode)
		}
		return
	}
	fmt.Printf("Created list id=%s name=%s\n", out.List.ID, out.List.Name)

	// subscribe the previously-resolved contact to the list (best-effort)
	if contact != "" {
		req := &contacts.UpdateListStatusForContactRequest{
			ContactList: &contacts.ContactList{
				Contact: contact,
				List:    out.List.ID,
				Status:  "1", // 1 = subscribe per API
			},
		}
		if upOut, upResp, upErr := contactsSvc.UpdateListStatus(context.Background(), req); upErr != nil {
			fmt.Println("UpdateListStatus error:", upErr)
			if upResp != nil {
				fmt.Println("status:", upResp.StatusCode)
			}
		} else {
			fmt.Println("Added contact to list. Response:", upOut)
		}
	}

	if *deleteAfter {
		delResp, delErr := listsSvc.DeleteList(context.Background(), out.List.ID)
		if delErr != nil {
			fmt.Println("DeleteList error:", delErr)
			if delResp != nil {
				fmt.Println("status:", delResp.StatusCode)
			}
			return
		}
		if delResp != nil {
			fmt.Println("Deleted list, status:", delResp.StatusCode)
		} else {
			fmt.Println("Deleted list")
		}
	}
}
