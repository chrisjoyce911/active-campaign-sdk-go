package contacts

import (
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

func TestNewRealServiceConstructors(t *testing.T) {
	// NewRealServiceFromDoer should accept any client.Doer
	d := &client.CoreClient{}
	svc := NewRealServiceFromDoer(d)
	if svc == nil {
		t.Fatalf("NewRealServiceFromDoer returned nil")
	}

	// NewRealService expects *client.CoreClient and returns a non-nil RealService
	cc := &client.CoreClient{}
	rs := NewRealService(cc)
	if rs == nil {
		t.Fatalf("NewRealService returned nil")
	}
}
