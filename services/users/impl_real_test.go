package users

import (
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

func TestUsers_NewRealServiceConstructors(t *testing.T) {
	_ = NewRealService(&client.CoreClient{})
}
