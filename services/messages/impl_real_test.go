package messages

import (
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

func TestMessages_NewRealServiceConstructors(t *testing.T) {
	_ = NewRealService(&client.CoreClient{})
	_ = NewRealServiceFromDoer(nil)
}
