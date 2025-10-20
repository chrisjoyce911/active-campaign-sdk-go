package users

import "github.com/chrisjoyce911/active-campaign-sdk-go/client"

// NewRealService returns a concrete UsersService wired to CoreClient.
func NewRealService(c *client.CoreClient) UsersService {
	return &service{client: c}
}
